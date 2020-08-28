package model

import (
	"encoding/json"
	"github.com/goudai-projects/gd-ops/utils"
	"gorm.io/gorm"
	"time"
)

type AppError struct {
	Id         string `json:"id"`
	Message    string `json:"message"`
	RequestId  string `json:"request_id,omitempty"`
	StatusCode int    `json:"code,omitempty"` // The http status code
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

func (er *AppError) Localization(lang ...string) {
	er.Message = utils.Translate(lang, er.Id, er.params)
}

type Pageable struct {
	Page int `form:"page"`
	Size int `form:"size"`
}

type Page struct {
	Page    int         `json:"page"`
	Size    int         `json:"size"`
	Total   int64       `json:"total"`
	Content interface{} `json:"content"`
}

func NewPage(pageable *Pageable, total int64, content interface{}) *Page {
	return &Page{
		Page:    pageable.Page,
		Size:    pageable.Size,
		Total:   total,
		Content: content,
	}
}

type Model struct {
	Id        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
