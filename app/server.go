package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goudai-projects/gd-ops/api/actuator"
	"github.com/goudai-projects/gd-ops/config"
	"github.com/goudai-projects/gd-ops/log"
	"github.com/goudai-projects/gd-ops/store"
	"github.com/goudai-projects/gd-ops/store/sqlstore"
	"gorm.io/gorm"
	"sync"
)

type Server struct {
	config                *config.Config
	sqlStore              *sqlstore.SqlSupplier
	Store                 store.Store
	router                *gin.Engine
	ServerInitializedOnce sync.Once
}

func NewServer(config *config.Config) *Server {
	s := &Server{
		config: config,
	}
	s.initRouter()
	s.sqlStore = sqlstore.NewSqlSupplier(config.Database)
	s.Store = s.sqlStore

	return s
}

func (s *Server) initRouter() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(JSONAppErrorHandler())
	// add routes
	actuator.Routes(router)

	s.router = router
}

func (s *Server) Start() {
	s.StartHttpServer()
}

func (s *Server) StartHttpServer() {
	err := s.router.Run(fmt.Sprintf("%s:%d", s.config.Server.Address, s.config.Server.Port))
	if err != nil {
		log.Error("http server start fail")
	}
}

func (s *Server) StopHttpServer() {

}

func (s *Server) Shutdown() {
	s.StopHttpServer()
}

func (s *Server) DB() *gorm.DB {
	return s.sqlStore.GetDB()
}

func (s *Server) Router() *gin.Engine {
	return s.router
}
