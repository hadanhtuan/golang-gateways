package util

import (
	"encoding/json"
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

	err := json.Unmarshal([]byte(payload.Data), &data)
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
