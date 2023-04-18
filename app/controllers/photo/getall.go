package photo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllPhoto godoc
//	@Summary	Read all Photos
//	@Tags		Photo
//	@Produce	json
//	@Success	200	{array}	models.Photo
//	@Failure	400
//	@Failure	401
//	@Failure	404
//	@Router		/photo/GetAll [get]
//	@Security	BearerAuth
func (controller PhotoController) GetAll(c *gin.Context) {
	photo, err := controller.PhotoService.GetAll()

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
