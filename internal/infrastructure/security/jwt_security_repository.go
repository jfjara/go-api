package security

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/juanfran/mi-api/internal/domain"
	"github.com/juanfran/mi-api/internal/domain/model"
	"github.com/juanfran/mi-api/internal/infrastructure/logger"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

type JwtSecurityRepository struct {
	jwtKey []byte
}

func NewJwtSecurityRepository(jwtKey []byte) *JwtSecurityRepository {
	logger.Log.Info("Securized with JWT 'NewJwtSecurityRepository' setted...")
	return &JwtSecurityRepository{jwtKey: jwtKey}
}

func (j *JwtSecurityRepository) CreateToken(user *model.User) (string, error) {
	config := domain.GetConfig()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Minute * time.Duration(config.TokenLiveInMinutes)).Unix(),
	})
	result, error := token.SignedString(j.jwtKey)
	logger.Log.Debug("token generated", "token", result)
	return result, error
}

func (j *JwtSecurityRepository) SecurizePath(path string, app *fiber.App) fiber.Router {
	logger.Log.Debug("Securized path", "path", path)
	return app.Group(path, jwtware.New(jwtware.Config{
		SigningKey: j.jwtKey,
	}))
}
