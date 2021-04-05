package main

// go:generate swagger generate spec

import (
	"os"

	handlers "github.com/DUNA-E-Commmerce/{{cookiecutter.app_name}}/http/handlers"
	middlewares "github.com/DUNA-E-Commmerce/{{cookiecutter.app_name}}/http/middlewares"
	repo "github.com/DUNA-E-Commmerce/{{cookiecutter.app_name}}/repository"
	services "github.com/DUNA-E-Commmerce/{{cookiecutter.app_name}}/services"
	storage "github.com/DUNA-E-Commmerce/{{cookiecutter.app_name}}/storage"
	utils "github.com/DUNA-E-Commmerce/{{cookiecutter.app_name}}/utils"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

// RefreshTokenResponse a response model when the authorization token is refreshed
//
// swagger:response refreshTokenResponse
type RefreshTokenResponse struct {
	// in: body
	Body struct {
		// Administrator's ID
		//
		// Required: true
		Code int `json:"code"`
		// Authorization Token
		//
		// Required: true
		Token string `json:"token"`
		// Token's expiration date
		//
		// Required: true
		Expire string `json:"expire"`
	}
}

// AppConfiguration configuration
type AppConfiguration struct {
	Config        *utils.Config
	MerchantRepo  repo.MerchantRepository
	UserRepo      repo.UserRepository
	CryptoService services.CryptoService
}

func setupApp() *AppConfiguration {
	// Load configuration
	cfg, err := utils.GetConfig()
	if err != nil {
		panic(err)
	}

	// Connect to database
	storage.Init(cfg)

	// get db driver instance
	db := storage.GetDB()
	// initialize top level dependencies
	return &AppConfiguration{
		Config:        cfg,
		MerchantRepo:  repo.NewMerchantRepository(db),
		UserRepo:      repo.NewUserRepository(db),
		CryptoService: services.NewCryptoService(cfg, repo.NewMerchantPrivateKeyRepository(db)),
	}
}

func main() {
	appConfig := setupApp()
	// Set server routes
	server := NewRouter(appConfig)
	// Run server
	server.Run(os.Getenv("GIN_PORT"))
}

// NewRouter creates all app's routes
func NewRouter(cfg *AppConfiguration) *gin.Engine {
	engine := gin.New()
	engine.RedirectTrailingSlash = true
	engine.Use(requestid.New(), middlewares.Logger(utils.Logger()), gin.Recovery())

	// healthcheck
	health := new(handlers.HealthController)
	engine.GET("/health", health.Status)
	engine.Use(cors.AllowAll())

	merchant := handlers.NewMerchantController(cfg.Config, cfg.MerchantRepo, cfg.CryptoService)
	user := handlers.NewUserController(cfg.UserRepo)

	// Public routes
	api := engine.Group("api")
	{
		v1 := api.Group("v1")
		{
			// Middleware
			authMiddleware := middlewares.AuthMiddleware()
		}
	}
	return engine
}
