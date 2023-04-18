package comment

import (
	"mygram294/app/models"
)

func (service CommentService) GetOne(CommentID uint) (comment *models.Comment, err error) {
	comment, err = service.CommentRepository.GetOne(service.DB, CommentID)
	if err != nil {
		return nil, err
	}

	return comment, nil
}
