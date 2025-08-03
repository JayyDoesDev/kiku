package shortener

import (
	"encoding/json"
	"kiku/main/db"
	"kiku/main/routes/photos"
	"net/http"
	"os"
)

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Url string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil || data.Url == "" {
		http.Error(w, "No url property provided", http.StatusBadRequest)
		return
	}

	shortened := db.ShortenUrl(data.Url)
	if shortened.ID == "" {
		http.Error(w, "Failed to shorten URL", http.StatusInternalServerError)
		return
	}

	apiURL := os.Getenv("API_URL")
	shortened_url := apiURL + "/" + shortened.ID

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&photos.Response{Url: shortened_url})
}
