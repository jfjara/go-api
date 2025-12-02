package postgres

import (
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/juanfran/mi-api/internal/domain/model"
	db_mapper "github.com/juanfran/mi-api/internal/infrastructure/persistence/db-mapper"
	DBModel "github.com/juanfran/mi-api/internal/infrastructure/persistence/db-model"
)

type UserRepositoryPostgres struct {
	db *sql.DB
}

func NewUserRepositoryPostgres(db *sql.DB) *UserRepositoryPostgres {
	return &UserRepositoryPostgres{db: db}
}

func (r *UserRepositoryPostgres) GetByUsername(username string) (*model.User, error) {
	var user DBModel.UserEntity
	var attributesJSON []byte

	query := `SELECT id, username, password, attributes FROM users WHERE username=$1`
	row := r.db.QueryRow(query, username)
	if err := row.Scan(&user.ID, &user.Username, &user.Password, &attributesJSON); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrUserNotFound
		}
		return nil, err
	}

	if err := json.Unmarshal(attributesJSON, &user.Attributes); err != nil {
		return nil, err
	}

	userDomain := db_mapper.ToDomain(&user)
	return userDomain, nil
}
