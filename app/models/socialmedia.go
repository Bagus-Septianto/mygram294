package models

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name" validate:"required"`
	SocialMediaURL string `gorm:"not null" json:"social_media_url" form:"social_media_url" validate:"required"`
	UserID         uint
}