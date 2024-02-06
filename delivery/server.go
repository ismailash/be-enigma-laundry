package delivery

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ismailash/be-enigma-laundry/config"
	"github.com/ismailash/be-enigma-laundry/delivery/controller"
	"github.com/ismailash/be-enigma-laundry/manager"
)

type Server struct {
	ucManager manager.UseCaseManager
	engine    *gin.Engine
	host      string
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	infraManager, err := manager.NewInfraManager(cfg)
	if err != nil {
		log.Fatal(err)
	}

	repoManager := manager.NewRepoManager(infraManager)
	ucManager := manager.NewUseCaseManager(repoManager)
	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		ucManager: ucManager,
		engine:    engine,
		host:      host,
	}
}

func (s *Server) setupControllers() {
	rg := s.engine.Group("/api/v1")
	controller.NewBillController(s.ucManager.NewBillUseCase(), rg).Route()
}

func (s *Server) Run() {
	s.setupControllers()
	if err := s.engine.Run(s.host); err != nil {
		log.Fatal("server can't run")
	}
}
