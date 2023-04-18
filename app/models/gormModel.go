package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type GormModel struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt *time.Time `json:"created_at",omitempty`
	UpdatedAt *time.Time `json:"updated_at",omitempty`
}
