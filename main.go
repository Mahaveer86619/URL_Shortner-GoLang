package main

import (
	"fmt"
	"net/http"

	"github.com/Mahaveer86619/URL_Shortner-GoLang/urlshortner"
	"github.com/joho/godotenv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from your Golang server!")
}

func main() {
	godotenv.Load()

	fmt.Println("Starting server...")

	http.HandleFunc("/", handler)
	http.HandleFunc("/shorten", urlshortner.Shorten)

	fmt.Println("Server listening on port 2020")
	http.ListenAndServe(":2020", nil)
}
