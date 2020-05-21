package xgin

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type XGin struct {
	app       *fx.App
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

	x.app = fx.New(
		x.providers,
		x.invokes,
	)

	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	return x.app.Start(startCtx)
}

func (x *XGin) Stop(ctx context.Context) error {
	return x.app.Stop(ctx)
}

func (x *XGin) Run() error {
	x.Provide(Engine)
	return x.RunWithFunc(func(e *gin.Engine, router Router) {
		router.Route(e)
		port := 3000
		if x.config != nil {
			port = x.config.HttpPort
		}

		srv := &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: e,
		}

		// Initializing the server in a goroutine so that
		// it won't block the graceful shutdown handling below
		func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()
	})
}

func (x *XGin) Done() <-chan os.Signal {
	return x.app.Done()
}
