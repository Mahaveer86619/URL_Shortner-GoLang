package urlshortner

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Mahaveer86619/URL_Shortner-GoLang/config"
)

type ShortenedURL struct {
	ID          string    `json:"id"`
	OriginalURL string    `json:"original_url"`
	ShortURL    string    `json:"short_url"`
	CreatedDate time.Time `json:"created_date"`
}

func gnerateShortURL(originalURL string) (string, error) {
	//* Hasher initialisation
	hasher := md5.New()
	//* Hashing the original URL into a byte slice
	hasher.Write([]byte(originalURL))
	//* Converting the byte slice into a string
	data := hasher.Sum(nil)
	hash := hex.EncodeToString(data)
	//* Returning the first 8 characters of the hashed string
	return hash[:8], nil
}

func createURL(originalURL string) (string, error) {
	ctx := context.Background()

	//* Get the short URL from generator
	shortURL, err := gnerateShortURL(originalURL)
	if err != nil {
		log.Fatal("Error while generating short URL: ", err)
	}

	fmt.Println("Short URL: ", shortURL)

	//* Get a reference to the database
	client, err := config.GetClient()
	if err != nil {
		return "", err
	}

	//* Create a reference to the new short URL document
	ref := client.Collection("shortUrls").NewDoc()

	//* Make a new ShortURL struct
	shortURLData := &ShortenedURL{
		OriginalURL: originalURL,
		ShortURL:    shortURL,
		CreatedDate: time.Now(),
	}

	//* Add the ShortURL struct data to the new document
	_, err = ref.Set(ctx, shortURLData)
	if err != nil {
		return "", err
	}

	return shortURL, nil
}

func Shorten(w http.ResponseWriter, r *http.Request) {
	// Set content type for JSON response
    w.Header().Set("Content-Type", "application/json")

    // Decode request body into a map
    var data map[string]string
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&data)
    if err != nil {
        http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
        return
    }

    // Check for required field "url" in the request body
    originalURL, ok := data["url"]
    if !ok {
        http.Error(w, "Missing required field 'url' in request body", http.StatusBadRequest)
        return
    }

    // Shorten the URL
    shortURL, err := createURL(originalURL)
    if err != nil {
        http.Error(w, "Error shortening URL: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Construct the response object
    response := map[string]string{"short_url": shortURL}

    // Encode the response object as JSON
    jsonData, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Write the JSON response
    w.WriteHeader(http.StatusOK)
    w.Write(jsonData)
}

// func getURL(shortURL string) (URL, error) {

// }
