package models

import (
	"encoding/json"

	"github.com/kemal576/go-pw-manager/utils"
)

type Response struct {
	Success bool
	Message string
}

func NewResponse(success bool, message string) Response {
	return Response{Success: success, Message: message}
}

func (r Response) ToJson() string {
	byte, err := json.Marshal(r)
	utils.ErrorCheck(err)
	return string(byte)
}
