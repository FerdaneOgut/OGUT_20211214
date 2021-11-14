package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/FerdaneOgut/video-uploader-api/db"
	"github.com/FerdaneOgut/video-uploader-api/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	genericInit()
	r := gin.Default()

	r.Use(cors.Default())
	//	v1 := r.Group("/v1/api")
	r.POST("/video", routes.AddVideo)
	r.GET("/video", routes.GetVideos)
	r.GET("/video/:id", routes.ServeVideo)
	r.GET("/categories", routes.GetCategories)

	r.Run()
}

func genericInit() {
	// make sure we have a working tempdir, because:
	// os.TempDir(): The directory is neither guaranteed to exist nor have accessible permissions.
	tempDir := os.TempDir()
	if err := os.MkdirAll(tempDir, 1777); err != nil {
		log.Fatalf("Failed to create temporary directory %s: %s", tempDir, err)
	}
	tempFile, err := ioutil.TempFile("", "genericInit_")
	if err != nil {
		log.Fatalf("Failed to create tempFile: %s", err)
	}
	_, err = fmt.Fprintf(tempFile, "Hello, World!")
	if err != nil {
		log.Fatalf("Failed to write to tempFile: %s", err)
	}
	if err := tempFile.Close(); err != nil {
		log.Fatalf("Failed to close tempFile: %s", err)
	}
	if err := os.Remove(tempFile.Name()); err != nil {
		log.Fatalf("Failed to delete tempFile: %s", err)
	}
	log.Printf("Using temporary directory %s", tempDir)
}
