package socialmedia

import (
	"errors"
	"mygram294/app/models"

	"gorm.io/gorm"
)

func (r SocialMediaRepository) DeleteSocialMedia(db *gorm.DB, SocialMediaID uint) (err error) {
	dbreturn := db.Where("id = ?", SocialMediaID).Delete(&models.SocialMedia{})
	if dbreturn.Error != nil {
		return dbreturn.Error
	}
	// No RowsAffected (= 0), hacky solution: put if statement here and make errors
	if dbreturn.RowsAffected == 0 {
		return errors.New("data doesn't exist")
	}
	return nil
}
