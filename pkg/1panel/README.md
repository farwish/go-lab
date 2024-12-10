# 1Panel

`仅支持 linux.`

## Development

https://1panel.cn/docs/dev_manual/dev_manual/

```bash
$ git clone https://github.com/1Panel-dev/1Panel.git

# frontend
$ cd 1Panel/frontend
$ npm run build:dev
$ npm i
$ npm run dev

# backend
$ cd 1Panel/cmd/server
$ go get
$ go run main.go
```

## 1pctl cli

https://github.com/1Panel-dev/installer/blob/main/1pctl

- status, start, stop, restart 命令是直接调用了 systemctl 命令操作 1panel.service
- uninstall 命令是停止及禁止自启动 1panel.service, 并删掉了 /opt/1panel /usr/local/bin/{1pctl,1panel} /etc/systemd/system/1panel.service
- user-info, listen-ip, version, update, reset, restore 命令是调用的构建好的 1panel 命令。

```bash
$ 1pctl

1Panel 控制脚本

Usage:
  ./1pctl [COMMAND] [ARGS...]
  ./1pctl --help

Commands:
  status              查看 1Panel 服务运行状态
  start               启动 1Panel 服务
  stop                停止 1Panel 服务
  restart             重启 1Panel 服务
  uninstall           卸载 1Panel 服务
  user-info           获取 1Panel 用户信息
  listen-ip           切换 1Panel 监听 IP
  version             查看 1Panel 版本信息
  update              修改 1Panel 系统信息
  reset               重置 1Panel 系统信息
  restore             恢复 1Panel 服务及数据
```

## 1panel cli

```bash
$ 1panel -h

1Panel ，一款现代化的 Linux 面板

Usage:
  1panel [flags]
  1panel [command]

Available Commands:
  app         应用相关命令
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  listen-ip   切换监听 IP
  reset       重置系统信息
  restore     回滚 1Panel 服务及数据
  update      修改面板信息
  user-info   获取面板信息
  version     获取系统版本信息

Flags:
  -h, --help   help for 1panel

Use "1panel [command] --help" for more information about a command.
```

- listen-ip, reset, update 是对 sqlite 的设置操作。
- user-info, version 是对 sqlite 的查询操作。
- restore 谨慎操作。


```bash
# cmd/server/cmd/root.go

var RootCmd = &cobra.Command{
	Use: "1panel",
	RunE: func(cmd *cobra.Command, args []string) error {
		server.Start()
		return nil
	},
}
```

直接运行 1panel 命令会启动 Gin-api 的 Http Server ，cron 管理器。 

Start 函数启动了多种服务，包括配置文件读取、国际化支持、日志记录、数据库连接、数据迁移、应用程序初始化、数据验证、会话管理、定时任务、业务逻辑初始化、钩子服务以及 HTTP/HTTPS 服务。这些服务共同构成了应用的运行环境，确保应用能够正常启动和运行。

其中 `Cron 调度器` 和 `API http server` 会在后台持续运行。
