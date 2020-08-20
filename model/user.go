package model

import (
	"encoding/json"
	"gorm.io/gorm"
	"io"
	"time"
)

const (
	DEFAULT_LOCALE = "zh-CN"
)

type User struct {
	gorm.Model
	Username           string
	Password           string
	Email              string
	Nickname           string
	LastPasswordUpdate time.Time
	Locale             string
}

type UserUpdated struct {
	Old *User
	New *User
}

type UserSearch struct {
	Pageable
	Username string
}

func (u *User) ToJson() string {
	b, _ := json.Marshal(u)
	return string(b)
}

func (u *User) PreSave() {

}

func (u *User) PreUpdate() {

}

func (u *User) IsValid() *AppError {

	return nil
}

func (u *User) Sanitize(options map[string]bool) {

}

func UserFromJson(data io.Reader) *User {
	var user *User
	_ = json.NewDecoder(data).Decode(&user)
	return user
}
