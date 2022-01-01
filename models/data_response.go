package models

import (
	"encoding/json"

	"github.com/kemal576/go-pw-manager/utils"
)

type DataResponse struct {
	Success bool
	Message string
	Data    interface{}
}

func NewDataResponse(success bool, message string, data interface{}) DataResponse {
	return DataResponse{Success: success, Message: message, Data: data}
}

func (r DataResponse) ToJson() string {
	byte, err := json.Marshal(r)
	utils.ErrorCheck(err)
	return string(byte)
}
