package apiUser

import (
	userService "user-gateway/proto/user"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	ServiceClient userService.UserServiceClient
}

func NewHandler(serviceClient userService.UserServiceClient) *UserHandler {
	return &UserHandler{ServiceClient: serviceClient}
}

// @Tags		Posts
// @Security	BearerAuth
// @Accept		json
// @Produce	json
// @Success	200
// @Failure	500
// @Param		id	path	string	true	"param"
// @Router		/post/{id} [delete]
func (pc *UserHandler) GetUser(c *gin.Context) {
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	// defer cancel()

	// res := ResponseDefault[bool]{}
	// id := c.Param("id")

	// headerJWT, _ := c.Get(constants.JWTPayload)
	// jwtPayload := headerJWT.(*JWTPayloadSchema)

	// result, err := pc.ServiceClient.DeletePost(ctx, &post_service.MessageDeletePost{
	// 	PostId: id,
	// 	UserId: jwtPayload.AccountId,
	// })

	// utils.FormatDefaultResponse(c, &res, result, err)
}
