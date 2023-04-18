package photo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ReadPhoto godoc
//	@Summary		Read a Photo by ID
//	@Description	Reading a Photo based on PhotoID
//	@Tags			Photo
//	@Accept			json
//	@Produce		json
//	@Param			PhotoID	path		int	true	"Photo ID"
//	@Success		200		{object}	models.Photo
//	@Failure		400
//	@Failure		401
//	@Failure		404
//	@Router			/photo/GetOne/{PhotoID} [get]
//	@Security		BearerAuth
func (controller PhotoController) GetOne(c *gin.Context) {
	PhotoID, _ := strconv.Atoi(c.Param("PhotoID"))
	photo, err := controller.PhotoService.GetOne(uint(PhotoID))
	
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
		"photo": photo,
	})
}
