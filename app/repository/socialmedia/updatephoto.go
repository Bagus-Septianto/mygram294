package socialmedia

import (
	"errors"
	"mygram294/app/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r SocialMediaRepository) UpdateSocialMedia(db *gorm.DB, p *models.SocialMedia, SocialMediaID uint) (socialMedia *models.SocialMedia, err error) {
	dbReturn := db.Model(&p).Clauses(clause.Returning{}).Where("id = ?", SocialMediaID).Updates(models.SocialMedia{Name: p.Name, SocialMediaURL: p.SocialMediaURL})
	if dbReturn.Error != nil {
		return nil, dbReturn.Error
	}
	// No RowsAffected (= 0), hacky solution: put if statement here and make errors
	if dbReturn.RowsAffected == 0 {
		return nil, errors.New("data doesn't exist")
	}
	return p, nil
}
