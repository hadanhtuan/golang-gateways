package main

import (
	"user-gateway/internal"

	"github.com/hadanhtuan/go-sdk"
	"github.com/hadanhtuan/go-sdk/config"
	"github.com/stripe/stripe-go"
)

// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {

	c, _ := config.InitConfig(".")

	stripe.Key = c.Stripe.SecretKey
	app := sdk.App{
		Config: c,
	}

	internal.InitGRPC(&app)
	internal.InitRoute(&app)
}
