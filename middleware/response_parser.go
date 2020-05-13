package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseParser(c *gin.Context) {
	c.Next()
	if c.Writer.Size() > 0 {
		return
	}
	switch c.Writer.Status() {
	case http.StatusNotFound:
		c.JSON(http.StatusNotFound, gin.H{
			"message": "找不到资源",
		})
	case http.StatusOK:
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}
