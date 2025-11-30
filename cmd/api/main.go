package main

import (
	_ "github.com/juanfran/mi-api/docs"

	"github.com/juanfran/mi-api/internal/domain"
	"github.com/juanfran/mi-api/internal/infrastructure/http"
	"github.com/juanfran/mi-api/internal/infrastructure/logger"
	"github.com/juanfran/mi-api/internal/infrastructure/persistence"
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

	jwtKey := []byte(config.JWTKey)

	passwordHasher := security.NewBCryptPasswordHasher()
	userRepository := persistence.NewUserRepositoryMemory(passwordHasher)
	securityRepository := security.NewJwtSecurityRepository(jwtKey)
	authService := usecase.NewAuthService(userRepository, passwordHasher, securityRepository)
	authHandler := http.NewAuthHandler(authService)
	authHandler.RegisterRoutes(app)

	api := securityRepository.SecurizePath("/api", app)
	//TO DO: aqui habria que declarar los endpoints a definir que estan securizados. Crear un handler mejor para cada operacion
	authHandler.RegisterSecuredRoutes(api)
	app.Get("/swagger/*", fiberSwagger.HandlerDefault)
	app.Listen(":3000")
}
