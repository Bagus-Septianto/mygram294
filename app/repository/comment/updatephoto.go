package comment

import (
	"errors"
	"mygram294/app/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r CommentRepository) UpdateComment(db *gorm.DB, p *models.Comment, CommentID uint) (comment *models.Comment, err error) {
	dbReturn := db.Model(&p).Clauses(clause.Returning{}).Where("id = ?", CommentID).Updates(models.Comment{PhotoID: p.PhotoID, Message: p.Message})
	if dbReturn.Error != nil {
		return nil, dbReturn.Error
	}
	// No RowsAffected (= 0), hacky solution: put if statement here and make errors
	if dbReturn.RowsAffected == 0 {
		return nil, errors.New("data doesn't exist")
	}
	return p, nil
}
