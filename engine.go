package xgin

import (
	"github.com/gin-gonic/gin"
	"github.com/iceyang/xgin/middleware"
)

func Engine() *gin.Engine {
	e := gin.New()

	e.Use(middleware.Recovery)
	e.Use(middleware.ResponseParser)

	return e
}
