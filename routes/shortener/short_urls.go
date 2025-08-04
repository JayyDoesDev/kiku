package shortener

import (
	"encoding/json"
	"kiku/main/db"
	"net/http"
)

func ShortUrlsHandler(w http.ResponseWriter, r *http.Request) {
	short_urls := db.GetShortenedUrls()
	urls := []db.ShortUrl{}

	for _, url := range short_urls {
		urls = append(urls, db.ShortUrl{ID: url.ID, Destination: url.Destination})
	}

	json.NewEncoder(w).Encode(urls)
}
