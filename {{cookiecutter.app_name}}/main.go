package main

// go:generate swagger generate spec

import (
	"os"

	"github.com/DUNA-E-Commmerce/{{cookiecutter.app_name}}/controllers"
	middlewares "github.com/DUNA-E-Commmerce/{{cookiecutter.app_name}}/middlewares"

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
	Config *utils.Config
}

func setupApp() *AppConfiguration {
	// Load configuration
	cfg, err := utils.GetConfig()
	if err != nil {
		panic(err)
	}

	// Connect to database
	storage.Init(cfg)

	// initialize top level dependencies
	return &AppConfiguration{
		Config: cfg,
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
	health := new(controllers.HealthController)
	engine.GET("/health", health.Status)
	engine.Use(cors.AllowAll())

	controllers := SetUpApp()

	// Public routes
	api := engine.Group("api")
	{
		v1 := api.Group("v1")
		{

			// User Router
			helloRoute := v1.Group("hello")
			{
				// Hello World
				helloRoute.GET("", controllers.helloController.SayHello)

				helloRoute.GET("all", controllers.helloController.GetAll)
			}
		}
	}
	return engine
}
