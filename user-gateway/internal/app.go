package internal

import (
	"fmt"
	"net/http"
	apiBooking "user-gateway/api/booking"
	apiUser "user-gateway/api/user"
	"user-gateway/internal/middleware"
	bookingService "user-gateway/proto/booking"
	userService "user-gateway/proto/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sdk "github.com/hadanhtuan/go-sdk"
	grpcClient "github.com/hadanhtuan/go-sdk/client"
	"github.com/hadanhtuan/go-sdk/common"
)

func InitGRPC(app *sdk.App) error {
	app.Handler = map[string]interface{}{}

	userServiceHost := fmt.Sprintf(
		"%s:%s",
		app.Config.GRPC.UserServiceHost,
		app.Config.GRPC.UserServicePort,
	)
	bookingServiceHost := fmt.Sprintf(
		"%s:%s",
		app.Config.GRPC.BookingServiceHost,
		app.Config.GRPC.BookingServicePort,
	)
	fmt.Println(userServiceHost)
	userConn, err := grpcClient.NewGRPCClientConn(userServiceHost)
	if err != nil {
		return fmt.Errorf("Failed to connect to %s: %v", userServiceHost, err)
	}
	bookingConn, err := grpcClient.NewGRPCClientConn(bookingServiceHost)
	if err != nil {
		return fmt.Errorf("Failed to connect to %s: %v", bookingServiceHost, err)
	}

	// TODO: Bug if defer in here: defer userConn.Close()
	// USER
	userServiceClient := userService.NewUserServiceClient(userConn)
	app.Handler[app.Config.GRPC.UserServicePort] = apiUser.NewUserController(userServiceClient)

	//BOOKING
	bookingServiceClient := bookingService.NewBookingServiceClient(bookingConn)
	app.Handler[app.Config.GRPC.BookingServicePort] = apiBooking.NewBookingController(bookingServiceClient)

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
	basePath := router.Group(config.HttpServer.ApiPath)

	basePath.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, common.APIResponse{Status: http.StatusOK, Message: "Pong!"})
	})

	//Init Route
	// USER
	apiUser.InitRoute(basePath, app)

	//BOOKING
	apiBooking.InitRoute(basePath, app)

	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{config.HttpServer.TrustedDomain})

	gatewayRoute := fmt.Sprintf("%s:%s", config.HttpServer.TrustedDomain, config.HttpServer.AppPort)

	router.Run(gatewayRoute)

	return nil
}
