package comment

import (
	"mygram294/app/models"
	"mygram294/pkg/common"
	"mygram294/pkg/helpers"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CreateComment godoc
//	@Summary		Create a Comment
//	@Description	Creating a Comment
//	@Tags			Comment
//	@Accept			json
//	@Produce		json
//	@Param			Comment	body		string			true	"Comment Fields"	SchemaExample({"photo_id":1,"message":"Lorem ipsum dolor sit amet"})
//	@Success		201		{object}	models.Comment	"Comment"
//	@Failure		400
//	@Failure		401
//	@Failure		404
//	@Router			/comment/CreateComment [post]
//	@Security		BearerAuth
func (controller CommentController) CreateComment(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	comment := models.Comment{}

	if contentType == common.AppJSON {
		c.ShouldBindJSON(&comment)
	} else {
		c.ShouldBind(&comment)
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	commentReturn, err := controller.CommentService.CreateComment(&comment, userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"comment": commentReturn,
	})
}
