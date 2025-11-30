package domain

import "github.com/juanfran/mi-api/internal/domain/model"

type UserRepository interface {
	GetByUsername(username string) (*model.User, error)
}
