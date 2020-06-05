package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-webapp-template/libs/apiError"
	serviceV1 "go-webapp-template/services/api/v1"
)

func getUsers(context *gin.Context) {
	response, err := serviceV1.GetUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, apiError.Info{
			Message: err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, response)
}
