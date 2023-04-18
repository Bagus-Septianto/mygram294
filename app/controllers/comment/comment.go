package comment

import (
	"mygram294/app/services"

	"gorm.io/gorm"
)

type CommentController struct {
	CommentService services.ICommentService
	DB           *gorm.DB
}

func NewCommentController(commentService services.ICommentService, db *gorm.DB) *CommentController {
	return &CommentController{
		CommentService: commentService,
		DB:           db,
	}
}
