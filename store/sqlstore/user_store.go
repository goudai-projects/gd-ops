package sqlstore

import (
	"errors"
	"github.com/goudai-projects/gd-ops/model"
	"github.com/goudai-projects/gd-ops/store"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type SqlUserStore struct {
	SqlStore
}

func newSqlUserStore(sqlStore SqlStore) store.UserStore {
	userStore := &SqlUserStore{
		SqlStore: sqlStore,
	}
	userStore.CreateTableIfNotExists()
	return userStore
}

func (s SqlUserStore) Save(user *model.User) (*model.User, *model.AppError) {
	user.PreSave()
	if err := user.IsValid(); err != nil {
		return nil, err
	}
	if err := s.GetDB().Create(user).Error; err != nil {
		return nil, model.NewAppError("store.sql_user.save.app_error", nil, http.StatusInternalServerError)
	}
	return user, nil
}

func (s SqlUserStore) Update(user *model.User) (*model.UserUpdated, *model.AppError) {
	user.PreUpdate()
	if err := user.IsValid(); err != nil {
		return nil, err
	}
	var oldUser model.User
	err := s.GetDB().Find(&oldUser, user.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, model.NewAppError("store.sql_user.update.find.app_error", nil, http.StatusBadRequest)
	}
	user.CreatedAt = oldUser.CreatedAt
	user.Password = oldUser.Password
	user.LastPasswordUpdate = oldUser.LastPasswordUpdate

	if err := s.GetDB().Save(user).Error; err != nil {
		return nil, model.NewAppError("store.sql_user.update.updating.app_error", nil, http.StatusInternalServerError)
	}
	user.Sanitize(map[string]bool{})
	oldUser.Sanitize(map[string]bool{})
	return &model.UserUpdated{
		Old: &oldUser,
		New: user,
	}, nil
}

func (s SqlUserStore) UpdatePassword(userId, newPassword string) *model.AppError {
	var user model.User
	err := s.GetDB().Find(&user, userId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.NewAppError("store.sql_user.update.find.app_error", nil, http.StatusBadRequest)
	}
	user.Password = newPassword
	user.LastPasswordUpdate = time.Now()
	if err := s.GetDB().Save(user).Error; err != nil {
		return model.NewAppError("store.sql_user.update.updating.app_error", nil, http.StatusInternalServerError)
	}
	return nil
}

func (s SqlUserStore) Get(id string) (*model.User, *model.AppError) {
	var user model.User
	err := s.GetDB().First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, model.NewAppError("store.sql_user.get.app_error", nil, http.StatusBadRequest)
	}
	return &user, nil
}

func (s SqlUserStore) GetAll() ([]*model.User, *model.AppError) {
	var users []*model.User
	err := s.GetDB().Where("deleted_at == null").Find(&users).Error
	if err != nil {
		return nil, model.NewAppError("store.sql_user.get.app_error", nil, http.StatusBadRequest)
	}
	return users, nil
}

func (s SqlUserStore) SearchAll(search *model.UserSearch) ([]*model.User, *model.AppError) {
	var users []*model.User
	tx := s.GetDB().Where("1 = 1")
	if username := search.Username; username != "" {
		tx = tx.Where("username like ?", "%"+username+"%")
	}
	tx.Find(&users)
	return users, nil
}

func (s SqlUserStore) SearchAllPaged(search *model.UserSearch) ([]*model.User, int64, *model.AppError) {
	var users []*model.User
	var total int64
	tx := s.GetDB().Where("1 = 1")
	if username := search.Username; username != "" {
		tx = tx.Where("username like ?", "%"+username+"%")
	}
	tx.Limit(search.Size).Offset(search.Page * search.Size)
	tx.Find(&users).Count(&total)
	return users, total, nil
}

func (s SqlUserStore) CreateTableIfNotExists() {
	_ = s.GetDB().AutoMigrate(&model.User{})
}
