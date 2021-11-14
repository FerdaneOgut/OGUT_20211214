package routes

import (
	"log"
	"net/http"

	"github.com/FerdaneOgut/video-uploader-api/db"
	"github.com/FerdaneOgut/video-uploader-api/models"
	"github.com/FerdaneOgut/video-uploader-api/services"
	errorutils "github.com/FerdaneOgut/video-uploader-api/utils/errorUtils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AddVideo(c *gin.Context) {
	var input services.VideoDto
	if err := c.ShouldBind(&input); err != nil {
		log.Println(err)
		r := errorutils.NewBadRequestError(err.Error())
		c.JSON(r.Code, r)
		return
	}

	validErr := input.Validate()
	if validErr != nil {
		log.Println(validErr)
		c.JSON(validErr.Code, validErr)
		return
	}
	log.Println(input)
	input.UID = uuid.New().String()
	saveErr := input.SaveFile(c)
	if saveErr != nil {
		log.Println(saveErr)
		c.JSON(saveErr.Code, saveErr)
		return
	}
	saveThumbnailsErr := input.CreateThumbnails()
	if saveThumbnailsErr != nil {
		c.JSON(saveThumbnailsErr.Code, saveThumbnailsErr)
		return
	}
	vd := input.ToVideo()
	addErr := services.SaveVideo(&vd)
	if addErr != nil {
		c.JSON(addErr.Code, addErr)
		return
	}
	c.JSON(http.StatusCreated, vd)
}
func GetVideos(c *gin.Context) {
	var videos []models.Video
	res := db.DB.Preload("Category").Order("ID desc").Find(&videos)
	if res.Error != nil {
		e := errorutils.NewInternalServerError(res.Error.Error())
		c.JSON(e.Code, e)
		return
	}
	c.JSON(http.StatusOK, videos)
}

func ServeVideo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		e := errorutils.NewBadRequestError("Invalid id")
		c.JSON(e.Code, e)
		return
	}
	var video models.Video
	vRes := db.DB.Where("ID=?", id).Find(&video)
	if vRes.Error != nil {

		e := errorutils.NewNotFoundError("Invalid video")
		c.JSON(e.Code, e)
		return
	}
	http.ServeFile(c.Writer, c.Request, video.Path)
}
