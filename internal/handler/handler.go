package handler

import (
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/templates/pages/index.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}
