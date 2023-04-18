package comment

import (
	"mygram294/app/models"
	"mygram294/pkg/common"
	"mygram294/pkg/helpers"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// UpdateComment godoc
//	@Summary		Update a Comment
//	@Description	Updating a Comment by CommentID
//	@Tags			Comment
//	@Accept			json
//	@Produce		json
//	@Param			CommentID	path		int				true	"Comment ID"
//	@Param			Comment		body		string			true	"Comment Fields"	SchemaExample({"photo_id":1,"message":"Lorem ipsum dolor sit amet"})
//	@Success		200			{object}	models.Comment	"Comment"
//	@Failure		400
//	@Failure		401
//	@Failure		404
//	@Router			/comment/UpdateComment/{CommentID} [put]
//	@Security		BearerAuth
func (controller CommentController) UpdateComment(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	comment := models.Comment{}

	if contentType == common.AppJSON {
		c.ShouldBindJSON(&comment)
	} else {
		c.ShouldBind(&comment)
	}

	CommentID, _ := strconv.Atoi(c.Param("CommentID"))
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	commentReturn, err := controller.CommentService.UpdateComment(&comment, userID, uint(CommentID))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"comment": commentReturn,
	})
}
