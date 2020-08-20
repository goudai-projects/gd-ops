package model

import (
	"encoding/json"
)

type AppError struct {
	Id         string `json:"id"`
	Message    string `json:"message"`
	RequestId  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"` // The http status code
	params     map[string]interface{}
}

func (er *AppError) Error() string {
	return er.Message
}

func (er *AppError) ToJson() string {
	b, _ := json.Marshal(er)
	return string(b)
}

func NewAppError(id string, params map[string]interface{}, status int) *AppError {
	ap := &AppError{}
	ap.Id = id
	ap.params = params
	ap.Message = id
	ap.StatusCode = status
	return ap
}

type Pageable struct {
	Page int
	Size int
}
