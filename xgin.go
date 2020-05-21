package xgin

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

func (x *XGin) Run() error {
	x.Provide(Engine)
	return x.RunWithFunc(func(lc fx.Lifecycle, e *gin.Engine, router Router) {
		router.Route(e)
		port := 3000
		if x.config != nil {
			port = x.config.HttpPort
		}

		srv := &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: e,
		}

		lc.Append(fx.Hook{
			OnStart: func(context.Context) error {
				log.Println("Starting HTTP server.")
				go func() {
					if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
						log.Fatalf("listen: %s\n", err)
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				log.Println("Stopping HTTP server.")
				return srv.Shutdown(ctx)
			},
		})
	})
}

// Done returns a channel of signals to block on after starting the
// application. Applications listen for the SIGINT and SIGTERM signals; during
// development, users can send the application SIGTERM by pressing Ctrl-C in
// the same terminal as the running process.
func (x *XGin) Done() <-chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	return c
}

// Stop XGin Instance.
func (x *XGin) Stop() error {
	stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	x.app.Done()
	return x.app.Stop(stopCtx)
}
