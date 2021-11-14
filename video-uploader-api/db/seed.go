package db

import (
	"log"

	"github.com/FerdaneOgut/video-uploader-api/models"
	"gorm.io/gorm"
)

var (
	categories = []models.Category{
		{Name: "Exercise"},
		{Name: "Education"},
		{Name: "Recipe"},
	}
)

func seed(db *gorm.DB) {
	for i := range categories {
		v := categories[i]

		res := db.Where(models.Category{Name: v.Name}).FirstOrCreate(&v)
		if res.Error != nil {
			log.Fatalf("cannot seed categories: %v", res.Error)
		}
	}

}
