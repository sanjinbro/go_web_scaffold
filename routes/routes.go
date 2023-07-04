package routes

import (
	"go_web_scaffold/logger"
	"go_web_scaffold/settings"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	r.GET("/detail", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":    settings.Conf.Name,
			"version": settings.Conf.Version,
		})
	})
	return r
}
