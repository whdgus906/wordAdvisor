package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whdgus906/wordAdvisor/config"
	"github.com/whdgus906/wordAdvisor/middlewares"
)

func main() {
	cfg := &config.Config{}
	cfg.ReadCfg()

	app := gin.New()
	app.Use(middlewares.CORS())
	app.Use(gin.Recovery())
	app.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Page not found",
		})
	})

	RouteMapping(app.Group("/"), cfg)

	app.Run(":80")
}
