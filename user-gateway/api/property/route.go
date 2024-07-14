package apiProperty

import (
	"github.com/gin-gonic/gin"
	"github.com/hadanhtuan/go-sdk"
)

func InitRoute(router *gin.RouterGroup, app *sdk.App) error {
	propertyController := app.Handler[app.Config.GRPC.PropertyServicePort].(*PropertyController)

	propertyGroup := router.Group("/property")
	// Property
	propertyGroup.POST("/get", propertyController.GetProperty)
	propertyGroup.POST("/count-by-status", propertyController.CountPropertyStatus)
	propertyGroup.POST("", propertyController.CreateProperty)
	propertyGroup.PATCH("", propertyController.UpdateProperty)
	propertyGroup.DELETE("/:propertyId", propertyController.DeleteProperty)

	// Booking
	propertyGroup.POST("/booking", propertyController.CreateBooking)
	propertyGroup.POST("/booking/get", propertyController.GetBooking)
	propertyGroup.POST("/booking/analyze", propertyController.AnalyzeBooking)
	propertyGroup.POST("/booking/count-by-status", propertyController.CountBookingStatus)

	// Review
	propertyGroup.POST("/review", propertyController.CreateReview)
	propertyGroup.POST("/review/get", propertyController.GetReview)
	propertyGroup.PATCH("/review", propertyController.UpdateReview)
	propertyGroup.DELETE("/review/:reviewId", propertyController.DeleteReview)

	// Amenity
	propertyGroup.POST("/amenity", propertyController.CreateAmenity)
	propertyGroup.POST("/amenity/get", propertyController.GetAmenity)
	propertyGroup.PATCH("/amenity", propertyController.UpdateAmenity)
	propertyGroup.DELETE("/amenity/:amenityId", propertyController.DeleteAmenity)

	// Favorite
	propertyGroup.POST("/favorite", propertyController.CreateFavorite)
	propertyGroup.POST("/favorite/get", propertyController.GetFavorite)
	propertyGroup.DELETE("/favorite/:favoriteId", propertyController.DeleteFavorite)

	return nil
}
