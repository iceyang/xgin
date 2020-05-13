package xgin

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var (
	ErrRouterProviderError = errors.New("router provider error")
)

type XGin struct {
	fxOption fx.Option
	config   *Config
	err      error
}

func New() *XGin {
	return &XGin{
		fxOption: fx.Options(),
	}
}

func (x *XGin) Config(config *Config) {
	x.config = config
}

func (x *XGin) Provide(constructors ...interface{}) {
	x.fxOption = fx.Options(
		x.fxOption,
		fx.Provide(constructors...),
	)
}

func (x *XGin) Router(routerConstructor interface{}) error {
	t := reflect.TypeOf(routerConstructor)
	typeName := t.Out(0).String()
	if typeName != "xgin.Router" {
		x.err = ErrRouterProviderError
		return x.err
	}

	x.Provide(routerConstructor)
	return nil
}

func (x *XGin) Run() error {
	if x.err != nil {
		return x.err
	}
	app := fx.New(
		x.fxOption,
		fx.Provide(Engine),
		fx.Invoke(func(e *gin.Engine, router Router) {
			router.Route(e)
			port := 3000
			if x.config != nil {
				port = x.config.HttpPort
			}
			e.Run(fmt.Sprintf(":%d", port))
		}),
	)

	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		return err
	}

	<-app.Done()

	return nil
}
