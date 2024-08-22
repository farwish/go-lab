# custom caddy

## How to custom caddy

This is a custom caddy, with ourselves' module plugged in.

```
go run main.go list-modules
```

----------------------------------------------------------

Knowledge see `github.com/caddyserver/caddy/blob/master/cmd/caddy/main.go` annotation.
```
// There is no need to modify the Caddy source code to customize your builds. 
// You can easily build a custom Caddy with these simple steps:

//  1. Copy `cmd/caddy/main.go` into a new folder.
//  2. Edit the imports to include the modules you want plugged in
//  3. Run `go mod init caddy`. (`go mod tidy`)
//  4. Run `go install` or `go build` - you now have a custom binary!
```

This `github.com/caddyserver/caddy/blob/master/cmd/caddy/main.go` is native caddy with standard module plugged in.

---

## Cmd detail explain

1 To see all available commands:

`go run main.go`


2 Starts the Caddy process and blocks indefinitely:

`go run main.go run [-c <path> [-a <name>]] [--envfile <path>] [-e] [-r] [-w] [--pidfile <file>] [-h]`

实现见：cmd/commandfuncs.go:cmdRun()

```
Usage:
  caddy run [--config <path> [--adapter <name>]] [--envfile <path>] [--environ] [--resume] [--watch] [--pidfile <file>] [flags]

Flags:
  -a, --adapter string    Name of config adapter to apply
  -c, --config string     Configuration file
      --envfile strings   Environment file(s) to load
  -e, --environ           Print environment
  -h, --help              help for run
      --pidfile string    Path of file to which to write process ID
      --pingback string   Echo confirmation bytes to this address on success
  -r, --resume            Use saved config, if any (and prefer over --config file)
  -w, --watch             Watch config file for changes and reload it automatically
```

Caddy Windows环境变量示意：

`go run main.go run -e`

  caddy.HomeDir=C:\Users\Administrator

  caddy.AppDataDir=C:\Users\Administrator\AppData\Roaming\Caddy

  caddy.AppConfigDir=C:\Users\Administrator\AppData\Roaming\Caddy

  caddy.ConfigAutosavePath=C:\Users\Administrator\AppData\Roaming\Caddy\autosave.json

  caddy.Version=(devel)

---

## Caddy source-code read

1. First, Clone `github.com/farwish/caddy-printcode`

2. Second, Edit ourselves `cmd/caddy/caddyfile`

```
vi cmd/caddy/caddyfile
```

```
http://localhost:8888 {
  reverse_proxy
}
```

3. Third, Test code by `go run main.go run` to see output print.

---

## Caddy relate resources:

#### Install

https://github.com/caddyserver/caddy/?tab=readme-ov-file#install

#### for Development

https://github.com/caddyserver/caddy/?tab=readme-ov-file#for-development

#### Architecture

https://caddyserver.com/docs/architecture

#### Extending Caddy (module)

https://caddyserver.com/docs/extending-caddy

#### Admin API

https://caddyserver.com/docs/api

#### Config Adapters

(The Caddyfile is a built-in config adapter)

https://caddyserver.com/docs/config-adapters
