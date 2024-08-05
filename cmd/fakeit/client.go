package fakeit

import (
	"fmt"
	"net"
	"net/http"

	"github.com/twitchtv/twirp"
	"github.com/uvio-network/apigocode/pkg/user"
	"github.com/uvio-network/apiserver/pkg/envvar"
)

type Client struct {
	User user.API
}

func NewClient(env envvar.Env) Client {
	return Client{
		User: user.NewAPIJSONClient(
			fmt.Sprintf("http://%s", net.JoinHostPort(env.HttpHost, env.HttpPort)),
			&http.Client{},
			twirp.WithClientPathPrefix(""),
		),
	}
}
