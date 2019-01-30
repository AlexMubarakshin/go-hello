package routes

import (
	"fmt"
	"html/template"
	"net/http"

	"../models"
	"../utils"
)

var Posts map[string]*models.Post

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	template.ExecuteTemplate(w, "create", nil)
}

func SavePost(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")

	var post *models.Post
	if id != "" {
		post = Posts[id]
		post.Title = title
		post.Content = content
	} else {
		id = utils.GenerateId()
		post := models.NewPost(id, title, content)

		Posts[post.Id] = post
	}

	http.Redirect(w, r, "/", 301)
}

func EditPostHandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	id := r.FormValue("id")
	post, found := Posts[id]
	if !found {
		http.NotFound(w, r)
		return
	}

	template.ExecuteTemplate(w, "create", post)
}
