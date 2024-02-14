package apiUser

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	userProto "user-gateway/proto/user"
	"github.com/hadanhtuan/go-sdk/common"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	ServiceClient userProto.UserServiceClient
}

func NewController(serviceClient userProto.UserServiceClient) *UserController {
	return &UserController{ServiceClient: serviceClient}
}

func (uc *UserController) Login(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var payload userProto.MessageLogin
	err := c.BindJSON(&payload)
	if err != nil {
		return
	}

	fmt.Println(payload.Username)
	fmt.Println(payload.Password)

	result, _ := uc.ServiceClient.Login(ctx, &userProto.MessageLogin{
		Username: payload.Username,
		Password: payload.Password,
	})
	fmt.Println("can go here")
	fmt.Println(result)
	fmt.Println("can not go here")

	var data interface{}
	var newResult common.APIResponse
	err = json.Unmarshal([]byte(result.Data), &data)


	newResult.Status = result.Status
	newResult.Message = result.Message
	newResult.ErrorCode = result.ErrorCode
	newResult.Total = result.Total
	newResult.Data = data

	c.JSON(int(result.Status), newResult)
}
