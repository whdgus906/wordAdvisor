package main

import (
	"github.com/gin-gonic/gin"
	"github.com/whdgus906/wordAdvisor/config"
	"github.com/whdgus906/wordAdvisor/controllers"
)

func RouteMapping(g *gin.RouterGroup, cfg *config.Config) {
	ac := &controllers.ArticleController{Cfg: cfg}

	g.POST("/analysis", ac.Analysis)
}
