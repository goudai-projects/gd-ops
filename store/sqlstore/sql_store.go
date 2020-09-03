package sqlstore

import (
	"gorm.io/gorm"
)

type SqlStore interface {
	GetDB() *gorm.DB
	//User() store.UserStore
	//Cluster() store.ClusterStore
}
