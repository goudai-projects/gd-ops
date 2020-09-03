package sqlstore

import (
	"errors"
	"github.com/goudai-projects/gd-ops/model"
	"github.com/goudai-projects/gd-ops/store"
	"gorm.io/gorm"
	"net/http"
)

type SqlClusterStore struct {
	SqlStore
}

func newSqlClusterStore(sqlStore SqlStore) store.ClusterStore {
	c := &SqlClusterStore{
		SqlStore: sqlStore,
	}
	c.CreateTableIfNotExists()
	return c
}

func (s *SqlClusterStore) Save(cluster *model.Cluster) (*model.Cluster, *model.AppError) {
	cluster.PreSave()
	if err := cluster.IsValid(); err != nil {
		return nil, err
	}
	if err := s.GetDB().Create(cluster).Error; err != nil {
		return nil, model.NewAppError("store.sql_cluster.save.app_error", nil, http.StatusInternalServerError)
	}
	return cluster, nil
}

func (s *SqlClusterStore) Update(cluster *model.Cluster) (*model.ClusterUpdated, *model.AppError) {
	cluster.PreUpdate()
	if err := cluster.IsValid(); err != nil {
		return nil, err
	}
	var old model.Cluster
	err := s.GetDB().Find(&old, cluster.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, model.NewAppError("store.sql_cluster.update.find.app_error", nil, http.StatusBadRequest)
	}
	if err := s.GetDB().Save(cluster).Error; err != nil {
		return nil, model.NewAppError("store.sql_cluster.update.updating.app_error", nil, http.StatusInternalServerError)
	}
	cluster.Sanitize(map[string]bool{})
	old.Sanitize(map[string]bool{})
	return &model.ClusterUpdated{
		Old: &old,
		New: cluster,
	}, nil
}

func (s *SqlClusterStore) Get(id string) (*model.Cluster, *model.AppError) {
	var cluster model.Cluster
	err := s.GetDB().Find(&cluster, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, model.NewAppError("store.sql_cluster.get.app_error", nil, http.StatusBadRequest)
	}
	return &cluster, nil
}

func (s SqlClusterStore) GetAll() ([]*model.Cluster, *model.AppError) {
	var users []*model.Cluster
	err := s.GetDB().Where("deleted_at == null").Find(&users).Error
	if err != nil {
		return nil, model.NewAppError("store.sql_cluster.get.app_error", nil, http.StatusBadRequest)
	}
	return users, nil
}

func (s SqlClusterStore) SearchAll(search *model.ClusterSearch) ([]*model.Cluster, *model.AppError) {
	panic("implement me")
}

func (s SqlClusterStore) SearchAllPaged(search *model.ClusterSearch) ([]*model.Cluster, int64, *model.AppError) {
	var users []*model.Cluster
	var total int64
	tx := s.GetDB().Where("1 = 1")
	tx.Limit(search.Size).Offset(search.Page * search.Size)
	tx.Find(&users).Count(&total)
	return users, total, nil
}

func (s SqlClusterStore) CreateTableIfNotExists() {
	_ = s.GetDB().AutoMigrate(&model.Cluster{})
}
