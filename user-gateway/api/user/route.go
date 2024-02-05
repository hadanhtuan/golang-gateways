package apiUser

import (
	pkg "github.com/hadanhtuan/go-sdk"
	"github.com/gin-gonic/gin"
)

func InitRoute(router *gin.RouterGroup, app *pkg.App) error {
	userHandler := app.Handler[app.Config.GRPC.UserServicePort].(*UserHandler)

	userGroup := router.Group("/user")

	userGroup.GET("/me", userHandler.GetUser)
	return nil
}
