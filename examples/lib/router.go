package lib

import (
	"github.com/gin-gonic/gin"
	"github.com/iceyang/xgin"
)

type router struct {
	controller *Controller
}

func (r *router) Route(e *gin.Engine) {
	e.GET("/", r.controller.Get)
	e.GET("/errors", r.controller.Error)
	e.GET("/errors/internal", r.controller.InternalError)
}

func NewRouter(controller *Controller) xgin.Router {
	return &router{
		controller: controller,
	}
}
