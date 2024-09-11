package usermiddleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/uvio-network/apiserver/pkg/server/context/subjectclaim"
	"github.com/uvio-network/apiserver/pkg/server/context/usercreated"
	"github.com/uvio-network/apiserver/pkg/server/context/userid"
	"github.com/uvio-network/apiserver/pkg/storage/userstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type MiddlewareConfig struct {
	Log logger.Interface
	Use userstorage.Interface
}

type Middleware struct {
	log logger.Interface
	use userstorage.Interface
}

func NewMiddleware(c MiddlewareConfig) *Middleware {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Use == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Use must not be empty", c)))
	}

	return &Middleware{
		log: c.Log,
		use: c.Use,
	}
}

func (m *Middleware) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error

		var ctx context.Context
		{
			ctx = r.Context()
		}

		// Lookup the OAuth subject claim if available. If no subject claim is
		// present we are dealing with an unauthenticated request and simply execute
		// the next handler, since data may be returned to anyone on the internet.
		var sub string
		{
			sub = subjectclaim.FromContext(ctx)
			if sub == "" {
				h.ServeHTTP(w, r)
				return
			}
		}

		// Try to find a user object for the given subject claim. If the given
		// subject claim causes an internal mapping error and the calling user
		// requested to be created, then simply forward the request to the server
		// handler responsible for user creation.
		var obj *userstorage.Object
		{
			obj, err = m.use.SearchSubject(sub)
			if r.URL.String() == "/user.API/Create" && errors.Is(err, userstorage.SubjectClaimMappingError) {
				// At this point we have a new user, which we flag in the requests
				// context, so that further components down the stack can act
				// accordingly.
				{
					ctx = usercreated.NewContext(ctx, true)
				}

				// Continue processing the request using the wrapped context.
				{
					h.ServeHTTP(w, r.Clone(ctx))
				}
			} else if err != nil {
				{
					m.werror(ctx, w, tracer.Mask(err))
				}
			} else {
				// At this point we have an authenticated user for which we set the
				// relevant user ID to the request context for further use down the
				// stack.
				{
					ctx = userid.NewContext(ctx, obj.ID)
				}

				// Continue processing the request using the wrapped context. The next
				// handler may execute another middleware or the RPC handler for the
				// actual business logic.
				{
					h.ServeHTTP(w, r.Clone(ctx))
				}
			}
		}
	})
}

func (m *Middleware) werror(ctx context.Context, wri http.ResponseWriter, err error) {
	e, o := err.(*tracer.Error)
	if o {
		m.log.Log(
			ctx,
			"level", "error",
			"message", e.Error(),
			"code", strconv.Itoa(http.StatusInternalServerError),
			"description", e.Desc,
			"docs", e.Docs,
			"kind", e.Kind,
			"stack", tracer.Stack(e),
		)
	} else {
		m.log.Log(
			ctx,
			"level", "error",
			"code", strconv.Itoa(http.StatusInternalServerError),
			"message", err.Error(),
			"stack", tracer.Stack(err),
		)
	}

	{
		wri.WriteHeader(http.StatusInternalServerError)
		_, _ = wri.Write([]byte(`{"message":"` + err.Error() + `"}`))
	}
}
