package db_mapper

import (
	"github.com/juanfran/mi-api/internal/domain/model"
	db_model "github.com/juanfran/mi-api/internal/infrastructure/persistence/db-model"
)

func ToDomain(userEntity *db_model.UserEntity) *model.User {
	if userEntity == nil {
		return nil
	}

	return &model.User{
		Username: userEntity.Username,
		Password: userEntity.Password,
		Name:     userEntity.Attributes.Name,
		Surname1: userEntity.Attributes.Surname1,
		Surname2: userEntity.Attributes.Surname2,
	}
}

func ToEntity(user *model.User) *db_model.UserEntity {
	if user == nil {
		return nil
	}

	entity := &db_model.UserEntity{
		Username: user.Username,
		Password: user.Password,
	}
	entity.Attributes.Name = user.Name
	entity.Attributes.Surname1 = user.Surname1
	entity.Attributes.Surname2 = user.Surname2

	return entity
}
