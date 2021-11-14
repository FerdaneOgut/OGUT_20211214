package services

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"os/exec"
	"strings"

	"github.com/FerdaneOgut/video-uploader-api/db"
	"github.com/FerdaneOgut/video-uploader-api/models"
	errorutils "github.com/FerdaneOgut/video-uploader-api/utils/errorUtils"
	"github.com/gin-gonic/gin"
)

var validExt = []string{
	"MP4", "MOV", "mp4", "mov",
}

const MAX_UPLOAD_SIZE = 1024 * 1024 * 200

type VideoDto struct {
	Title        string                `form:"title" binding:"required"`
	Category     *int                  `form:"category", binding:"required"`
	File         *multipart.FileHeader `form:"file" binding:"required"`
	UID          string
	Path         string
	Thumbnail64  []byte
	Thumbnail128 []byte
	Thumbnail256 []byte
}

func (v *VideoDto) Validate() *errorutils.ErrorResponse {
	if v.Title == "" {
		return errorutils.NewBadRequestError("Title is required")
	}
	if v.Category == nil || *v.Category == 0 {
		return errorutils.NewBadRequestError("Category is required")
	}
	if v.File == nil {
		return errorutils.NewBadRequestError("File is required")
	}
	fmt.Println(v.File.Size)
	if v.File.Size > MAX_UPLOAD_SIZE {
		return errorutils.NewBadRequestError("Too large data size.")
	}
	fileExt := strings.Split(v.File.Filename, ".")[1]
	_, exist := findExtention(fileExt)
	if !exist {
		return errorutils.NewBadRequestError("Unsupported data type! Valid types are:  " + strings.Join(validExt, ","))
	}

	return nil
}
func (v *VideoDto) SaveFile(c *gin.Context) *errorutils.ErrorResponse {

	mainP := fmt.Sprintf("./data/videos/%s", v.UID)

	createFileErr := os.MkdirAll(mainP, 0700)
	if createFileErr != nil {
		log.Println("createfile")
		log.Println(createFileErr.Error())
		return errorutils.NewInternalServerError(createFileErr.Error())
	}
	path := fmt.Sprintf("%s/%s", mainP, v.File.Filename)
	v.Path = path
	saveErr := c.SaveUploadedFile(v.File, path)
	if saveErr != nil {

		log.Println("createfile save")
		log.Println(saveErr.Error())
		return errorutils.NewInternalServerError(saveErr.Error())
	}
	return nil
}

func (v *VideoDto) CreateThumbnails() *errorutils.ErrorResponse {

	frame, err64 := getFirstFrame(v.Path, 64)
	if err64 != nil {
		return errorutils.NewInternalServerError(err64.Error())
	}
	frame128, err128 := getFirstFrame(v.Path, 128)
	if err128 != nil {
		return errorutils.NewInternalServerError(err128.Error())
	}
	frame256, err256 := getFirstFrame(v.Path, 256)
	if err256 != nil {
		return errorutils.NewInternalServerError(err256.Error())
	}
	mainPImages := fmt.Sprintf("./data/images/%s", v.UID)

	createFileImgErr := os.MkdirAll(mainPImages, os.ModePerm)
	if createFileImgErr != nil {
		log.Println("thumbnail")
		log.Println(createFileImgErr.Error())
		return errorutils.NewInternalServerError(createFileImgErr.Error())
	}
	image64Path := fmt.Sprintf("%s/%s", mainPImages, "thumnbnail_64.jpeg")
	image128Path := fmt.Sprintf("%s/%s", mainPImages, "thumnbnail_128.jpeg")
	image256Path := fmt.Sprintf("%s/%s", mainPImages, "thumnbnail_256.jpeg")
	write64Err := ioutil.WriteFile(image64Path, *frame, 0777)
	if write64Err != nil {
		return errorutils.NewInternalServerError(write64Err.Error())
	}
	write128Err := ioutil.WriteFile(image128Path, *frame128, 0777)
	if write128Err != nil {
		return errorutils.NewInternalServerError(write128Err.Error())
	}
	write256Err := ioutil.WriteFile(image256Path, *frame256, 0777)
	if write256Err != nil {
		return errorutils.NewInternalServerError(write256Err.Error())
	}
	v.Thumbnail64 = *frame
	v.Thumbnail128 = *frame128
	v.Thumbnail256 = *frame256
	return nil
}

func (v *VideoDto) ToVideo() models.Video {
	vd := models.Video{}
	vd.CategoryID = *v.Category
	vd.Path = v.Path
	vd.Thumbnail128 = v.Thumbnail128
	vd.Thumbnail256 = v.Thumbnail256
	vd.Thumbnail64 = v.Thumbnail64
	vd.Title = v.Title
	vd.UID = v.UID
	return vd
}

func SaveVideo(v *models.Video) *errorutils.ErrorResponse {
	res := db.DB.Create(v)
	if res.Error != nil {
		return errorutils.NewInternalServerError(res.Error.Error())
	}
	return nil
}

func findExtention(val string) (int, bool) {
	for i, item := range validExt {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func getFirstFrame(path string, size int) (*[]byte, error) {
	//gets the frames after 1 sec
	cmd := exec.Command(
		"ffmpeg",
		"-i", path,
		"-ss", "00:00:01.00",
		"-vframes", "1",
		"-s", fmt.Sprintf("%dx%d", size, size),
		"-f", "singlejpeg", "-")
	var buffer bytes.Buffer
	cmd.Stdout = &buffer
	if cmd.Run() != nil {
		fmt.Println(cmd.Run().Error())
		return nil, fmt.Errorf("Cannot run or found ffmpeg file by path: %s", path)
	}
	b := buffer.Bytes()

	return &b, nil
}
