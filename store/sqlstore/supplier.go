package sqlstore

import (
	"github.com/goudai-projects/gd-ops/log"
	"github.com/goudai-projects/gd-ops/store"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlSupplierStores struct {
	user store.UserStore
}

type SqlSupplier struct {
	db     *gorm.DB
	dsn    string
	config *gorm.Config
	stores SqlSupplierStores
}

func NewSqlSupplier(dsn string) *SqlSupplier {
	supplier := &SqlSupplier{
		dsn:    dsn,
		config: &gorm.Config{},
	}
	supplier.initConnection()

	supplier.stores.user = newSqlUserStore(supplier)

	return supplier
}

func (supplier *SqlSupplier) initConnection() {
	db, err := gorm.Open(mysql.Open(supplier.dsn), supplier.config)
	if err != nil {
		log.Fatal("connect to database fail")
	}
	supplier.db = db
}

func (supplier *SqlSupplier) User() store.UserStore {
	return supplier.stores.user
}
