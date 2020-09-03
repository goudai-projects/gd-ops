package app

import (
	"github.com/goudai-projects/gd-ops/model"
	"time"
)

func (s *Server) ClusterList() ([]*model.Cluster, *model.AppError) {
	return s.Store.Cluster().GetAll()
}

func (s *Server) ClusterCreate(userId uint64, dto *model.ClusterCreateDto) (*model.Cluster, *model.AppError) {
	return s.Store.Cluster().Save(&model.Cluster{
		UserId:             userId,
		Name:               dto.Name,
		KubeConfigAsString: dto.KubeConfigAsString,
		EnvId:              dto.EnvId,
		ZoneId:             dto.ZoneId,
		CreatedTime:        time.Time{},
	})
}
