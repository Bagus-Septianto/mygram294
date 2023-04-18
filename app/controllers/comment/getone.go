package comment

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ReadComment godoc
//	@Summary		Read a Comment by ID
//	@Description	Reading a Comment based on CommentID
//	@Tags			Comment
//	@Accept			json
//	@Produce		json
//	@Param			CommentID	path		int	true	"Comment ID"
//	@Success		200			{object}	models.Comment
//	@Failure		400
//	@Failure		401
//	@Failure		404
//	@Router			/comment/GetOne/{CommentID} [get]
//	@Security		BearerAuth
func (controller CommentController) GetOne(c *gin.Context) {
	CommentID, _ := strconv.Atoi(c.Param("CommentID"))
	comment, err := controller.CommentService.GetOne(uint(CommentID))
	
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"comment": comment,
	})
}
