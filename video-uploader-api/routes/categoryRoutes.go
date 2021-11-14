package routes

import (
	"net/http"

	"github.com/FerdaneOgut/video-uploader-api/db"
	"github.com/FerdaneOgut/video-uploader-api/models"
	errorutils "github.com/FerdaneOgut/video-uploader-api/utils/errorUtils"
	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var categories []models.Category
	categoryRes := db.DB.Find(&categories)
	if categoryRes.Error != nil {
		errorutils.NewInternalServerError(categoryRes.Error.Error())
		return
	}
	c.JSON(http.StatusOK, categories)
}
