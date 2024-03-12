package apiBooking

import (
	"context"
	"fmt"
	"strconv"
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

func (bc *BookingController) GetPropertyDetail(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var payload bookingProto.MsgGetProperty
	propertyId := c.Param("propertyId")

	if propertyId == "" {
		c.JSON(int(common.APIStatus.BadRequest), &sdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Invalid request",
		})
		return
	}
	payload.PropertyId = propertyId
	result, _ := bc.ServiceBookingClient.GetPropertyDetail(ctx, &payload)
	newResult := util.ConvertResult(result)

	c.JSON(int(newResult.Status), newResult)
}

func (bc *BookingController) GetAllProperty(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var payload bookingProto.MsgQueryProperty
	payload.Paginate = &sdk.Pagination{}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	
	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &sdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Invalid request",
		})
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &sdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Invalid request",
		})
		return
	}
	if page < 0 && page >= 100 {
		page = 1
	}

	if limit < 0 && limit >= 100 {
		limit = 10
	}
	payload.Paginate.Offset = int32((page - 1) * limit)
	payload.Paginate.Limit = int32(limit)
	result, _ := bc.ServiceBookingClient.GetAllProperty(ctx, &payload)
	newResult := util.ConvertResult(result)

	c.JSON(int(newResult.Status), newResult)
}

func (bc *BookingController) CreateProperty(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var payload bookingProto.MsgCreateProperty
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

	var payload bookingProto.MsgUpdateProperty
	err := c.BindJSON(&payload)
	propertyId := c.Param("propertyId")

	if propertyId == "" || err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &sdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Invalid request",
		})
		return
	}
	payload.PropertyId = propertyId
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
