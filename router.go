package xgin

import "github.com/gin-gonic/gin"

type Router interface {
	Route(*gin.Engine)
}
