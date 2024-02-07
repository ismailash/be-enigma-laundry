package delivery

import (
	"fmt"
	"github.com/ismailash/be-enigma-laundry/delivery/middleware"
	"github.com/ismailash/be-enigma-laundry/usecase"
	"github.com/ismailash/be-enigma-laundry/utils/common"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ismailash/be-enigma-laundry/config"
	"github.com/ismailash/be-enigma-laundry/delivery/controller"
	"github.com/ismailash/be-enigma-laundry/manager"
)

type Server struct {
	ucManager  manager.UseCaseManager
	auth       usecase.AuthUseCase
	engine     *gin.Engine
	host       string
	jwtService common.JwtToken
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
	jwtService := common.NewJwtToken(cfg.TokenConfig)

	return &Server{
		ucManager:  ucManager,
		engine:     engine,
		host:       host,
		auth:       usecase.NewAuthUseCase(ucManager.NewUserUseCase(), jwtService),
		jwtService: jwtService,
	}
}

func (s *Server) setupControllers() {
	rg := s.engine.Group("/api/v1")
	authMiddleware := middleware.NewAuthMiddleware(s.jwtService)
	controller.NewBillController(s.ucManager.NewBillUseCase(), rg, authMiddleware).Route()
	controller.NewAuthController(s.auth, rg, s.jwtService).Route()
}

func (s *Server) Run() {
	s.setupControllers()
	if err := s.engine.Run(s.host); err != nil {
		log.Fatal("server can't run")
	}
}
