package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	"github.com/iceyang/boom"
)

func Recovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			if e, ok := err.(*boom.Error); ok {
				if e.Status() >= 500 {
					fmt.Println(e.ErrorStack())
				}
				c.AbortWithStatusJSON(e.Status(), gin.H{
					"message": e.Message(),
				})
				return
			}
			fmt.Println(errors.New(err).ErrorStack())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "服务出错，请稍后尝试",
			})
		}
	}()
	c.Next()
}
