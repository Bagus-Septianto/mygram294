package socialmedia

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllSocialMedia godoc
//	@Summary	Read all SocialMedia
//	@Tags		SocialMedia
//	@Produce	json
//	@Success	200	{array}	models.SocialMedia
//	@Failure	400
//	@Failure	401
//	@Failure	404
//	@Router		/socialmedia/GetAll [get]
//	@Security	BearerAuth
func (controller SocialMediaController) GetAll(c *gin.Context) {
	socialMedia, err := controller.SocialMediaService.GetAll()

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
		"socialMedia": socialMedia,
	})
}
