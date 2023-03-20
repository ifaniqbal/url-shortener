package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-redis/redis"
)

// URL represents a long URL and its associated short URL
type URL struct {
	LongURL  string `json:"long_url"`
	ShortURL string `json:"short_url"`
}

// HashURL returns a base64-encoded SHA-256 hash of the given string
func HashURL(s string) string {
	hash := sha256.Sum256([]byte(s))
	return base64.StdEncoding.EncodeToString(hash[:])
}

// CreateShortURL generates a short URL for the given long URL
func CreateShortURL(longURL string) string {
	return HashURL(longURL)[:8]
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var u URL
			err := json.NewDecoder(r.Body).Decode(&u)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			u.ShortURL = CreateShortURL(u.LongURL)
			err = client.Set(u.ShortURL, u.LongURL, 0).Err()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(u)
		} else if r.Method == http.MethodGet {
			shortURL := r.URL.Path[1:]
			longURL, err := client.Get(shortURL).Result()
			if err == redis.Nil {
				http.NotFound(w, r)
				return
			} else if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, longURL, http.StatusFound)
		} else {
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
