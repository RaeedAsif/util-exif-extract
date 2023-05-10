package route

import (
	"net/http"

	"github.com/RaeedAsif/util-exif-extract/handler"
)

func InitRoutes(filename string) {
	http.HandleFunc("/", handler.ViewIndex)

	http.HandleFunc("/result", func(w http.ResponseWriter, r *http.Request) {
		handler.ViewResultHandler(w, r, filename)
	})

	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		handler.ViewResultHandler(w, r, "error_"+filename)
	})
}
