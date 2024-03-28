package main

import (
	"user-gateway/internal"
	"github.com/hadanhtuan/go-sdk"
	config "github.com/hadanhtuan/go-sdk/config"
	"github.com/stripe/stripe-go"
)

// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {

	stripe.Key = "sk_test_51LfT9zHxaiuet7b6VatpmqxiwYfU8sjbmeYLu8uaiVDWgvrJXsP9cHx5k2Zr4KWqRt5jjnAPHNTdA7ThXT15l2MQ00LuQrsY4s"
	config, _ := config.InitConfig("")
	app := sdk.App{
		Config: config,
	}

	internal.InitGRPC(&app)
	internal.InitRoute(&app)
}
