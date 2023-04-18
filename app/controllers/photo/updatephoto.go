package photo

import (
	"mygram294/app/models"
	"mygram294/pkg/common"
	"mygram294/pkg/helpers"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// UpdatePhoto godoc
//	@Summary		Update a Photo
//	@Description	Updating a Photo by PhotoID
//	@Tags			Photo
//	@Accept			json
//	@Produce		json
//	@Param			PhotoID	path		int				true	"Photo ID"
//	@Param			Photo	body		string			true	"Photo Fields"	SchemaExample({"title":"Lorem Ipsum","caption":"Lorem ipsum dolor sit amet","photo_url":"https://css.category.id/cat.jpg"})
//	@Success		200		{object}	models.Photo	"Photo"
//	@Failure		400
//	@Failure		401
//	@Failure		404
//	@Router			/photo/UpdatePhoto/{PhotoID} [put]
//	@Security		BearerAuth
func (controller PhotoController) UpdatePhoto(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	photo := models.Photo{}

	if contentType == common.AppJSON {
		c.ShouldBindJSON(&photo)
	} else {
		c.ShouldBind(&photo)
	}

	PhotoID, _ := strconv.Atoi(c.Param("PhotoID"))
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	photoReturn, err := controller.PhotoService.UpdatePhoto(&photo, userID, uint(PhotoID))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"photo": photoReturn,
	})
}
