package shortener

import (
	"encoding/json"
	"html/template"
	"kiku/main/db"
	"kiku/main/routes/photos"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Page struct {
	ID          string
	Destination string
	Url         string
}

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

func GoToLongLinkHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		idx := params["id"]

		shortned_url := db.FindShortenedUrl(idx)
		if shortned_url != nil {
			data := Page{
				ID:          shortned_url.ID,
				Destination: shortned_url.Destination,
				Url:         os.Getenv("API_URL"),
			}
			if err := tmpl.ExecuteTemplate(w, "redirect.html", data); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}

		http.NotFound(w, r)
	}
}
