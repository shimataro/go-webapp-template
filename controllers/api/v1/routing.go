package v1

import (
	"github.com/gin-gonic/gin"
)

func Routing(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	routerGroup.GET("/users", getUsers)

	return routerGroup
}
