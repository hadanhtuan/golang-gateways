package apiBooking

import (
	"github.com/gin-gonic/gin"
	pkg "github.com/hadanhtuan/go-sdk"
)

func InitRoute(router *gin.RouterGroup, app *pkg.App) error {
	bookingController := app.Handler[app.Config.GRPC.BookingServicePort].(*BookingController)

	bookingGroup := router.Group("/booking")
	// Property
	bookingGroup.GET("/:bookingId", bookingController.GetBookingDetail)
	bookingGroup.POST("/property/get", bookingController.GetProperty)
	bookingGroup.POST("/property", bookingController.CreateProperty)
	bookingGroup.PATCH("/property", bookingController.UpdateProperty)
	bookingGroup.DELETE("/property/:propertyId", bookingController.DeleteProperty)

	// Review
	bookingGroup.POST("/review", bookingController.CreateReview)
	bookingGroup.POST("/review/get", bookingController.GetReview)
	bookingGroup.PATCH("/review", bookingController.UpdateReview)
	bookingGroup.DELETE("/review/:reviewId", bookingController.DeleteReview)

	// Amenity
	bookingGroup.POST("/amenity", bookingController.CreateAmenity)
	bookingGroup.POST("/amenity/get", bookingController.GetAmenity)
	bookingGroup.PATCH("/amenity", bookingController.UpdateAmenity)
	bookingGroup.DELETE("/amenity/:amenityId", bookingController.DeleteAmenity)

	return nil
}
