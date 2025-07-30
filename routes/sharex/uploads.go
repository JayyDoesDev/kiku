// sharex/uploads.go
package sharex

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

type Page struct {
	Id    string
	Image string
	Url   string
}

func UploadsHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		idx := params["id"]
		files, err := os.ReadDir("./storage")
		if err != nil {
			http.Error(w, "Unable to fetch storage", http.StatusInternalServerError)
			return
		}
		if len(files) == 0 {
			http.Error(w, "Storage empty", http.StatusInternalServerError)
			return
		}
		for _, file := range files {
			if strings.TrimSuffix(file.Name(), filepath.Ext(file.Name())) == idx {
				data := Page{
					Id:    file.Name(),
					Image: "/files/" + file.Name(),
					Url:   os.Getenv("API_URL"),
				}
				if err := tmpl.ExecuteTemplate(w, "uploads.html", data); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				return
			}
		}
		http.NotFound(w, r)
	}
}
