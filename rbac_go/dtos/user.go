package dtos

import "boilerplate-api/models"

// CreateUserRequestData Request body data to create user
type CreateUserRequestData struct {
	models.User
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

// GetUserResponse Dtos for User model
type GetUserResponse struct {
	Email    string      `json:"email"`
	FullName string      `json:"full_name"`
	Phone    string      `json:"phone"`
	Gender   string      `json:"gender"`
	Role     models.Role `json:"role"`
}
