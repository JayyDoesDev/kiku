package sharex

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

func UploadsHandler(w http.ResponseWriter, r *http.Request) {
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
			http.ServeFile(w, r, filepath.Join("storage", file.Name()))
			return
		}
	}
}
