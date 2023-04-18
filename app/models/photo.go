package models

type Photo struct {
	GormModel
	Title    string `gorm:"not null" json:"title" form:"title" validate:"required"`
	Caption  string `json:"caption" form:"caption"`
	PhotoURL string `gorm:"not null" json:"photo_url" form:"photo_url" validate:"required"`
	UserID   uint
}