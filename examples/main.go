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
	log.Println("[xgin] Running")

	<-x.Done()

	log.Println("[xgin] Stopping")
	if err := x.Stop(); err != nil {
		log.Fatalf("[xgin] Stop with error: %+v\n", err)
	}
	log.Println("[xgin] Stopped")
}
