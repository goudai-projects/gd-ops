package store

import "github.com/goudai-projects/gd-ops/model"

type Store interface {
	User() UserStore
}

type UserStore interface {
	Save(user *model.User) (*model.User, *model.AppError)
	Update(user *model.User) (*model.UserUpdated, *model.AppError)
	UpdatePassword(userId, newPassword string) *model.AppError
	Get(id string) (*model.User, *model.AppError)
	GetAll() ([]*model.User, *model.AppError)
	SearchAll(search *model.UserSearch) ([]*model.User, *model.AppError)
	SearchAllPaged(search *model.UserSearch) ([]*model.User, int64, *model.AppError)
	CreateTableIfNotExists()
}
