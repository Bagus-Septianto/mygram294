package models

type Comment struct {
	GormModel
	UserID  uint
	PhotoID uint `gorm:"not null" json:"photo_id" form:"photo_id" validate:"number,required"`
	Message string `gorm:"not null" json:"message" form:"message" validate:"required"`
}