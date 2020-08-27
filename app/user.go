package app

import (
	"context"
	"github.com/goudai-projects/gd-ops/log"
	"github.com/goudai-projects/gd-ops/model"
)

func (s *Server) CreateUser(ctx context.Context, user *model.User) (*model.User, *model.AppError) {
	log.Infof("Create User, username: %s", user.Username)
	return user, nil
}

func (s *Server) SearchUsersPage(ctx context.Context, searchModel *model.UserSearch) ([]*model.User, int64, *model.AppError) {

	return s.Store.User().SearchAllPaged(searchModel)
}

func (s *Server) GetAllUser(ctx context.Context) ([]*model.User, *model.AppError) {

	return s.Store.User().GetAll()
}

func (s *Server) GetUser(ctx context.Context, userId string) (*model.User, *model.AppError) {
	return s.Store.User().Get(userId)
}
