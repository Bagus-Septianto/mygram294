package comment

import (
	"mygram294/app/models"
	"mygram294/pkg/common"
	"mygram294/pkg/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeleteComment godoc
//	@Summary		Delete a Comment
//	@Description	Deleting a Comment by CommentID
//	@Tags			Comment
//	@Produce		json
//	@Param			CommentID	path	int	true	"Comment ID"
//	@Success		204
//	@Failure		400
//	@Failure		401
//	@Failure		404
//	@Router			/comment/DeleteComment/{CommentID} [delete]
//	@Security		BearerAuth
func (controller CommentController) DeleteComment(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	comment := models.Comment{}

	if contentType == common.AppJSON {
		c.ShouldBindJSON(&comment)
	} else {
		c.ShouldBind(&comment)
	}

	CommentID, _ := strconv.Atoi(c.Param("CommentID"))
	err := controller.CommentService.DeleteComment(uint(CommentID))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
