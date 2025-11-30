package dto

type RegisterRequest struct {
	Username       string `json:"username" validate:"required,min=3,max=32"`
	Password       string `json:"password" validate:"required,min=4,max=64"`
	RepeatPassword string `json:"repeat-password" validate:"required,min=4,max=64"`
	Name           string `json:"name" validate:"required,min=3,max=32"`
	Surname        string `json:"surname" validate:"required,min=3,max=32"`
	Genre          string `json:"genre" validate:"required,min=4,max=64"`
	Age            int    `json:"age" validate:"required"`
}
