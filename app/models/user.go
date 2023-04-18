package models

import (
	// "gorm.io/gorm"
)

type User struct {
	GormModel
	Username string `gorm:"not null;uniqueIndex" json:"username" form:"username" validate:"required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" validate:"required,email"`
	Password string `gorm:"not null" json:"password" form:"password" validate:"required,min=6"`
	Age      int    `gorm:"not null" json:"age" form:"age" validate:"required,gt=8"`
}
