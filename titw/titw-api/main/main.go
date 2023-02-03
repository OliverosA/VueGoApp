package main

import (
	"fmt"
	"net/http"

	"github.com/OliverosA/titw-api/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func main() {

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	/* r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		mime.AddExtensionType(".js", "application/javascript")
		mime.AddExtensionType(".html", "text/html")
		mime.AddExtensionType(".css", "text/css")
		fileServer := http.FileServer(http.Dir("../../titw-fe/emails-app/dist/"))

		http.StripPrefix("/", fileServer)

		fileServer.ServeHTTP(w, r)
	}) */

	/* mime.AddExtensionType(".js", "application/javascript")
	mime.AddExtensionType(".html", "text/html")
	mime.AddExtensionType(".css", "text/css")

	fileserver := http.FileServer(http.Dir("../../titw-fe/emails-app/dist/"))

	//http.Handle("/", http.StripPrefix("/", fileserver))

	r.Handle("/static/", http.StripPrefix("/static", fileserver)) */

	r.Route("/emails", func(r chi.Router) {

		r.Get("/", handlers.GetEmails)
		r.Post("/term", handlers.SearchTerm)

	})

	fmt.Println("Serving on port: 3000")
	http.ListenAndServe(":3000", r)

}
