package routes

import (
	"fmt"
	"html/template"
	"net/http"

	"../models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request, posts map[string]*models.Post) {
	template, err := template.ParseFiles("templates/home.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	template.ExecuteTemplate(w, "home", posts)
}
