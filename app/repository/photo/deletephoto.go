package photo

import (
	"errors"
	"mygram294/app/models"

	"gorm.io/gorm"
)

func (r PhotoRepository) DeletePhoto(db *gorm.DB, PhotoID uint) (err error) {
	dbreturn := db.Where("id = ?", PhotoID).Delete(&models.Photo{})
	if dbreturn.Error != nil {
		return dbreturn.Error
	}
	// No RowsAffected (= 0), hacky solution: put if statement here and make errors
	if dbreturn.RowsAffected == 0 {
		return errors.New("data doesn't exist")
	}
	return nil
}
