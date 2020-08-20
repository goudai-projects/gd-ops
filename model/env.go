package model

import "time"

type Env struct {
	Id          string `gorm:"primarykey"`
	Name        string
	ZoneId      string
	CreatedTime time.Time
}
