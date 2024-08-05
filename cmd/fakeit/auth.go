package fakeit

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"net/http"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/twitchtv/twirp"
	"github.com/xh3b4sd/tracer"
)

const (
	kid = "localhost-fake-kid"
)

func bytHan(byt []byte, typ string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{
			w.Header().Set("Content-Type", typ)
		}

		{
			_, err := w.Write(byt)
			if err != nil {
				tracer.Panic(err)
			}
		}
	}
}

func newCtx(key jwk.Key, sub string) context.Context {
	var err error

	clm := map[string]string{
		jwt.AudienceKey: "fake",
		jwt.IssuerKey:   "fake",
		jwt.SubjectKey:  sub,
	}

	var tok []byte
	{
		tok, err = newTok(key, clm)
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	var hdr http.Header
	{
		hdr = http.Header{
			"authorization": []string{"Bearer " + string(tok)},
		}
	}

	var ctx context.Context
	{
		ctx, err = twirp.WithHTTPRequestHeaders(context.Background(), hdr)
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	return ctx
}

func newTok(key jwk.Key, clm map[string]string) ([]byte, error) {
	var err error

	var tok jwt.Token
	{
		tok = jwt.New()
	}

	for k, v := range clm {
		err := tok.Set(k, v)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var sig []byte
	{
		sig, err = jwt.Sign(tok, jwa.ES256, key)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return sig, nil
}

func newKey() (jwk.Key, jwk.Set, error) {
	var err error

	var pri *ecdsa.PrivateKey
	{
		pri, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			return nil, nil, tracer.Mask(err)
		}
	}

	var key jwk.Key
	{
		key, err = jwk.New(pri)
		if err != nil {
			return nil, nil, tracer.Mask(err)
		}

		err = key.Set(jwk.KeyIDKey, kid)
		if err != nil {
			return nil, nil, tracer.Mask(err)
		}
	}

	var set jwk.Set
	{
		pub, err := jwk.New(pri.PublicKey)
		if err != nil {
			return nil, nil, tracer.Mask(err)
		}

		err = pub.Set(jwk.AlgorithmKey, jwa.ES256)
		if err != nil {
			return nil, nil, tracer.Mask(err)
		}

		err = pub.Set(jwk.KeyIDKey, kid)
		if err != nil {
			return nil, nil, tracer.Mask(err)
		}

		{
			set = jwk.NewSet()
			set.Add(pub)
		}
	}

	return key, set, nil
}

func srvJwk(set jwk.Set) {
	var err error

	var byt []byte
	{
		byt, err = json.Marshal(set)
		if err != nil {
			tracer.Panic(err)
		}
	}

	{
		http.HandleFunc("/jwk", bytHan(byt, "application/json"))
	}

	{
		err := http.ListenAndServe("localhost:7171", nil)
		if err != nil {
			tracer.Panic(err)
		}
	}
}
