package main

import (
	_ "github.com/juanfran/mi-api/docs"

	fiberprometheus "github.com/ansrivas/fiberprometheus/v2"
	"github.com/juanfran/mi-api/internal/domain"
	"github.com/juanfran/mi-api/internal/infrastructure/http"
	"github.com/juanfran/mi-api/internal/infrastructure/logger"
	"github.com/juanfran/mi-api/internal/infrastructure/persistence/postgres"
	"github.com/juanfran/mi-api/internal/infrastructure/security"
	"github.com/juanfran/mi-api/internal/usecase"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/gofiber/swagger"
)

// @title Example user register API with Go!
// @version 1.0
// @description API Example to login and create users. Credentials are generated with JWT security
// @host localhost:3000
// @BasePath /
func main() {
	config := domain.GetConfig()
	logger.Init(config.LoggerDebug)

	logger.Log.Info("Starting server...")
	app := fiber.New()
	jwtKey := []byte(config.Security.JWTKey)
	if err := postgres.InitDB(); err != nil {
        logger.Log.Error("Error conectando a la base de datos:", "error", err)
    }

	passwordHasher := security.NewBCryptPasswordHasher()
	userRepository := postgres.NewUserRepositoryPostgres(postgres.Conn)
	securityRepository := security.NewJwtSecurityRepository(jwtKey)
	authService := usecase.NewAuthService(userRepository, passwordHasher, securityRepository)
	authHandler := http.NewAuthHandler(authService)

	prometheus := fiberprometheus.New("my_service")
	app.Use(prometheus.Middleware)
	api := securityRepository.SecurizePath("/api", app)
	
	authHandler.RegisterSecuredRoutes(api)
	authHandler.RegisterRoutes(app)
	app.Get("/swagger/*", fiberSwagger.HandlerDefault)
	
	prometheus.RegisterAt(app, "/metrics")
	
	app.Listen(":3000")
}

