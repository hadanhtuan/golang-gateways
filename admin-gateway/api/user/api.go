package apiUser

import (
	"context"
	"encoding/json"
	"strings"
	"time"
	"admin-gateway/internal/util"
	userProto "admin-gateway/proto/user"

	"github.com/gin-gonic/gin"
	"github.com/hadanhtuan/go-sdk"
	"github.com/hadanhtuan/go-sdk/common"
)

type UserController struct {
	ServiceClient userProto.UserServiceClient
}

func NewUserController(serviceClient userProto.UserServiceClient) *UserController {
	return &UserController{ServiceClient: serviceClient}
}

func (uc *UserController) Login(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var payload userProto.MsgLogin
	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Error parsing body. Error detail " + err.Error(),
		})
		return
	}

	ipAddress := c.ClientIP()
	userAgent := c.Request.UserAgent()
	deviceId := sdk.HashDevice(ipAddress, userAgent)

	payload.UserAgent = userAgent
	payload.IpAddress = ipAddress
	payload.DeviceId = deviceId

	result, _ := uc.ServiceClient.Login(ctx, &payload)

	newResult := util.ConvertResult(result)
	c.JSON(int(newResult.Status), newResult)
}

func (uc *UserController) Register(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var payload userProto.MsgRegister
	err := c.BindJSON(&payload)

	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Error parsing body. Error detail " + err.Error(),
		})
		return
	}

	ipAddress := c.ClientIP()
	userAgent := c.Request.UserAgent()
	deviceId := sdk.HashDevice(ipAddress, userAgent)
	payload.UserAgent = userAgent
	payload.IpAddress = ipAddress
	payload.DeviceId = deviceId

	result, _ := uc.ServiceClient.Register(ctx, &payload)

	newResult := util.ConvertResult(result)
	c.JSON(int(newResult.Status), newResult)
}

func (uc *UserController) RefreshToken(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var payload userProto.MsgToken
	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(int(common.APIStatus.BadRequest), &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Error parsing body. Error detail " + err.Error(),
		})
		return
	}

	result, _ := uc.ServiceClient.RefreshToken(ctx, &payload)

	newResult := util.ConvertResult(result)
	c.JSON(int(newResult.Status), newResult)
}

func (uc *UserController) Logout(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var payload userProto.MsgID

	ginPayload, _ := c.Get(common.JWT_PAYLOAD)
	jwtPayload := ginPayload.(*common.JWTPayload)

	payload.Id = jwtPayload.LoginLogID //TODO: logout by login log ID

	result, _ := uc.ServiceClient.Logout(ctx, &payload)

	newResult := util.ConvertResult(result)
	c.JSON(int(newResult.Status), newResult)
}

func (uc *UserController) GetProfile(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var payload userProto.MsgID

	ginPayload, _ := c.Get(common.JWT_PAYLOAD)
	jwtPayload := ginPayload.(*common.JWTPayload)

	payload.Id = jwtPayload.UserID

	result, _ := uc.ServiceClient.GetProfile(ctx, &payload)

	newResult := util.ConvertResult(result)
	c.JSON(int(newResult.Status), newResult)
}

// TODO: middleware verify token
func (uc *UserController) AuthorizeRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		token := c.GetHeader("Authorization")

		res := &common.APIResponse{}

		if token == "" {
			res.Message = "Not found token"
			c.AbortWithStatusJSON(int(common.APIStatus.Unauthorized), res)
			return
		}

		jwt := ExtractJWTHeader(token)

		result, err := uc.ServiceClient.VerifyToken(ctx, &userProto.MsgToken{
			Token: jwt,
		})

		if result.Status == common.APIStatus.Unauthorized || err != nil {
			c.AbortWithStatusJSON(int(common.APIStatus.Unauthorized), result)
			return
		}

		var data common.JWTPayload
		json.Unmarshal([]byte(result.Data), &data)
		c.Set(common.JWT_PAYLOAD, &data)
		c.Next()
	}
}

func ExtractJWTHeader(token string) string {
	jwt := token
	authHeaderParts := strings.Split(token, " ")
	if len(authHeaderParts) == 2 {
		jwt = authHeaderParts[1]
	}
	return jwt
}
