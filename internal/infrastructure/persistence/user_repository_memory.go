package persistence

import (
	"errors"

	"github.com/juanfran/mi-api/internal/domain/model"
	"github.com/juanfran/mi-api/internal/domain/repository"
	"github.com/juanfran/mi-api/internal/infrastructure/logger"
)

type UserRepositoryMemory struct {
	users map[string]model.User
}

func NewUserRepositoryMemory(hasher repository.PasswordHasher) *UserRepositoryMemory {
	logger.Log.Info("User Repository 'UserRepositoryMemory' setted...")
	hash, err := hasher.Hash("1234")
	if err != nil {
		panic("error generando el hash de la contraseña inicial")
	}

	return &UserRepositoryMemory{
		users: map[string]model.User{
			"jfjara": {Username: "jfjara", Password: hash},
		},
	}
}

func (r *UserRepositoryMemory) GetByUsername(username string) (*model.User, error) {
	user, ok := r.users[username]
	if !ok {
		return nil, errors.New("usuario no encontrado")
	}
	return &user, nil
}
