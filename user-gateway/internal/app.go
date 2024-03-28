package internal

import (
	"fmt"
	"net/http"
	apiPayment "user-gateway/api/payment"
	apiProperty "user-gateway/api/property"
	apiSearch "user-gateway/api/search"
	apiUser "user-gateway/api/user"

	protoPayment "user-gateway/proto/payment"
	protoProperty "user-gateway/proto/property"
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

	userServiceUrl := fmt.Sprintf(
		"%s:%s",
		app.Config.GRPC.UserServiceHost,
		app.Config.GRPC.UserServicePort,
	)
	bookingServiceUrl := fmt.Sprintf(
		"%s:%s",
		app.Config.GRPC.PropertyServiceHost,
		app.Config.GRPC.PropertyServicePort,
	)
	searchServiceUrl := fmt.Sprintf(
		"%s:%s",
		app.Config.GRPC.SearchServiceHost,
		app.Config.GRPC.SearchServicePort,
	)
	paymentServiceHost := fmt.Sprintf(
		"%s:%s",
		app.Config.GRPC.PaymentServiceHost,
		app.Config.GRPC.PaymentServicePort,
	)

	userConn, err := grpcClient.NewGRPCClientConn(userServiceUrl)
	if err != nil {
		return fmt.Errorf("failed to connect to %s: %v", userServiceUrl, err)
	}
	propertyConn, err := grpcClient.NewGRPCClientConn(bookingServiceUrl)
	if err != nil {
		return fmt.Errorf("failed to connect to %s: %v", bookingServiceUrl, err)
	}
	searchConn, err := grpcClient.NewGRPCClientConn(searchServiceUrl)
	if err != nil {
		return fmt.Errorf("failed to connect to %s: %v", searchServiceUrl, err)
	}
	paymentConn, err := grpcClient.NewGRPCClientConn(paymentServiceHost)
	if err != nil {
		return fmt.Errorf("failed to connect to %s: %v", paymentServiceHost, err)
	}

	// TODO: Bug if defer in here: defer userConn.Close()
	// USER
	userServiceClient := protoUser.NewUserServiceClient(userConn)
	app.Handler[app.Config.GRPC.UserServicePort] = apiUser.NewUserController(userServiceClient)

	//BOOKING
	propertyServiceClient := protoProperty.NewPropertyServiceClient(propertyConn)
	app.Handler[app.Config.GRPC.PropertyServicePort] = apiProperty.NewPropertyController(propertyServiceClient)

	//SEARCH
	searchServiceClient := protoSearch.NewSearchServiceClient(searchConn)
	app.Handler[app.Config.GRPC.SearchServicePort] = apiSearch.NewSearchController(searchServiceClient)

	//PAYMENT
	paymentServiceClient := protoPayment.NewPaymentServiceClient(paymentConn)
	app.Handler[app.Config.GRPC.PaymentServicePort] = apiPayment.NewPaymentController(paymentServiceClient)

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
	apiProperty.InitRoute(basePath, app)

	//SEARCH
	apiSearch.InitRoute(basePath, app)

	//PAYMENT
	apiPayment.InitRoute(basePath, app)

	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{config.HttpServer.TrustedDomain})

	gatewayRoute := fmt.Sprintf("%s:%s", config.HttpServer.TrustedDomain, config.HttpServer.AppPort)

	router.Run(gatewayRoute)

	return nil
}
