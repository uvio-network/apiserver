package authmiddleware

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/twitchtv/twirp"
	"github.com/uvio-network/apiserver/pkg/server/context/subjectclaim"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type MiddlewareConfig struct {
	Aud string
	Iss string
	Log logger.Interface
	URL string
}

type Middleware struct {
	log logger.Interface
	mid *jwtmiddleware.JWTMiddleware
}

func NewMiddleware(c MiddlewareConfig) *Middleware {
	if c.Aud == "" {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Aud must not be empty", c)))
	}
	if c.Iss == "" {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Iss must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.URL == "" {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.URL must not be empty", c)))
	}

	var err error

	var fnc func(ctx context.Context) (interface{}, error)
	{
		fnc = jwks.NewCachingProvider(
			musUrl(c.URL),
			5*time.Minute,
			jwks.WithCustomJWKSURI(musUrl(c.URL)),
		).KeyFunc
	}

	var val *validator.Validator
	{
		val, err = validator.New(fnc, validator.ES256, c.Iss, []string{c.Aud})
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	var mid *jwtmiddleware.JWTMiddleware
	{
		mid = jwtmiddleware.New(
			val.ValidateToken,
			jwtmiddleware.WithCredentialsOptional(true),
			jwtmiddleware.WithValidateOnOptions(false),
			jwtmiddleware.WithErrorHandler(errHan(c.Log)),
		)
	}

	return &Middleware{
		log: c.Log,
		mid: mid,
	}
}

func (m *Middleware) Handler(h http.Handler) http.Handler {
	// CheckJWT extracts and validates the bearer access token provided with the
	// request's authorization header, if any. Any valid claims are put into the
	// request's context and can be accessed like shown below.
	return m.mid.CheckJWT(http.HandlerFunc(func(wri http.ResponseWriter, req *http.Request) {
		var ctx context.Context
		{
			ctx = req.Context()
		}

		{
			cla, _ := ctx.Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
			if cla != nil {
				req = req.Clone(subjectclaim.NewContext(ctx, cla.RegisteredClaims.Subject))
			}
		}

		{
			h.ServeHTTP(wri, req)
		}
	}))
}

func errHan(log logger.Interface) func(wri http.ResponseWriter, req *http.Request, err error) {
	return func(wri http.ResponseWriter, req *http.Request, err error) {
		{
			log.Log(
				context.Background(),
				"level", "error",
				"message", err.Error(),
				"path", req.URL.String(),
				"stack", tracer.Stack(err),
			)
		}

		{
			err = twirp.WriteError(wri, err)
			if err != nil {
				tracer.Panic(tracer.Mask(err))
			}
		}
	}
}

func musUrl(str string) *url.URL {
	u, e := url.Parse(str)
	if e != nil {
		tracer.Panic(tracer.Mask(e))
	}

	return u
}
