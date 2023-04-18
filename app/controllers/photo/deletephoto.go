package photo

import (
	"mygram294/app/models"
	"mygram294/pkg/common"
	"mygram294/pkg/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeletePhoto godoc
//	@Summary		Delete a Photo
//	@Description	Deleting a Photo by PhotoID
//	@Tags			Photo
//	@Produce		json
//	@Param			PhotoID	path	int	true	"Photo ID"
//	@Success		204
//	@Failure		400
//	@Failure		401
//	@Failure		404
//	@Router			/photo/DeletePhoto/{PhotoID} [delete]
//	@Security		BearerAuth
func (controller PhotoController) DeletePhoto(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	photo := models.Photo{}

	if contentType == common.AppJSON {
		c.ShouldBindJSON(&photo)
	} else {
		c.ShouldBind(&photo)
	}

	PhotoID, _ := strconv.Atoi(c.Param("PhotoID"))
	err := controller.PhotoService.DeletePhoto(uint(PhotoID))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
