package apiUser

import (
	"github.com/gin-gonic/gin"
	pkg "github.com/hadanhtuan/go-sdk"
)

func InitRoute(router *gin.RouterGroup, app *pkg.App) error {
	userController := app.Handler[app.Config.GRPC.UserServicePort].(*UserController)

	userGroup := router.Group("/user")

	userGroup.POST("/login", userController.Login)
	userGroup.POST("/register", userController.Register)
	userGroup.POST("/refresh-token", userController.RefreshToken)
	userGroup.GET("/logout", userController.AuthorizeRequest(), userController.Logout)
	userGroup.GET("/profile", userController.AuthorizeRequest(), userController.GetProfile)
	return nil
}
