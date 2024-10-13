package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/uvio-network/apiserver/pkg/worker/budget"
	"github.com/uvio-network/apiserver/pkg/worker/workerhandler"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/logger/meta"
	"github.com/xh3b4sd/rescue"
	"github.com/xh3b4sd/rescue/engine"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	// Han are the worker specific handlers implementing the actual business
	// logic.
	Han []workerhandler.Interface
	Log logger.Interface
	// Res is the rescue engine used to participate in the distributed task queue.
	Res rescue.Interface
}

type Worker struct {
	han []workerhandler.Interface
	log logger.Interface
	res rescue.Interface
}

func New(c Config) *Worker {
	if len(c.Han) == 0 {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Han must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Res == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Res must not be empty", c)))
	}

	return &Worker{
		han: c.Han,
		log: c.Log,
		res: c.Res,
	}
}

func (w *Worker) Daemon() {
	{
		w.log.Log(
			context.Background(),
			"level", "info",
			"message", "worker searching for tasks",
			"addr", w.res.Listen(),
		)
	}

	{
		w.create()
	}

	go func() {
		for {
			w.expire()
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		for {
			w.search()
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			w.ticker()
			time.Sleep(5 * time.Second)
		}
	}()

	{
		select {}
	}
}

func (w *Worker) create() {
	var err error

	// Ensure any task template that the respective handlers require.
	for _, x := range w.han {
		var tas *task.Task
		{
			tas = x.Create()
		}

		if tas == nil {
			continue
		}

		var lis []*task.Task
		{
			lis, err = w.res.Lister(&task.Task{Meta: tas.Meta})
			if err != nil {
				w.lerror(tracer.Mask(err))
			}
		}

		// If the task that our worker handler wants to see created does not exist
		// at all, we simply create it here.
		if len(lis) == 0 {
			err = w.res.Create(tas)
			if err != nil {
				w.lerror(tracer.Mask(err))
			}

			continue
		}

		// If the task that we are looking for is a task template, then there is
		// only exactly one object for it. So if there are more than exactly one
		// task object, then it might mean that we found tasks with other
		// specificcations which we should not touch, and so we ignore them.
		if len(lis) > 1 {
			continue
		}

		// Skip this task template if it already resembles the desired state.
		if lis[0].Has(tas) {
			continue
		}

		{
			lis[0].Core.Set().Bypass(true)
		}

		// Delete the existing task template.
		{
			err = w.res.Delete(lis[0])
			if err != nil {
				w.lerror(tracer.Mask(err))
			}
		}

		// Update to the new task template.
		{
			err = w.res.Create(tas)
			if err != nil {
				w.lerror(tracer.Mask(err))
			}
		}
	}
}

func (w *Worker) expire() {
	err := w.res.Expire()
	if err != nil {
		w.lerror(tracer.Mask(err))
	}
}

func (w *Worker) lerror(err error) {
	e, o := err.(*tracer.Error)
	if o {
		w.log.Log(
			context.Background(),
			"level", "error",
			"message", e.Error(),
			"description", e.Desc,
			"docs", e.Docs,
			"kind", e.Kind,
			"stack", tracer.Stack(e),
		)
	} else {
		w.log.Log(
			context.Background(),
			"level", "error",
			"message", err.Error(),
			"stack", tracer.Stack(err),
		)
	}
}

func (w *Worker) search() {
	var err error

	var tas *task.Task
	{
		tas, err = w.res.Search()
		if engine.IsTaskNotFound(err) {
			return
		} else if err != nil {
			w.lerror(tracer.Mask(err))
		}
	}

	// It may happen that a worker loses the underlying connection to the Redis
	// server, for whatever reason. In such a case the task received from above is
	// nil and we simply return here in order to prevent runtime panics.
	if tas == nil {
		return
	}

	var bud *budget.Budget
	{
		bud = budget.New()
	}

	// We track the current and desired amount of handlers for the current task in
	// order to decide whether to delete the task after all handlers got invoked.
	// The desired amount of handlers that can process the current task are those
	// that return true when calling Handler.Filter. The desired amount of workers
	// are then those that do not return an error when calling Handler.Ensure.
	var cur int
	var des int

	for _, x := range w.han {
		if !x.Filter(tas) {
			continue
		}

		{
			des++
		}

		{
			w.log.Log(
				logCtx(tas),
				"level", "info",
				"message", "processing worker task",
			)
		}

		{
			err = x.Ensure(tas, bud)
			if err != nil {
				w.lerror(tracer.Mask(err))
			} else if !bud.Break() || tas.Pag() || (tas.Cron != nil && tas.Cron.Exi().Adefer()) {
				cur++
			}
		}

		// We want to requeue a task that carries a paging pointer, whether the task
		// execution failed or not. So if the given task carries a paging pointer
		// that indicates to continue processing at a certain stage of the worker
		// process itself, then the rescue engine will expire the task immediately
		// and update the synced paging state so that another worker can pick up the
		// task again ASAP.
		if tas.Pag() {
			w.log.Log(
				logCtx(tas),
				"level", "info",
				"message", "task being requeued",
				"paging", tas.Sync.Get(task.Paging),
			)
		}

		// We want to requeue a task that carries a deferral statement, whether the
		// task execution failed or not. So if the given task carries a deferral
		// statement that indicates to continue processing at a later point in time,
		// then the rescue engine will expire the task immediately and update the
		// task's next tick so that another worker can pick up the task again after
		// the desired hold off period.
		if tas.Cron != nil && tas.Cron.Exi().Adefer() {
			w.log.Log(
				logCtx(tas),
				"level", "info",
				"message", "task being requeued",
				"@defer", tas.Cron.Get().Adefer(),
			)
		}

		// We have to account for the worker budget when processing a task. Calling
		// Handler.Ensure may use up the entire budget and it may break through the
		// budget or it may not. Breaking through the budget means that there is
		// still work left to do. And so not breaking the worker budget tells us
		// here that Handler.Ensure successfully resolved the task from its own
		// point of view, allowing us to count with it towards the desired amount of
		// handlers we that we track.
		if bud.Break() {
			w.log.Log(
				logCtx(tas),
				"level", "warning",
				"message", "task budget exhausted",
			)
		}
	}

	// If the current and desired amount of handlers match, we can delete the
	// task, assuming that it got properly resolved.
	if cur != 0 && des != 0 && cur == des {
		err = w.res.Delete(tas)
		if err != nil {
			w.lerror(tracer.Mask(err))
		}
	}
}

func (w *Worker) ticker() {
	err := w.res.Ticker()
	if err != nil {
		w.lerror(tracer.Mask(err))
	}
}

func logCtx(tas *task.Task) context.Context {
	ctx := context.Background()

	if tas.Core != nil {
		for k, v := range *tas.Core {
			ctx = meta.Add(ctx, k, v)
		}
	}

	if tas.Cron != nil {
		for k, v := range *tas.Cron {
			ctx = meta.Add(ctx, k, v)
		}
	}

	if tas.Gate != nil {
		for k, v := range *tas.Gate {
			ctx = meta.Add(ctx, k, v)
		}
	}

	if tas.Meta != nil {
		for k, v := range *tas.Meta {
			ctx = meta.Add(ctx, k, v)
		}
	}

	if tas.Node != nil {
		for k, v := range *tas.Node {
			ctx = meta.Add(ctx, k, v)
		}
	}

	if tas.Root != nil {
		for k, v := range *tas.Root {
			ctx = meta.Add(ctx, k, v)
		}
	}

	if tas.Sync != nil {
		for k, v := range *tas.Sync {
			ctx = meta.Add(ctx, k, v)
		}
	}

	return ctx
}
