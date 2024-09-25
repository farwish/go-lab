# frp

## Build binary how to

*See: dockerfiles/Dockerfile-for-frps*

```
$ make frps
```

## Command overview

```bash
$ go run cmd/frps/main.go cmd/frps/root.go -h
```

```
frps is the server of frp (https://github.com/fatedier/frp)

Usage:
  frps [flags]

Flags:
      --allow-ports string               allow ports
      --bind-addr string                 bind address (default "0.0.0.0")
  -p, --bind-port int                    bind port (default 7000)
  -c, --config string                    config file of frps
      --dashboard-addr string            dashboard address (default "0.0.0.0")
      --dashboard-port int               dashboard port
      --dashboard-pwd string             dashboard password (default "admin")
      --dashboard-tls-cert-file string   dashboard tls cert file
      --dashboard-tls-key-file string    dashboard tls key file
      --dashboard-tls-mode               if enable dashboard tls mode
      --dashboard-user string            dashboard user (default "admin")
      --disable-log-color                disable log color in console
      --enable-prometheus                enable prometheus dashboard
  -h, --help                             help for frps
      --kcp-bind-port int                kcp bind udp port
      --log-file string                  log file (default "console")
      --log-level string                 log level (default "info")
      --log-max-days int                 log max days (default 3)
      --max-ports-per-client int         max ports per client
      --proxy-bind-addr string           proxy bind address (default "0.0.0.0")
      --strict-config                    strict config parsing mode, unknown fields will cause errors (default true)
      --subdomain-host string            subdomain host
      --tls-only                         frps tls only
  -t, --token string                     auth token
  -v, --version                          version of frps
      --vhost-http-port int              vhost http port
      --vhost-http-timeout int           vhost http response header timeout (default 60)
      --vhost-https-port int             vhost https port
```

```bash
$ go run cmd/frps/main.go cmd/frps/root.go -v  # 0.60.0
```

## Development

*run*
```
# Server
$ go run cmd/frps/main.go cmd/frps/root.go -c frps.toml

# Client
$ go run cmd/frpc/main.go -c frpc.toml
```

## Reference

*服务端配置*
https://gofrp.org/zh-cn/docs/reference/server-configures/

*客户端配置*
https://gofrp.org/zh-cn/docs/reference/client-configures/

*代理配置*
https://gofrp.org/zh-cn/docs/reference/proxy/
