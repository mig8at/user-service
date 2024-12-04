package dto

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Nickname string `json:"nickname" validate:"required,alphanum,min=3,max=30"`
	Bio      string `json:"bio" validate:"max=500"`
	Avatar   string `json:"avatar" validate:"url"`
}

type CreateUser struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Nickname string `json:"nickname" validate:"required,alphanum,min=3,max=30"`
	Bio      string `json:"bio" validate:"max=500"`
	Avatar   string `json:"avatar" validate:"url"`
}

type UpdateUser struct {
	Name     string `json:"name" validate:"omitempty,min=2,max=100"`
	Nickname string `json:"nickname" validate:"omitempty,alphanum,min=3,max=30"`
	Bio      string `json:"bio" validate:"omitempty,max=500"`
	Avatar   string `json:"avatar" validate:"omitempty,url"`
}
