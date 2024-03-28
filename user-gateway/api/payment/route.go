package apiPayment

import (
	"github.com/gin-gonic/gin"
	pkg "github.com/hadanhtuan/go-sdk"
)

func InitRoute(router *gin.RouterGroup, app *pkg.App) error {
	paymentController := app.Handler[app.Config.GRPC.PaymentServicePort].(*PaymentController)

	paymentGroup := router.Group("/payment")
	// Payment
	paymentGroup.POST("/create-payment-intent", paymentController.CreatePaymentIntent)
	paymentGroup.POST("/hook", paymentController.HandleHook)

	return nil
}
