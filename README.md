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



Running redis stack locally.

```
docker run --rm --name redis-stack-apiserver -p 6379:6379 -p 8001:8001 redis/redis-stack:latest
```



[Twirp]: https://github.com/twitchtv/twirp
