package sqlstore

import (
	"github.com/goudai-projects/gd-ops/store"
	"gorm.io/gorm"
)

type SqlStore interface {
	GetDB() *gorm.DB
	User() store.UserStore
}
