package db

import (
	"fmt"
	"log"
	"strconv"

	"github.com/FerdaneOgut/video-uploader-api/config"
	"github.com/FerdaneOgut/video-uploader-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect connect to db
func Connect() {
	var err error
	p := config.GetValue("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("error parsing env variable! check your env file")
	}

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.GetValue("DB_HOST"), port, config.GetValue("DB_USER"), config.GetValue("DB_PASSWORD"), config.GetValue("DB_NAME"))
	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	migrate()
}

func migrate() {
	DB.AutoMigrate(&models.Category{}, &models.Video{})
	fmt.Println("Migrated")
	seed(DB)
	fmt.Println("Data seeded")
}
