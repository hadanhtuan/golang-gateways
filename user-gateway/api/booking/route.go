package apiBooking

import (
	// "user-gateway/internal/model"
	"github.com/gin-gonic/gin"
	pkg "github.com/hadanhtuan/go-sdk"
)

func InitRoute(router *gin.RouterGroup, app *pkg.App) error {
	bookingController := app.Handler[app.Config.GRPC.BookingServicePort].(*BookingController)

	bookingGroup := router.Group("/booking")
	// Property
	bookingGroup.GET("/:bookingId", bookingController.GetBookingDetail)
	bookingGroup.GET("/property/:propertyId", bookingController.GetPropertyDetail)
	bookingGroup.GET("/property", bookingController.GetAllProperty)
	bookingGroup.POST("/property", bookingController.CreateProperty)
	bookingGroup.PATCH("/property/:propertyId", bookingController.UpdateProperty)
	bookingGroup.DELETE("/property/:propertyId", bookingController.DeleteProperty)
	//Review
	bookingGroup.POST("/review", bookingController.CreateReview)
	bookingGroup.POST("/reviewGET", bookingController.GetReview)
	bookingGroup.PATCH("/review/:reviewId", bookingController.UpdateReview)
	bookingGroup.DELETE("/review/:reviewId", bookingController.DeleteReview)

	return nil
}
