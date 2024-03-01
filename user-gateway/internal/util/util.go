package util

import (
	"encoding/json"
	"fmt"
	"user-gateway/proto/sdk"

	"github.com/hadanhtuan/go-sdk/common"
)

func ConvertResult(payload *sdk.BaseResponse) *common.APIResponse {
	var data interface{}

	if payload == nil {
		return &common.APIResponse{
			Message: "Internal Server Error",
			Status:  common.APIStatus.ServerError,
		}
	}
	fmt.Printf("Payload: %v\n", payload)
	err := json.Unmarshal([]byte(payload.Data), &data)
	fmt.Printf("Payload1: %v\n", err)
	if err != nil {
		return &common.APIResponse{
			Message: payload.Message,
			Status:  common.APIStatus.ServerError,
		}
	}

	return &common.APIResponse{
		Message: payload.Message,
		Status:  payload.Status,
		Total:   payload.Total,
		Data:    data,
	}
}
