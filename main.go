package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const PORT = ":8080"

func main() {
	http.HandleFunc("/", indexRoute)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))

	println("Server listen on ", PORT)

	if err := http.ListenAndServe(PORT, nil); err != nil {
		panic(err)
	}
}

func indexRoute(writer http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(writer, err.Error())
	}

	template.ExecuteTemplate(writer, "index", nil)
}
