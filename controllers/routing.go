package controllers

import (
	"github.com/gin-gonic/gin"

	"go-webapp-template/controllers/api"
)

func Routing(engine *gin.Engine) *gin.Engine {
	api.Routing(engine.Group("/api"))

	return engine
}
