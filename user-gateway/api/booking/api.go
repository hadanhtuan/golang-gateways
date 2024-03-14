package apiBooking

import (
	"context"
	"fmt"
	"time"
	"user-gateway/internal/util"
	bookingProto "user-gateway/proto/booking"
	"user-gateway/proto/sdk"

	"github.com/gin-gonic/gin"
	"github.com/hadanhtuan/go-sdk/common"
)

type BookingController struct {
	ServiceBookingClient bookingProto.BookingServiceClient
}

func NewBookingController(serviceBookingClient bookingProto.BookingServiceClient) *BookingController {
	return &BookingController{ServiceBookingClient: serviceBookingClient}
}

func (bc *BookingController) GetProperty(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var payload bookingProto.MsgQueryProperty
	payload.Paginate = &sdk.Pagination{
		Offset: 0,
		Limit:  10,
	}
	err := c.BindJSON(&payload)

	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &sdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Invalid request",
		})
		return
	}

	result, _ := bc.ServiceBookingClient.GetProperty(ctx, &payload)
	newResult := util.ConvertResult(result)

	c.JSON(int(newResult.Status), newResult)
}

func (bc *BookingController) CreateProperty(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var payload bookingProto.MsgProperty
	err := c.BindJSON(&payload)

	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &sdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Invalid request",
		})
		return
	}
	result, _ := bc.ServiceBookingClient.CreateProperty(ctx, &payload)
	newResult := util.ConvertResult(result)
	c.JSON(int(newResult.Status), newResult)
}

func (bc *BookingController) UpdateProperty(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var payload bookingProto.MsgProperty
	err := c.BindJSON(&payload)

	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &sdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Invalid request",
		})
		return
	}
	result, _ := bc.ServiceBookingClient.UpdateProperty(ctx, &payload)
	newResult := util.ConvertResult(result)

	c.JSON(int(newResult.Status), newResult)
}

func (bc *BookingController) DeleteProperty(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload bookingProto.MsgDeleteProperty
	propertyId := c.Param("propertyId")
	if propertyId == "" {
		c.JSON(int(common.APIStatus.BadRequest), &sdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Invalid request",
		})
		return
	}
	payload.PropertyId = propertyId
	result, _ := bc.ServiceBookingClient.DeleteProperty(ctx, &payload)
	newResult := util.ConvertResult(result)

	c.JSON(int(newResult.Status), newResult)
}

func (bc *BookingController) GetBookingDetail(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	id := c.Param("bookingId")

	result, _ := bc.ServiceBookingClient.GetBookingDetail(ctx, &bookingProto.MsgGetBooking{
		BookingId: id,
	})
	fmt.Println(result)
	newResult := util.ConvertResult(result)
	c.JSON(int(newResult.Status), newResult)
}

func (bc *BookingController) CreateReview(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload bookingProto.MsgCreateReview
	err := c.BindJSON(&payload)

	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &sdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Invalid request",
		})
		return
	}
	result, _ := bc.ServiceBookingClient.CreateReview(ctx, &payload)
	newResult := util.ConvertResult(result)
	c.JSON(int(newResult.Status), newResult)
}

func (bc *BookingController) UpdateReview(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload bookingProto.MsgUpdateReview
	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &sdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Invalid request",
		})
		return
	}
	result, _ := bc.ServiceBookingClient.UpdateReview(ctx, &payload)
	newResult := util.ConvertResult(result)

	c.JSON(int(newResult.Status), newResult)
}

func (bc *BookingController) DeleteReview(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload bookingProto.MsgDeleteReview
	reviewId := c.Param("reviewId")
	if reviewId == "" {
		c.JSON(int(common.APIStatus.BadRequest), &sdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Invalid request",
		})
		return
	}
	payload.ReviewId = reviewId
	result, _ := bc.ServiceBookingClient.DeleteReview(ctx, &payload)
	newResult := util.ConvertResult(result)

	c.JSON(int(newResult.Status), newResult)
}

func (bc *BookingController) GetReview(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload bookingProto.MessageQueryReview
	payload.Paginate = &sdk.Pagination{
		Offset: 0,
		Limit:  10,
	}
	err := c.BindJSON(&payload)

	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &sdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Invalid request",
		})
		return
	}

	result, _ := bc.ServiceBookingClient.GetReview(ctx, &payload)
	newResult := util.ConvertResult(result)
	c.JSON(int(newResult.Status), newResult)
}
