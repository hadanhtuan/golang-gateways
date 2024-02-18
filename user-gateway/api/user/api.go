package apiUser

import (
	"context"
	"fmt"
	"time"
	"user-gateway/internal/util"
	userProto "user-gateway/proto/user"

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

	newResult := util.ConvertResult2(result)
	c.JSON(int(newResult.Status), newResult)
}
