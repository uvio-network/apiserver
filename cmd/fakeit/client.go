package fakeit

import (
	"fmt"
	"net"
	"net/http"

	"github.com/twitchtv/twirp"
	"github.com/uvio-network/apigocode/pkg/post"
	"github.com/uvio-network/apigocode/pkg/user"
	"github.com/uvio-network/apigocode/pkg/vote"
	"github.com/uvio-network/apiserver/pkg/envvar"
)

type Client struct {
	Post post.API
	User user.API
	Vote vote.API
}

func NewClient(env envvar.Env) Client {
	var url string
	{
		url = fmt.Sprintf("http://%s", net.JoinHostPort(env.HttpHost, env.HttpPort))
	}

	var cli *http.Client
	{
		cli = &http.Client{}
	}

	var opt []twirp.ClientOption
	{
		opt = []twirp.ClientOption{
			twirp.WithClientPathPrefix(""),
		}
	}

	return Client{
		Post: post.NewAPIJSONClient(url, cli, opt...),
		User: user.NewAPIJSONClient(url, cli, opt...),
		Vote: vote.NewAPIJSONClient(url, cli, opt...),
	}
}
