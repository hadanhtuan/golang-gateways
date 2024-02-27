package apiUser

import (
	// "user-gateway/internal/model"
	"github.com/gin-gonic/gin"
	pkg "github.com/hadanhtuan/go-sdk"
)


func InitRoute(router *gin.RouterGroup, app *pkg.App) error {
	userController := app.Handler[app.Config.GRPC.UserServicePort].(*UserController)

	userGroup := router.Group("/user")

	userGroup.POST("/login", userController.Login)
	userGroup.POST("/register", userController.Register)
	return nil
}
