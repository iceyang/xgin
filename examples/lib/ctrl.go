package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iceyang/boom"
)

type Controller struct {
	service *Service
}

func (c *Controller) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.service.FindOne())
}

func (c *Controller) Error(ctx *gin.Context) {
	panic(boom.Boom(http.StatusBadRequest, "请求错误"))
}

func (c *Controller) InternalError(ctx *gin.Context) {
	c.service.PanicBoomError()
}

func NewController() *Controller {
	return &Controller{}
}
