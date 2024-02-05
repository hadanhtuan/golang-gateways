package main

import (
	. "github.com/hadanhtuan/go-sdk"
	config "github.com/hadanhtuan/go-sdk/config"

	// logger "github.com/hadanhtuan/go-sdk/logger"
	. "user-gateway/internal"
)

// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	config, _ := config.InitConfig("")
	// log := logger.NewApiLogger(config)
	app := App{
		Config: config,
	}

	InitGRPC(&app)
	InitRoute(&app)

}
