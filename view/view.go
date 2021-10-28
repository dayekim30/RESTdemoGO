package view

import (
	"html/template"
	"log"
	"net/http"
)

func ViewParse(w http.ResponseWriter, view string) {
	t, err := template.ParseFiles(view)
	if err != nil {
		log.Print("template parsing error", err)
	}
	err = t.Execute(w, nil)
}
