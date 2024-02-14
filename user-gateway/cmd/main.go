package main

import (
	. "github.com/hadanhtuan/go-sdk"
	config "github.com/hadanhtuan/go-sdk/config"
	. "user-gateway/internal"
)

// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	config, _ := config.InitConfig("")
	app := App{
		Config: config,
	}

	InitGRPC(&app)
	InitRoute(&app)
}
