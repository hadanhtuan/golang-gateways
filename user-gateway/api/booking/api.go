package apiBooking

import (
	"context"
	"fmt"
	"reflect"
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
	var payload bookingProto.MsgGetPropertyRequest
	// err := c.BindJSON(&payload)
	// fmt.Printf("Payload: %+v\n", payload)
	// fmt.Printf("Error: %+v\n", err)
	propertyId := c.Param("propertyId")
	fmt.Println("result", propertyId)
	if propertyId == "" {
		c.JSON(int(common.APIStatus.BadRequest), &sdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Invalid request",
		})
		return
	}
	payload.PropertyId = propertyId
	result, _ := bc.ServiceBookingClient.GetPropertyDetail(ctx, &payload)
	fmt.Println("result", result)
	newResult := util.ConvertResult(result)
	fmt.Println("newresult", newResult)

	c.JSON(int(newResult.Status), newResult)
}
func (bc *BookingController) GetAllPropertyDetail(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload bookingProto.MsgGetAllPropertyRequest
	err := c.BindJSON(&payload)
	fmt.Printf("Payload: %+v\n", payload)
	fmt.Printf("Error: %+v\n", err)
	propertyId := c.Param("propertyId")
	if propertyId == "" || err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &sdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Invalid request",
		})
		return
	}
	// check payload.page and payload.limit is type int
	if reflect.TypeOf(payload.Page).Kind() != reflect.Int || reflect.TypeOf(payload.Limit).Kind() != reflect.Int {
		c.JSON(int(common.APIStatus.BadRequest), &sdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Invalid request",
		})
		return
	}
	if payload.Page < 0 && payload.Page >= 100 {
		payload.Page = 1
	}

	if payload.Limit < 0 && payload.Limit >= 100 {
		payload.Limit = 10
	}
	result, _ := bc.ServiceBookingClient.GetAllProperty(ctx, &payload)
	fmt.Println("result", result)
	newResult := util.ConvertResult(result)
	fmt.Println("newresult", newResult)

	c.JSON(int(newResult.Status), newResult)
}

func (bc *BookingController) CreateProperty(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload bookingProto.MsgCreatePropertyRequest
	err := c.BindJSON(&payload)
	fmt.Printf("Payload: %+v\n", payload)
	fmt.Printf("Error: %+v\n", err)

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
	var payload bookingProto.MsgUpdatePropertyRequest
	err := c.BindJSON(&payload)
	fmt.Printf("Payload: %+v\n", payload)
	fmt.Printf("Error: %+v\n", err)
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
	fmt.Println("result", result)
	newResult := util.ConvertResult(result)
	fmt.Println("newresult", newResult)

	c.JSON(int(newResult.Status), newResult)
}
func (bc *BookingController) DeleteProperty(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload bookingProto.MsgDeletePropertyRequest
	// err := c.BindJSON(&payload)
	// fmt.Printf("Payload: %+v\n", payload)
	// fmt.Printf("Error: %+v\n", err)
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
	fmt.Println("result", result)
	newResult := util.ConvertResult(result)
	fmt.Println("newresult", newResult)

	c.JSON(int(newResult.Status), newResult)
}

func (bc *BookingController) GetBookingDetail(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	//parse id in params to int64
	// id := c.Param("id")
	// id64, err := strconv.ParseInt(id, 10, 64)
	// if err != nil {
	// 	panic(err)
	// }
	id := c.Param("bookingId")
	bookingID, err := strconv.ParseInt(id, 10, 64)
	fmt.Println(err, bookingID)
	if err != nil {
		res := &common.APIResponse{
			Message: "Booking ID is invalid",
			Status:  common.APIStatus.BadRequest,
		}
		c.JSON(int(res.Status), res)
		return
	}
	result, _ := bc.ServiceBookingClient.GetBookingDetail(ctx, &bookingProto.MsgGetBookingRequest{
		BookingId: bookingID,
	})
	fmt.Println(result)
	newResult := util.ConvertResult(result)
	c.JSON(int(newResult.Status), newResult)
}
