package apiUser

import (
	"context"
	"time"
	"user-gateway/internal/util"
	userProto "user-gateway/proto/user"

	"github.com/gin-gonic/gin"
	"github.com/hadanhtuan/go-sdk"
	"github.com/hadanhtuan/go-sdk/common"
)

type UserController struct {
	ServiceClient userProto.UserServiceClient
}

func NewController(serviceClient userProto.UserServiceClient) *UserController {
	return &UserController{ServiceClient: serviceClient}
}

func (uc *UserController) Login(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var payload userProto.MsgLogin
	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), nil)
		return
	}

	ipAddress := c.ClientIP()
	userAgent := c.Request.UserAgent()
	deviceId := sdk.HashDevice(payload.Email, ipAddress, userAgent)

	result, _ := uc.ServiceClient.Login(ctx, &userProto.MsgLogin{
		Email:     payload.Email,
		Password:  payload.Password,
		UserAgent: userAgent,
		IpAddress: ipAddress,
		DeviceId:  deviceId,
	})

	newResult := util.ConvertResult(result)
	c.JSON(int(newResult.Status), newResult)
}

func (uc *UserController) Register(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var payload userProto.MsgRegister
	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), nil)
		return
	}

	ipAddress := c.ClientIP()
	userAgent := c.Request.UserAgent()
	deviceId := sdk.HashDevice(payload.Email, ipAddress, userAgent)


	result, _ := uc.ServiceClient.Register(ctx, &userProto.MsgRegister{
		Email:     payload.Email,
		Password:  payload.Password,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		UserAgent: userAgent,
		IpAddress: ipAddress,
		DeviceId:  deviceId,
	})

	newResult := util.ConvertResult(result)
	c.JSON(int(newResult.Status), newResult)
}
