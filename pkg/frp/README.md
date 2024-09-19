## frp

#### Development

*run*
```
# Server
$ go run cmd/frps/main.go cmd/frps/root.go -c frps.toml

# Client
$ go run cmd/frpc/main.go -c frpc.toml
```

*frps.toml*
```
# 服务端监听地址，用于接收 frpc 的连接，默认监听 0.0.0.0。
bindAddr = "0.0.0.0"
bindPort = 7000 # 服务端监听端口，默认值为 7000。

# 默认为 127.0.0.1，如果需要公网访问 Dashboard，需要修改为 0.0.0.0。
webServer.addr = "0.0.0.0"
webServer.port = 7500

# dashboard 用户名密码，可选，默认为空
webServer.user = "admin"
webServer.password = "admin"
```

*frpc.toml*
```
serverAddr = "127.0.0.1"
serverPort = 7000

[[proxies]]
name = "test-tcp"
type = "tcp"
localIP = "127.0.0.1"
localPort = 22
remotePort = 6000
```

#### Build binary

*See: dockerfiles/Dockerfile-for-frps*

```
$ make frps
```

#### Reference

*服务端配置*
https://gofrp.org/zh-cn/docs/reference/server-configures/

*客户端配置*
https://gofrp.org/zh-cn/docs/reference/client-configures/

*代理配置*
https://gofrp.org/zh-cn/docs/reference/proxy/
