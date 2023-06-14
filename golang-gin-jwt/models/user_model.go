package models

import (
	"time"

	"github.com/google/uuid"
)

// User struct represents users table in darabase
type User struct {
	ID        *uuid.UUID `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string     `gorm:"type:varchar(255)" json:"name"`
	Email     string     `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password  string     `gorm:"->;<-;not null" json:"-"`
	Role      *string    `gorm:"type:varchar(50);default:'user';not null"`
	Provider  *string    `gorm:"type:varchar(50);default:'local';not null"`
	Photo     *string    `gorm:"not null;default:'default.png'"`
	Verified  *bool      `gorm:"not null;default:false"`
	CreatedAt *time.Time `gorm:"not null;default:now()"`
	UpdatedAt *time.Time `gorm:"not null;default:now()"`
}
