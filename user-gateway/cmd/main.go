package main

import (
	"fmt"
	v1 "go-tripper-gateway/cmd/user-gateway/routes/v1"
	"go-tripper-gateway/docs"
	config "go-tripper-gateway/internal/configs"
	"go-tripper-gateway/internal/middlewares"
	user_gateway "go-tripper-gateway/services/user-gateway"
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	application, _ := user_gateway.InitApp()
	defer shutdownApp(application)

	appCnf := config.Config
	corsCnf := config.CorsConfig
	logger := middlewares.SetupLog(appCnf.LOG_PATH)

	router := gin.New()

	router.Use(cors.New(corsCnf))
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger, true))

	v1.InitRoute(router, application, appCnf)

	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{appCnf.TRUSTED_DOMAIN})

	gatewayRoute := fmt.Sprintf("%s:%s", appCnf.TRUSTED_DOMAIN, appCnf.AppPort)
	swaggerRoute := fmt.Sprintf("/%s/*any", appCnf.SWAGGER_PATH)

	// docs.SwaggerInfo.Host = gatewayRoute
	docs.SwaggerInfo.BasePath = "/" + appCnf.API_PATH

	router.GET(swaggerRoute, ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(gatewayRoute)
}

func shutdownApp(application *user_gateway.UserGatewayApplication) {
	application.CacheModule.Close()
}
