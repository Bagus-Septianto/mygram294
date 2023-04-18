package comment

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllComment godoc
//	@Summary	Read all Comments
//	@Tags		Comment
//	@Produce	json
//	@Success	200	{array}	models.Comment
//	@Failure	400
//	@Failure	401
//	@Failure	404
//	@Router		/comment/GetAll [get]
//	@Security	BearerAuth
func (controller CommentController) GetAll(c *gin.Context) {
	comment, err := controller.CommentService.GetAll()

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
