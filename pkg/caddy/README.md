# custom caddy

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


Other important resources:

## Install

https://github.com/caddyserver/caddy/?tab=readme-ov-file#install

## for Development

https://github.com/caddyserver/caddy/?tab=readme-ov-file#for-development

## Architecture

https://caddyserver.com/docs/architecture

## Extending Caddy (module)

https://caddyserver.com/docs/extending-caddy

## Admin API

https://caddyserver.com/docs/api

## Config Adapters

(The Caddyfile is a built-in config adapter)

https://caddyserver.com/docs/config-adapters
