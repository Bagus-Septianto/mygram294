package comment

import (
	"mygram294/app/models"

	"gorm.io/gorm"
)

func (r CommentRepository) GetAll(db *gorm.DB) (comment []*models.Comment, err error) {
	err = db.Debug().Find(&comment).Error
	return comment, err
}
