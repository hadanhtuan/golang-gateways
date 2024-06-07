package apiSearch

import (
	"context"
	"fmt"
	"time"
	"user-gateway/internal/util"
	"user-gateway/proto/sdk"
	searchProto "user-gateway/proto/search"

	"github.com/gin-gonic/gin"
	"github.com/hadanhtuan/go-sdk/common"
)

type SearchController struct {
	ServiceSearchClient searchProto.SearchServiceClient
}

func NewSearchController(serviceSearchClient searchProto.SearchServiceClient) *SearchController {
	return &SearchController{ServiceSearchClient: serviceSearchClient}
}

func (bc *SearchController) Search(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var payload searchProto.MsgSearchProperty
	err := c.BindJSON(&payload)

	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &sdk.BaseResponse{
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

	result, _ := bc.ServiceSearchClient.SearchProperty(ctx, &payload)
	newResult := util.ConvertResult(result)

	c.JSON(int(newResult.Status), newResult)
}

func (bc *SearchController) GetNation(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var payload searchProto.MsgIP

	if payload.Paginate == nil {
		payload.Paginate = &sdk.Pagination{
			Offset: 0,
			Limit:  10,
		}
	}

	result, _ := bc.ServiceSearchClient.GetNation(ctx, &payload)
	fmt.Println(result.Status)
	newResult := util.ConvertResult(result)

	c.JSON(int(newResult.Status), newResult)
}
