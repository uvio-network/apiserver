package redigo

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/server/converter"
	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/uvio-network/apiserver/pkg/storage/notestorage"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/objectid"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"
)

type run struct{}

func (r *run) runE(cmd *cobra.Command, arg []string) error {
	var red redigo.Interface
	{
		red = redigo.Default()
	}

	var sto storage.Interface
	{
		sto = storage.New(storage.Config{
			Log: logger.Default(),
			Red: red,
		})
	}

	// --------------------------------------------------------------------- //

	var cod int
	{
		cod = 0
	}

	if cod == 0 {
		err := r.creNot(sto, objectid.ID(arg[0]))
		if err != nil {
			return tracer.Mask(err)
		}
	}

	if cod == 1 {
		err := r.updExp(sto, arg)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (r *run) creNot(sto storage.Interface, use objectid.ID) error {
	var err error

	var not *notestorage.Object
	{
		not = &notestorage.Object{
			Kind:     objectlabel.NoteKindUserAbstained,
			Owner:    use,
			Resource: objectid.ID("1731243324474059"),
		}
	}

	{
		err = sto.Note().CreateNote([]*notestorage.Object{not})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (r *run) seaPos(sto storage.Interface, cla objectid.ID) (*poststorage.Object, error) {
	var err error

	var pos []*poststorage.Object
	{
		pos, err = sto.Post().SearchPost([]objectid.ID{cla})
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	if len(pos) != 1 {
		return nil, tracer.Mask(runtime.ExecutionFailedError)
	}

	return pos[0], nil
}

func (r *run) updExp(sto storage.Interface, arg []string) error {
	var err error

	//
	//     ./apiserver redigo cla exp
	//

	var cla objectid.ID
	{
		cla = objectid.ID(arg[0])
	}

	var exp time.Time
	{
		exp = converter.StringToTime(arg[1])
	}

	// --------------------------------------------------------------------- //

	var pos *poststorage.Object
	{
		pos, err = r.seaPos(sto, cla)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err = sto.Post().DeleteExpiry([]*poststorage.Object{pos})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		pos.Expiry = exp
	}

	{
		err = sto.Post().CreateExpiry([]*poststorage.Object{pos})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err = sto.Post().UpdatePost([]*poststorage.Object{pos})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
