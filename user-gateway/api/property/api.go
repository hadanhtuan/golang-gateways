package apiProperty

import (
	"context"
	"fmt"
	"time"
	"user-gateway/internal/util"
	propertyProto "user-gateway/proto/property"
	"user-gateway/proto/sdk"

	"github.com/gin-gonic/gin"
	"github.com/hadanhtuan/go-sdk/common"
)

type PropertyController struct {
	ServicePropertyClient propertyProto.PropertyServiceClient
}

func NewPropertyController(servicePropertyClient propertyProto.PropertyServiceClient) *PropertyController {
	return &PropertyController{ServicePropertyClient: servicePropertyClient}
}

func (bc *PropertyController) GetProperty(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var payload propertyProto.MsgQueryProperty
	err := c.BindJSON(&payload)

	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Error parsing body. Error detail " + err.Error(),
		})
		return
	}

	if payload.Paginate == nil {
		payload.Paginate = &sdk.Pagination{
			Offset: 0,
			Limit:  10,
		}
	}

	result, _ := bc.ServicePropertyClient.GetProperty(ctx, &payload)
	newResult := util.ConvertResult(result)

	c.JSON(int(newResult.Status), newResult)
}

func (bc *PropertyController) CreateProperty(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var payload propertyProto.MsgProperty
	err := c.BindJSON(&payload)

	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Error parsing body. Error detail " + err.Error(),
		})
		return
	}
	result, _ := bc.ServicePropertyClient.CreateProperty(ctx, &payload)
	newResult := util.ConvertResult(result)
	c.JSON(int(newResult.Status), newResult)
}

func (bc *PropertyController) CountPropertyStatus(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var payload propertyProto.MsgProperty
	err := c.BindJSON(&payload)

	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Error parsing body. Error detail " + err.Error(),
		})
		return
	}
	result, _ := bc.ServicePropertyClient.CountPropertyStatus(ctx, &payload)
	newResult := util.ConvertResult(result)
	c.JSON(int(newResult.Status), newResult)
}

func (bc *PropertyController) UpdateProperty(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var payload propertyProto.MsgProperty
	err := c.BindJSON(&payload)

	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Error parsing body. Error detail " + err.Error(),
		})
		return
	}
	result, _ := bc.ServicePropertyClient.UpdateProperty(ctx, &payload)
	newResult := util.ConvertResult(result)

	c.JSON(int(newResult.Status), newResult)
}

func (bc *PropertyController) DeleteProperty(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload propertyProto.MsgDeleteProperty
	propertyId := c.Param("propertyId")
	if propertyId == "" {
		c.JSON(int(common.APIStatus.BadRequest), &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Invalid request",
		})
		return
	}
	payload.PropertyId = propertyId
	result, _ := bc.ServicePropertyClient.DeleteProperty(ctx, &payload)
	newResult := util.ConvertResult(result)

	c.JSON(int(newResult.Status), newResult)
}

func (bc *PropertyController) CreateReview(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload propertyProto.MsgCreateReview
	err := c.BindJSON(&payload)

	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Error parsing body. Error detail " + err.Error(),
		})
		return
	}
	result, _ := bc.ServicePropertyClient.CreateReview(ctx, &payload)
	newResult := util.ConvertResult(result)
	c.JSON(int(newResult.Status), newResult)
}

func (bc *PropertyController) UpdateReview(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload propertyProto.MsgUpdateReview
	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Error parsing body. Error detail " + err.Error(),
		})
		return
	}
	result, _ := bc.ServicePropertyClient.UpdateReview(ctx, &payload)
	newResult := util.ConvertResult(result)

	c.JSON(int(newResult.Status), newResult)
}

func (bc *PropertyController) DeleteReview(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload propertyProto.MsgDeleteReview
	reviewId := c.Param("reviewId")
	if reviewId == "" {
		c.JSON(int(common.APIStatus.BadRequest), &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Invalid request",
		})
		return
	}
	payload.ReviewId = reviewId
	result, _ := bc.ServicePropertyClient.DeleteReview(ctx, &payload)
	newResult := util.ConvertResult(result)

	c.JSON(int(newResult.Status), newResult)
}

func (bc *PropertyController) GetReview(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload propertyProto.MsgQueryReview
	err := c.BindJSON(&payload)

	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Error parsing body. Error detail " + err.Error(),
		})
		return
	}

	if payload.Paginate == nil {
		payload.Paginate = &sdk.Pagination{
			Offset: 0,
			Limit:  10,
		}
	}

	result, _ := bc.ServicePropertyClient.GetReview(ctx, &payload)
	newResult := util.ConvertResult(result)
	c.JSON(int(newResult.Status), newResult)
}

func (bc *PropertyController) CreateAmenity(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload propertyProto.MsgAmenity
	err := c.BindJSON(&payload)

	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Error parsing body. Error detail " + err.Error(),
		})
		return
	}
	result, _ := bc.ServicePropertyClient.CreateAmenity(ctx, &payload)
	newResult := util.ConvertResult(result)
	c.JSON(int(newResult.Status), newResult)
}

func (bc *PropertyController) UpdateAmenity(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload propertyProto.MsgAmenity
	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Error parsing body. Error detail " + err.Error(),
		})
		return
	}
	result, _ := bc.ServicePropertyClient.UpdateAmenity(ctx, &payload)
	newResult := util.ConvertResult(result)

	c.JSON(int(newResult.Status), newResult)
}

func (bc *PropertyController) DeleteAmenity(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload propertyProto.MsgId
	amenityId := c.Param("amenityId")

	if amenityId == "" {
		c.JSON(int(common.APIStatus.BadRequest), &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Invalid request",
		})
		return
	}
	payload.Id = amenityId
	result, _ := bc.ServicePropertyClient.DeleteAmenity(ctx, &payload)
	newResult := util.ConvertResult(result)

	c.JSON(int(newResult.Status), newResult)
}

func (bc *PropertyController) GetAmenity(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload propertyProto.MsgQueryAmenity
	err := c.BindJSON(&payload)

	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Error parsing body. Error detail " + err.Error(),
		})
		return
	}

	if payload.Paginate == nil {
		payload.Paginate = &sdk.Pagination{
			Offset: 0,
			Limit:  10,
		}
	}

	result, _ := bc.ServicePropertyClient.GetAmenity(ctx, &payload)
	newResult := util.ConvertResult(result)
	c.JSON(int(newResult.Status), newResult)
}

func (bc *PropertyController) CreateFavorite(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload propertyProto.MsgFavorite
	err := c.BindJSON(&payload)

	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Error parsing body. Error detail " + err.Error(),
		})
		return
	}
	result, _ := bc.ServicePropertyClient.CreateFavorite(ctx, &payload)
	newResult := util.ConvertResult(result)
	c.JSON(int(newResult.Status), newResult)
}

func (bc *PropertyController) DeleteFavorite(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload propertyProto.MsgId
	favoriteId := c.Param("favoriteId")

	if favoriteId == "" {
		c.JSON(int(common.APIStatus.BadRequest), &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Invalid request",
		})
		return
	}
	payload.Id = favoriteId
	result, _ := bc.ServicePropertyClient.DeleteFavorite(ctx, &payload)
	newResult := util.ConvertResult(result)

	c.JSON(int(newResult.Status), newResult)
}

func (bc *PropertyController) GetFavorite(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload propertyProto.MsgQueryFavorite
	err := c.BindJSON(&payload)

	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Error parsing body. Error detail " + err.Error(),
		})
		return
	}

	if payload.Paginate == nil {
		payload.Paginate = &sdk.Pagination{
			Offset: 0,
			Limit:  10,
		}
	}

	result, _ := bc.ServicePropertyClient.GetFavorite(ctx, &payload)
	newResult := util.ConvertResult(result)
	c.JSON(int(newResult.Status), newResult)
}

func (bc *PropertyController) CreateBooking(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload propertyProto.MsgBooking
	err := c.BindJSON(&payload)

	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Error parsing body. Error detail " + err.Error(),
		})
		return
	}
	result, _ := bc.ServicePropertyClient.CreateBooking(ctx, &payload)
	newResult := util.ConvertResult(result)
	c.JSON(int(newResult.Status), newResult)
}

func (bc *PropertyController) GetBooking(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload propertyProto.MsgQueryBooking
	err := c.BindJSON(&payload)

	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Error parsing body. Error detail " + err.Error(),
		})
		return
	}

	if payload.Paginate == nil {
		payload.Paginate = &sdk.Pagination{
			Offset: 0,
			Limit:  10,
		}
	}

	result, _ := bc.ServicePropertyClient.GetBooking(ctx, &payload)
	newResult := util.ConvertResult(result)
	c.JSON(int(newResult.Status), newResult)
}

func (bc *PropertyController) AnalyzeBooking(c *gin.Context) {
	fmt.Println(5)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload propertyProto.MsgQueryBooking

	result, _ := bc.ServicePropertyClient.AnalyzeBooking(ctx, &payload)
	newResult := util.ConvertResult(result)
	c.JSON(int(newResult.Status), newResult)
}

func (bc *PropertyController) CountBookingStatus(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var payload propertyProto.MsgBooking
	err := c.BindJSON(&payload)

	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Error parsing body. Error detail " + err.Error(),
		})
		return
	}

	result, _ := bc.ServicePropertyClient.CountBookingStatus(ctx, &payload)
	newResult := util.ConvertResult(result)
	c.JSON(int(newResult.Status), newResult)
}
