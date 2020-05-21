package main

import (
	"fmt"
	"log"

	"github.com/iceyang/xgin"
	"github.com/iceyang/xgin/examples/lib"
)

func main() {
	x := xgin.New()
	x.Config(&xgin.Config{
		HttpPort: 8888,
	})

	x.Provide(
		lib.NewService,
		lib.NewController,
		lib.NewRouter,
	)

	x.Invoke(func() {
		fmt.Println("I'm a invoke function")
	})

	if err := x.Run(); err != nil {
		log.Fatalf("[xgin] Start with error: %+v\n", err)
	}
}
