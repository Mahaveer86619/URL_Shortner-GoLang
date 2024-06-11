package main

import (
	"fmt"
	"net/http"

	"github.com/Mahaveer86619/URL_Shortner-GoLang/urlshortner"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	Hello from the URL Shortner GoLang server!

	Shorten a URL:
	POST http://localhost:2020/shorten
	with a JSON body containing the URL 'url' to shorten

	Redirect to a shortened URL:
	GET http://localhost:2020/redirect?shortURL=your-short-url

	Happy shortening!😊
	`)
}

func main() {
	godotenv.Load()

	fmt.Println("Starting server...")

	r := chi.NewRouter()

	registerMethods(r)

	fmt.Println("Server listening on port 2020")
	http.ListenAndServe(":2020", r)
}

func registerMethods(r *chi.Mux) {
	r.Get("/", handler)
	r.Post("/shorten", urlshortner.ShortenHandler)
	r.Get("/redirect", urlshortner.RedirectHandler)
}
