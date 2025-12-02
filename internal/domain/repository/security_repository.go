package repository

import (
	"github.com/gofiber/fiber/v2"
	"github.com/juanfran/mi-api/internal/domain/model"
)

type SecurityRepository interface {
	CreateToken(user *model.User) (string, error)
	SecurizePath(path string, app *fiber.App) fiber.Router
}
