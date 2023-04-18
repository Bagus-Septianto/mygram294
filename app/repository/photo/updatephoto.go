package photo

import (
	"errors"
	"mygram294/app/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r PhotoRepository) UpdatePhoto(db *gorm.DB, p *models.Photo, PhotoID uint) (photo *models.Photo, err error) {
	dbReturn := db.Model(&p).Clauses(clause.Returning{}).Where("id = ?", PhotoID).Updates(models.Photo{Title: p.Title, Caption: p.Caption, PhotoURL: p.PhotoURL})
	if dbReturn.Error != nil {
		return nil, dbReturn.Error
	}
	// No RowsAffected (= 0), hacky solution: put if statement here and make errors
	if dbReturn.RowsAffected == 0 {
		return nil, errors.New("data doesn't exist")
	}
	return p, nil
}
