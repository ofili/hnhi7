package api

import (
	"github.com/pkg/errors"
	"html/template"
	"log"
	"net/http"
)

var (
	templates = template.Must(template.ParseGlob("views/*html"))
)

// Pages ...
func Pages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	err := templates.ExecuteTemplate(w, "success.html", nil); if err != nil{
		http.Error(w, errors.New("Something went wrong. If this continues contact an admin").Error(), http.StatusInternalServerError)
		log.Println("error loading template",err)
	}
}