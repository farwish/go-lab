package caddyabc

import (
	"github.com/caddyserver/caddy/v2"
)

func init() {
	caddy.Log().Info("init() caddyabc")
	caddy.RegisterModule(A{})
	caddy.RegisterModule(B{})
}

// A只是一个例子；可以是你自己的类型
type A struct {
	MyField string `json:"my_field,omitempty"`
	Number  int    `json:"number,omitempty"`
}

// 通过CaddyModule方法返回Caddy模块的信息
func (A) CaddyModule() caddy.ModuleInfo {
	caddy.Log().Info("into A CaddyModule()")
	return caddy.ModuleInfo{
		ID:  "caddy.abc.a",   // namespace is caddy.abc, module_name is a
		New: func() caddy.Module { return new(A) },
	}
}

func (a *A) Start() error {
	caddy.Log().Info("A is started")

	return nil
}

// Provision sets up the module.
func (g *A) Provision(ctx caddy.Context) error {
	// TODO: set up the module

	// ctx.Logger() is a *zap.Logger
	ctx.Logger().Info("into A Provision()")
	return nil
}

// Validate validates that the module has a usable config.
func (g A) Validate() error {
	// TODO: validate the module's setup
	return nil
}

// Interface guards
var (
	_ caddy.Provisioner           = (*A)(nil)
	_ caddy.Validator             = (*A)(nil)
)
