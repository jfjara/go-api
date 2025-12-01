package security

import (
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/juanfran/mi-api/internal/domain"
	"github.com/juanfran/mi-api/internal/domain/model"
	"github.com/juanfran/mi-api/internal/infrastructure/logger"
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

    claims := &jwt.RegisteredClaims{
        Subject:   user.Username,
        ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(config.Security.TokenLiveInHours))),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    result, err := token.SignedString(j.jwtKey)
    if err != nil {
        logger.Log.Error("error signing token", "error", err)
        return "", err
    }

    logger.Log.Debug("token generated", "token", result)
    return result, nil
}


func (j *JwtSecurityRepository) SecurizePath(path string, app *fiber.App) fiber.Router {
	logger.Log.Debug("Securized path", "path", path)
	return app.Group(path, jwtware.New(jwtware.Config{
		SigningKey: j.jwtKey,
	}))
}
