package apiProperty

import (
	"github.com/gin-gonic/gin"
	pkg "github.com/hadanhtuan/go-sdk"
)

func InitRoute(router *gin.RouterGroup, app *pkg.App) error {
	propertyController := app.Handler[app.Config.GRPC.PropertyServicePort].(*PropertyController)

	propertyGroup := router.Group("/property")
	// Property
	propertyGroup.POST("/get", propertyController.GetProperty)
	propertyGroup.POST("", propertyController.CreateProperty)
	propertyGroup.PATCH("", propertyController.UpdateProperty)
	propertyGroup.DELETE("/:propertyId", propertyController.DeleteProperty)

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
