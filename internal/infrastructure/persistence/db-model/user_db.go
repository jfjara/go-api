package db_model

import "time"

type UserEntity struct {
	ID         int
	Username   string
	Password   string
	Attributes struct {
		Name     string `json:"name"`
		Surname1 string `json:"surname1"`
		Surname2 string `json:"surname2"`
	} `json:"attributes"`
	CreatedAt time.Time
}
