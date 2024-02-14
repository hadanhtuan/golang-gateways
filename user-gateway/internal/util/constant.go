package util

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

var BODY_PAYLOAD = "BODY_PAYLOAD"

func ParseBody[T any](ctx *gin.Context) T {
	var payload T

	body, _ := ctx.Get(BODY_PAYLOAD)

	b, _ := json.Marshal(body)
	json.Unmarshal(b, &payload)

	return payload
}