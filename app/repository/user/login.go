package user

import (
	"mygram294/app/models"

	"gorm.io/gorm"
)

func (r UserRepository) Login(db *gorm.DB, username string) (user *models.User, err error) {
	err = db.Debug().Where("username = ?", username).Take(&user).Error
	return user, err
}
