package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/RaeedAsif/util-exif-extract/route"
	"github.com/RaeedAsif/util-exif-extract/service"
)

var (
	imagesDir = "images"
	csvFile   = "result"
)

func main() {
	log.Println("Starting...")
	fileList := make([]string, 0)
	err := filepath.Walk(imagesDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return err
	})
	if err != nil {
		fmt.Printf("Error directory %s: %s\n", imagesDir, err.Error())
		os.Exit(1)
	}

	fileWriter, err := service.NewWriter(csvFile + ".csv")
	if err != nil {
		log.Panic(err)
	}

	for _, path := range fileList {
		if isImageFile(path) {
			log.Println("Processing: ", path)
			gpsExif, err := service.NewGpsExif(path)
			if err != nil {
				fileWriter.WriteErrorRow(path, err)
				continue
			}

			err = fileWriter.WriteRow(path, gpsExif.GetLatitude(), gpsExif.GetLongitude())
			if err != nil {
				log.Println("Error writing row: ", err)
			}
		}
	}

	fileWriter.Close()
	log.Println("Done!")

	route.InitRoutes(csvFile + ".csv")

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func isImageFile(filename string) bool {
	validExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".bmp":  true,
		".svg":  true,
	}

	ext := strings.ToLower(filepath.Ext(filename))
	return validExts[ext]
}
