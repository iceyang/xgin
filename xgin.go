package xgin

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type XGin struct {
	providers fx.Option
	invokes   fx.Option
	config    *Config
}

func New() *XGin {
	return &XGin{
		providers: fx.Options(),
		invokes:   fx.Options(),
	}
}

func (x *XGin) Config(config *Config) {
	x.config = config
}

func (x *XGin) Provide(constructors ...interface{}) {
	x.providers = fx.Options(
		x.providers,
		fx.Provide(constructors...),
	)
}

func (x *XGin) Invoke(funcs ...interface{}) {
	x.invokes = fx.Options(
		x.invokes,
		fx.Invoke(funcs...),
	)
}

func (x *XGin) RunWithFunc(fun interface{}) error {
	x.Invoke(fun)

	app := fx.New(
		x.providers,
		x.invokes,
	)

	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		return err
	}

	<-app.Done()

	return nil
}

func (x *XGin) Run() error {
	x.Provide(Engine)
	return x.RunWithFunc(func(e *gin.Engine, router Router) {
		router.Route(e)
		port := 3000
		if x.config != nil {
			port = x.config.HttpPort
		}
		e.Run(fmt.Sprintf(":%d", port))
	})
}
