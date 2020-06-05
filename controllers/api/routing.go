package api

import (
	"github.com/gin-gonic/gin"

	"go-webapp-template/controllers/api/v1"
)

func Routing(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	routerGroup.Use(csrfMiddleware)
	v1.Routing(routerGroup.Group("/v1"))

	return routerGroup
}
