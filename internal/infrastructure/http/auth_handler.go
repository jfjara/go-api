package http

import (
	dto "github.com/juanfran/mi-api/internal/infrastructure/http/dto"
	"github.com/juanfran/mi-api/internal/infrastructure/logger"
	"github.com/juanfran/mi-api/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service *usecase.AuthService
}

func NewAuthHandler(s *usecase.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) RegisterRoutes(app *fiber.App) {
	app.Post("/login", h.Login)
}

func (h *AuthHandler) RegisterSecuredRoutes(api fiber.Router) {

	api.Get("/profile", func(c *fiber.Ctx) error {
		user := c.Locals("user")
		return c.JSON(fiber.Map{
			"message": "Bienvenido a tu perfil protegido",
			"token":   user,
		})
	})
}

// Login godoc
// @Summary User login
// @Description Authenticate an user and return a JWT Token
// @Tags auth
// @Accept json
// @Produce json
// @Param login body dto.LoginRequest true "Access credentials"
// @Success 200 {object} dto.LoginResponse
// @Failure 401 {object} dto.LoginResponse
// @Router /login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {

	var loginRequest dto.LoginRequest

	if err := c.BodyParser(&loginRequest); err != nil {
		logger.Log.Error("Payload not valid", "error", err)
		return c.Status(400).JSON(fiber.Map{"error": "Payload not valid"})
	}

	loginResponse, err := h.service.Login(loginRequest)
	if err != nil {
		logger.Log.Error("Login failed!", "error", err)
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(loginResponse)
}
