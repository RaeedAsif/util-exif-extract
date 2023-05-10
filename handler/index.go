package handler

import (
	"html/template"
	"net/http"
)

func ViewIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("html/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
