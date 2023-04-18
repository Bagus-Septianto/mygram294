package comment

import (
	"mygram294/app/models"

	"gorm.io/gorm"
)

func (r CommentRepository) GetOne(db *gorm.DB, CommentID uint) (comment *models.Comment, err error) {
	err = db.Debug().First(&comment, "id = ?", CommentID).Error

	return comment, err
}
