package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	chi "github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	// Serve dynamic content
	r.Get("/books", func(w http.ResponseWriter, r *http.Request) {
		tmplPath := filepath.Join("dist", "index.html")
		tmpl, err := template.ParseFiles(tmplPath)
		if err != nil {
			http.Error(w, "Failed to load template", http.StatusInternalServerError)
			return
		}

		data := map[string]interface{}{
			"Title": "Books Page",
			"Books": []string{"Book 1", "Book 2", "Book 3"},
		}

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
		}
	})

	fs := http.FileServer(http.Dir("dist"))
	r.Handle("/*", fs)

	fmt.Println("Starting the server at :3000")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}
