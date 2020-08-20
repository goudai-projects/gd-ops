package model

import "time"

type Zone struct {
	Id          string `gorm:"primarykey"`
	Name        string
	CreatedTime time.Time
}
