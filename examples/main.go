package main

import (
	"log"

	"github.com/iceyang/xgin"
	"github.com/iceyang/xgin/examples/lib"
)

func main() {
	x := xgin.New()
	x.Config(&xgin.Config{
		HttpPort: 8888,
	})

	x.Router(lib.NewRouter)
	x.Provide(
		lib.NewService,
		lib.NewController,
	)

	if err := x.Run(); err != nil {
		log.Fatalf("[xgin] Start error: %s\n", err)
	}
}
