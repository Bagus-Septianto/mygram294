package user

import (
	"mygram294/app/models"

	"gorm.io/gorm"
)

func (r UserRepository) Register(db *gorm.DB, user *models.User) (err error){
	err = db.Debug().Create(&user).Error
	return err
}
