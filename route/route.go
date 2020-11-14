package route

import (
	"go_webserver/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Setup ...
func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
		return
	})

	return r
}
