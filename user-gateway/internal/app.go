package internal

import (
	"context"
	"fmt"
	"net/http"
	apiUser "user-gateway/api/user"
	"user-gateway/docs"
	"user-gateway/internal/middleware"
	pkg "github.com/hadanhtuan/go-sdk"
	grpcClient "github.com/hadanhtuan/go-sdk/client"
	. "github.com/hadanhtuan/go-sdk/common"
	userService "user-gateway/proto/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitGRPC(app *pkg.App) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app.Handler = map[string]interface{}{}

	userConn, err := grpcClient.NewGRPCClientServiceConn(ctx, app.Config.GRPC.UserServicePort)
	if err != nil {
		return err
	}
	defer userConn.Close()

	userServiceClient := userService.NewUserServiceClient(userConn)
	app.Handler[app.Config.GRPC.UserServicePort] = apiUser.NewHandler(userServiceClient)

	return nil
}

func InitRoute(app *pkg.App) error {
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

	basePath := router.Group(config.HttpServer.ApiPath)

	basePath.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, APIResponse{Status: http.StatusOK, Message: "Pong!"})
	})

	//Init Route
	apiUser.InitRoute(&router.RouterGroup, app)
	fmt.Println(config.HttpServer.SwaggerPath)

	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{config.HttpServer.TrustedDomain})

	gatewayRoute := fmt.Sprintf("%s:%s", config.HttpServer.TrustedDomain, config.HttpServer.AppPort)
	swaggerRoute := fmt.Sprintf("/%s/*any", config.HttpServer.SwaggerPath)

	// docs.SwaggerInfo.Host = gatewayRoute
	docs.SwaggerInfo.BasePath = "/" + config.HttpServer.ApiPath

	router.GET(swaggerRoute, ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(gatewayRoute)

	return nil
}
