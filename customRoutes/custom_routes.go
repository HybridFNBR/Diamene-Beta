package addons

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DO NOT EDIT THE FUNC NAME OR THE PACKAGE NAME OR IT WILL NOT WORK
func CustomRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"pong": "pong",
		})
	})
}
