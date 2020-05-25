package main

import (
	"fmt"

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

	x.Run()
}
