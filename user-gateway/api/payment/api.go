package apiPayment

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
	"user-gateway/internal/util"
	protoPayment "user-gateway/proto/payment"

	"github.com/gin-gonic/gin"
	"github.com/hadanhtuan/go-sdk/common"
	"github.com/hadanhtuan/go-sdk/config"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/v76/webhook"
)

type PaymentController struct {
	ServicePaymentClient protoPayment.PaymentServiceClient
}

func NewPaymentController(servicePaymentClient protoPayment.PaymentServiceClient) *PaymentController {
	return &PaymentController{ServicePaymentClient: servicePaymentClient}
}

func (pc *PaymentController) CreatePaymentIntent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var payload protoPayment.MsgCreatePaymentIntent
	err := c.BindJSON(&payload)

	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Error parsing body. Error detail " + err.Error(),
		})
		return
	}

	result, _ := pc.ServicePaymentClient.CreatePaymentIntent(ctx, &payload)
	newResult := util.ConvertResult(result)

	c.JSON(int(newResult.Status), newResult)
}

func (pc *PaymentController) HandleHook(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	sigHeader := c.GetHeader("Stripe-Signature")
	endpointSecret := config.AppConfig.Stripe.EndpointSecret

	payload, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Error parsing event. Error detail: " + err.Error(),
		})
		return
	}

	event, err := webhook.ConstructEventWithOptions(payload, sigHeader, endpointSecret, webhook.ConstructEventOptions{
		IgnoreAPIVersionMismatch: true,
	})

	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Error parsing event. Error detail: " + err.Error(),
		})
		return
	}

	switch event.Type {
	case "payment_intent.succeeded":
		var paymentIntent stripe.PaymentIntent
		err := json.Unmarshal(event.Data.Raw, &paymentIntent)

		if err != nil {
			c.JSON(int(common.APIStatus.BadRequest), common.APIResponse{
				Status:  common.APIStatus.BadRequest,
				Message: "Error parsing webhook. Error detail: " + err.Error(),
			})
			return
		}

		payload := protoPayment.MsgPaymentIntent{
			StripeId:      paymentIntent.ID,
			BookingId:     paymentIntent.Metadata["bookingId"],
			PropertyId:    paymentIntent.Metadata["propertyId"],
			UserId:        paymentIntent.Metadata["userId"],
			Amount:        paymentIntent.Amount,
			Currency:      paymentIntent.Currency,
			ReceiptEmail:  paymentIntent.ReceiptEmail,
			CanceledAt:    &paymentIntent.CanceledAt,
			PaymentMethod: paymentIntent.PaymentMethodTypes[0],
			Status:        string(paymentIntent.Status),
			// Event:   paymentIntent.Source.SourceObject.Type,
		}
		fmt.Println(payload.BookingId)
		result, _ := pc.ServicePaymentClient.HookPayment(ctx, &payload)
		newResult := util.ConvertResult(result)
		c.JSON(int(newResult.Status), newResult)
		return

	case "payment_intent.canceled":
		var paymentMethod stripe.PaymentMethod
		err := json.Unmarshal(event.Data.Raw, &paymentMethod)
		if err != nil {
			c.JSON(int(common.APIStatus.BadRequest), common.APIResponse{
				Status:  common.APIStatus.BadRequest,
				Message: "Error parsing webhook. Error detail: " + err.Error(),
			})
			return
		}
	default:
		log.Println("Unhandled event type. Event detail: " + event.Type)
	}

	c.JSON(int(common.APIStatus.Ok), nil)

}
