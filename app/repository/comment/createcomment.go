package comment

import (
	"mygram294/app/models"

	"gorm.io/gorm"
)

func (r CommentRepository) CreateComment(db *gorm.DB, p *models.Comment) (comment *models.Comment, err error) {
	err = db.Debug().Create(&p).Error
	return p, err
}
