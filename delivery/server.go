package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jutionck/interview-bootcamp-apps/config"
	"github.com/jutionck/interview-bootcamp-apps/delivery/controller"
	"github.com/jutionck/interview-bootcamp-apps/delivery/middleware"
	"github.com/jutionck/interview-bootcamp-apps/manager"
	"github.com/jutionck/interview-bootcamp-apps/usecase"
	"github.com/jutionck/interview-bootcamp-apps/utils/exception"
	"github.com/jutionck/interview-bootcamp-apps/utils/logger"
	"github.com/jutionck/interview-bootcamp-apps/utils/security"
)

type Server struct {
	useCaseManager manager.UseCaseManager
	authService    usecase.AuthUseCase
	jwtService     security.JwtSecurity
	loggerService  logger.MyLogger
	engine         *gin.Engine
	host           string
}

func (s *Server) Run() {
	s.setupControllers()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}

func (s *Server) setupControllers() {
	s.engine.Use(middleware.NewLogMiddleware(s.loggerService).LogRequest())
	authMiddleware := middleware.NewAuthMiddleware(s.jwtService)
	controller.NewInitController(s.engine, s.useCaseManager.UserUseCase())
	controller.NewUserController(s.engine, s.useCaseManager.UserUseCase(), authMiddleware)
	controller.NewRecruiterController(s.engine, s.useCaseManager.RecruiterUseCase(), authMiddleware)
	controller.NewInterviewerController(s.engine, s.useCaseManager.InterviewerUseCase(), authMiddleware)
	controller.NewAuthController(s.engine, s.authService)
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	exception.CheckErr(err)
	infraManager, _ := manager.NewInfraManager(cfg)
	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)
	jwtService := security.NewJwtSecurity(cfg.TokenConfig)
	loggerService := logger.NewMyLogger(cfg.FileConfig)
	authUseCase := usecase.NewAuthUseCase(useCaseManager.UserUseCase(), jwtService)
	engine := gin.Default()
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &Server{
		useCaseManager: useCaseManager,
		engine:         engine,
		host:           host,
		authService:    authUseCase,
		jwtService:     jwtService,
		loggerService:  loggerService,
	}
}
