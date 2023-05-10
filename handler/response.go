package handler

import (
	"html/template"
	"net/http"

	"github.com/util-exif-extract/service"
)

func ViewResultHandler(w http.ResponseWriter, r *http.Request, filename string) {
	records, err := service.GetRecords(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("html/template.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, records)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
