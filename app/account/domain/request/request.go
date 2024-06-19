package request

type CreateAccountBodyRequest struct {
	Username string `json:"username" gorm:"username" validate:"required,min=6"`
	Password string `json:"password" gorm:"password" validate:"required,min=6"`
}

type LoginBodyRequest struct {
	Username string `json:"username" gorm:"username" validate:"required,min=6"`
	Password string `json:"password" gorm:"password" validate:"required,min=6"`
}

type UpdatePasswordBodyRequest struct {
	OldPassword string `json:"old_password" validate:"required,min=6"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}
