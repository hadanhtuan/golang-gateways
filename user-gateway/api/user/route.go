package apiUser

import (
	"github.com/gin-gonic/gin"
	"github.com/hadanhtuan/go-sdk"
)

func InitRoute(router *gin.RouterGroup, app *sdk.App) error {
	userController := app.Handler[app.Config.GRPC.UserServicePort].(*UserController)

	userGroup := router.Group("/user")

	userGroup.POST("/login", userController.Login)
	userGroup.POST("/get", userController.GetUsers)
	userGroup.POST("/register", userController.Register)
	userGroup.PATCH("", userController.UpdateUser)
	userGroup.POST("/refresh-token", userController.RefreshToken)
	userGroup.GET("/logout", userController.AuthorizeRequest(), userController.Logout)
	userGroup.GET("/profile", userController.AuthorizeRequest(), userController.GetProfile)
	return nil
}
