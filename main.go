package main

import (
	"net/http"

	"./models"
	"./routes"
)

const PORT = ":8080"

func main() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))

	routes.Posts = make(map[string]*models.Post, 0)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		routes.HomeHandler(w, r, routes.Posts)
	})

	http.HandleFunc("/create", routes.CreatePostHandler)
	http.HandleFunc("/save", routes.SavePost)

	println("Server listen on ", PORT)

	if err := http.ListenAndServe(PORT, nil); err != nil {
		panic(err)
	}
}
