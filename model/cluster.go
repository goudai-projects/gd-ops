package model

import "time"

type ClusterConfig struct {
	Id                 string `gorm:"primarykey"`
	Name               string
	KubeConfigAsString string
	EnvId              string
	ZoneId             string
	CreatedTime        time.Time
}

type ClusterCreateDto struct {
	Name               string
	KubeConfigAsString string
	EnvId              string
	ZoneId             string
}
