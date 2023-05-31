package dtos

import (
	"time"

	model "github.com/bj-budhathoki/learn-go/golang-gin-jwt/models"
	"github.com/google/uuid"
)

type SignUpInput struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,min=8"`
	Photo           string `json:"photo"`
}

type SignInInput struct {
	Email    string `json:"email"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FilterUserRecord(user *model.User) UserResponse {
	return UserResponse{
		ID:        *user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      *user.Role,
		Photo:     *user.Photo,
		Provider:  *user.Provider,
		CreatedAt: *user.CreatedAt,
		UpdatedAt: *user.UpdatedAt,
	}
}
