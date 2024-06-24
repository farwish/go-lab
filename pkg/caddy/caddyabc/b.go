package caddyabc

import (
	"github.com/caddyserver/caddy/v2"
)

// B只是一个例子；可以是你自己的类型
type B struct {
	MyField string `json:"my_field,omitempty"`
	Number  int    `json:"number,omitempty"`
}

// 通过CaddyModule方法返回Caddy模块的信息
func (B) CaddyModule() caddy.ModuleInfo {
	caddy.Log().Info("into B CaddyModule()")
	return caddy.ModuleInfo{
		ID:  "caddy.abc.b",   // namespace is caddy.abc, module_name is b
		New: func() caddy.Module { return new(B) },
	}
}

func (b *B) Start() error {
	caddy.Log().Info("B is started")

	return nil
}


// Provision sets up the module.
func (g *B) Provision(ctx caddy.Context) error {
	// TODO: set up the module

	// ctx.Logger() is a *zap.Logger
	ctx.Logger().Info("into A Provision()")
	return nil
}

// Validate validates that the module has a usable config.
func (g B) Validate() error {
	// TODO: validate the module's setup
	return nil
}

// Interface guards
var (
	_ caddy.Provisioner           = (*B)(nil)
	_ caddy.Validator             = (*B)(nil)
)

