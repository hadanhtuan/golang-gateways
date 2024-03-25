package internal

import (
	"fmt"
	"net/http"
	apiBooking "user-gateway/api/booking"
	apiSearch "user-gateway/api/search"
	apiUser "user-gateway/api/user"

	protoBooking "user-gateway/proto/booking"
	protoSearch "user-gateway/proto/search"
	protoUser "user-gateway/proto/user"

	"user-gateway/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hadanhtuan/go-sdk"
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
	searchServiceHost := fmt.Sprintf(
		"%s:%s",
		app.Config.GRPC.SearchServiceHost,
		app.Config.GRPC.SearchServicePort,
	)

	userConn, err := grpcClient.NewGRPCClientConn(userServiceHost)
	if err != nil {
		return fmt.Errorf("failed to connect to %s: %v", userServiceHost, err)
	}
	bookingConn, err := grpcClient.NewGRPCClientConn(bookingServiceHost)
	if err != nil {
		return fmt.Errorf("failed to connect to %s: %v", bookingServiceHost, err)
	}
	searchConn, err := grpcClient.NewGRPCClientConn(searchServiceHost)
	if err != nil {
		return fmt.Errorf("failed to connect to %s: %v", searchServiceHost, err)
	}

	// TODO: Bug if defer in here: defer userConn.Close()
	// USER
	userServiceClient := protoUser.NewUserServiceClient(userConn)
	app.Handler[app.Config.GRPC.UserServicePort] = apiUser.NewUserController(userServiceClient)

	//BOOKING
	bookingServiceClient := protoBooking.NewBookingServiceClient(bookingConn)
	app.Handler[app.Config.GRPC.BookingServicePort] = apiBooking.NewBookingController(bookingServiceClient)

	//SEARCH
	searchServiceClient := protoSearch.NewSearchServiceClient(searchConn)
	app.Handler[app.Config.GRPC.SearchServicePort] = apiSearch.NewSearchController(searchServiceClient)

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

	//SEARCH
	apiSearch.InitRoute(basePath, app)

	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{config.HttpServer.TrustedDomain})

	gatewayRoute := fmt.Sprintf("%s:%s", config.HttpServer.TrustedDomain, config.HttpServer.AppPort)

	router.Run(gatewayRoute)

	return nil
}
