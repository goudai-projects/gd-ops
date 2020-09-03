package sqlstore

import (
	"github.com/goudai-projects/gd-ops/config"
	"github.com/goudai-projects/gd-ops/log"
	"github.com/goudai-projects/gd-ops/store"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type SqlSupplierStores struct {
	user    store.UserStore
	cluster store.ClusterStore
}

type SqlSupplier struct {
	db       *gorm.DB
	settings *config.Database
	config   *gorm.Config
	stores   SqlSupplierStores
}

func NewSqlSupplier(settings config.Database) *SqlSupplier {
	supplier := &SqlSupplier{
		config: &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   settings.TablePrefix,
				SingularTable: true,
			},
		},
		settings: &settings,
	}
	supplier.initConnection()

	// init store
	supplier.stores.user = newSqlUserStore(supplier)
	supplier.stores.cluster = newSqlClusterStore(supplier)

	return supplier
}

func (supplier *SqlSupplier) initConnection() {
	db, err := gorm.Open(mysql.Open(supplier.settings.DSN), supplier.config)
	if err != nil {
		log.Fatal("connect to database fail")
	}
	supplier.db = db
}

func (supplier *SqlSupplier) User() store.UserStore {
	return supplier.stores.user
}

func (supplier *SqlSupplier) Cluster() store.ClusterStore {
	return supplier.stores.cluster
}

func (supplier *SqlSupplier) GetDB() *gorm.DB {
	return supplier.db
}
