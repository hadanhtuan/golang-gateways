package apiBooking

import (
	"context"
	"fmt"
	"time"
	"user-gateway/internal/util"
	bookingProto "user-gateway/proto/booking"

	"github.com/gin-gonic/gin"
)

type BookingController struct {
	ServiceBookingClient bookingProto.BookingServiceClient
}

func NewBookingController(serviceBookingClient bookingProto.BookingServiceClient) *BookingController {
	return &BookingController{ServiceBookingClient: serviceBookingClient}
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
	// id := c.Param("id")

	result, _ := bc.ServiceBookingClient.GetBookingDetail(ctx, &bookingProto.MsgGetBookingRequest{
		BookingId: 1,
	})
	fmt.Println(result)
	newResult := util.ConvertResult(result)
	c.JSON(int(newResult.Status), newResult)
}
