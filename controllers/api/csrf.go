package api

import (
	"github.com/gin-gonic/gin"
	"go-webapp-template/libs/apiError"
	"net/http"
)

const csrfHeaderName = "X-Requested-With"

func csrfMiddleware(context *gin.Context) {
	_, ok := context.Request.Header[csrfHeaderName]
	if !ok {
		context.JSON(http.StatusBadRequest, apiError.Info{
			Message: "リクエストエラー",
		})
		context.Abort()
		return
	}

	context.Next()
}
