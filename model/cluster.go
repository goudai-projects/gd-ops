package model

import (
	"time"
)

type Cluster struct {
	Id                 uint64    `gorm:"primarykey",required`
	UserId             uint64    `json:"userId",required`
	Name               string    `json:"name",required`
	KubeConfigAsString string    `json:"kubeConfigAsString",required`
	EnvId              string    `json:"envId",required`
	ZoneId             string    `json:"zoneId",required`
	CreatedTime        time.Time `json:"createdTime",required`
}

type ClusterCreateDto struct {
	Name               string `json:"name"`
	KubeConfigAsString string `json:"kubeConfigAsString"`
	EnvId              string `json:"envId"`
	ZoneId             string `json:"zoneId"`
}

type ClusterUpdated struct {
	Old *Cluster `json:"old"`
	New *Cluster `json:"new"`
}

type ClusterSearch struct {
	Pageable
}

func (u *Cluster) PreSave() {

}

func (u *Cluster) PreUpdate() {

}

func (u *Cluster) IsValid() *AppError {

	return nil
}

func (u *Cluster) Sanitize(options map[string]bool) {

}
