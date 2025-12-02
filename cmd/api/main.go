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
	config := loadConfig()
	initLogger(config)

	app := setupApp()
	initDatabase()
	
	setupAuth(app, config)

	setupRoutes(app)
	setupMetrics(app)

	startServer(app)
}

func loadConfig() *domain.AppConfiguration {
	return domain.GetConfig()
}

func initLogger(config *domain.AppConfiguration) {
	logger.Init(config.LoggerDebug)
	logger.Log.Info("Starting server...")
}

func setupApp() *fiber.App {
	return fiber.New()
}

func initDatabase() {
	if err := postgres.InitDB(); err != nil {
		logger.Log.Error("Database connection error:", "error", err)
		panic("")
	}
}

func setupAuth(app *fiber.App, config *domain.AppConfiguration) *http.AuthHandler {
	jwtKey := []byte(config.Security.JWTKey)

	passwordHasher := security.NewBCryptPasswordHasher()
	userRepository := postgres.NewUserRepositoryPostgres(postgres.Conn)
	securityRepository := security.NewJwtSecurityRepository(jwtKey)
	authService := usecase.NewAuthService(userRepository, passwordHasher, securityRepository)
	authHandler := http.NewAuthHandler(authService)

	api := securityRepository.SecurizePath("/api", app)
	authHandler.RegisterSecuredRoutes(api)
	authHandler.RegisterRoutes(app)

	return authHandler
}

func setupRoutes(app *fiber.App) {
	app.Get("/swagger/*", fiberSwagger.HandlerDefault)
}

func setupMetrics(app *fiber.App) {
	prometheus := fiberprometheus.New("my_service")
	app.Use(prometheus.Middleware)
	prometheus.RegisterAt(app, "/metrics")
}

func startServer(app *fiber.App) {
	app.Listen(":3000")
}
