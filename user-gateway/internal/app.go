package internal

import (
	"fmt"
	"net/http"
	"user-gateway/internal/middleware"
	"github.com/hadanhtuan/go-sdk/common"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sdk "github.com/hadanhtuan/go-sdk"
	apiUser "user-gateway/api/user"
	userService "user-gateway/proto/user"
	grpcClient "github.com/hadanhtuan/go-sdk/client"
)

func InitGRPC(app *sdk.App) error {
	app.Handler = map[string]interface{}{}

	userServiceHost := fmt.Sprintf(
		"%s:%s",
		app.Config.GRPC.UserServiceHost,
		app.Config.GRPC.UserServicePort,
	)
	fmt.Println(userServiceHost)
	userConn, err := grpcClient.NewGRPCClientServiceConn(userServiceHost)
	if err != nil {
		return fmt.Errorf("Failed to connect to %s: %v", userServiceHost, err)
	}
	// TODO: Bug if defer in here: defer userConn.Close()

	userServiceClient := userService.NewUserServiceClient(userConn)
	app.Handler[app.Config.GRPC.UserServicePort] = apiUser.NewController(userServiceClient)


	fmt.Println("Server down")
	return nil
}

func InitRoute(app *sdk.App) error {
	config := app.Config

	if config.HttpServer.ENV != "develop" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()

	router.Use(cors.New(config.Cors))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.TimeoutMiddleware(config.HttpServer.RequestTimeoutPerSecond))

	//TODO: missing config rate limit, will do it in future
	fmt.Println(config.HttpServer.ApiPath)
	basePath := router.Group(config.HttpServer.ApiPath)

	basePath.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, common.APIResponse{Status: http.StatusOK, Message: "Pong!"})
	})

	//Init Route
	apiUser.InitRoute(basePath, app)
	fmt.Println(config.HttpServer.SwaggerPath)

	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{config.HttpServer.TrustedDomain})

	gatewayRoute := fmt.Sprintf("%s:%s", config.HttpServer.TrustedDomain, config.HttpServer.AppPort)

	router.Run(gatewayRoute)

	return nil
}
