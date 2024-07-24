# apiserver

Golang based [Twirp] apiserver. Requires at least `go1.22.5`.

```
./apiserver -h
```
```
Golang based RPC apiserver.

Usage:
  apiserver [flags]
  apiserver [command]

Available Commands:
  daemon      Execute the long running process exposing RPC server handlers.
  version     Print version information of this command line tool.

Flags:
  -h, --help   help for apiserver

Use "apiserver [command] --help" for more information about a command.
```



### development

Running redis stack locally.

```
docker run --rm --name redis-stack-apiserver -p 6379:6379 -p 8001:8001 redis/redis-stack:latest
```

Running the apiserver locally.

```
./apiserver daemon
```
```
{ "time":"2024-07-24 10:07:46", "leve":"info", "mess":"server listening for calls", "addr":"127.0.0.1:7777", "call":"/Users/xh3b4sd/project/uvio-network/apiserver/pkg/server/server.go:86" }
{ "time":"2024-07-24 10:07:46", "leve":"info", "mess":"worker searching for tasks", "addr":"127.0.0.1:6379", "call":"/Users/xh3b4sd/project/uvio-network/apiserver/pkg/worker/worker.go:55" }
```

Calling the apiserver locally at `/` returns a simple health check response.

```
curl -s http://127.0.0.1:7777
```
```
OK
```

Calling the apiserver locally at `/version` returns the project's version information.

```
curl -s http://127.0.0.1:7777/version | jq .
```
```
{
  "arc": "arm64",
  "gos": "darwin",
  "sha": "n/a",
  "src": "https://github.com/uvio-network/apiserver",
  "tag": "n/a",
  "ver": "go1.22.5"
}
```


[Twirp]: https://github.com/twitchtv/twirp
