package sharex

import (
	"encoding/json"
	"fmt"
	"io"
	"kiku/main/lib"
	"net/http"
	"os"
	"path/filepath"
)

type Response struct {
	Success bool   `json:"success"`
	Url     string `json:"url"`
}

type Request struct {
	File        http.File `json:"file"`
	Name        string    `json:"name"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "File not found in request", http.StatusBadRequest)
		return
	}
	defer file.Close()

	saveDir := "./storage"
	os.MkdirAll(saveDir, 0755)

	str, err := lib.GenerateRandomString(10)
	if err != nil {
		http.Error(w, "Error generating filename", http.StatusInternalServerError)
		return
	}
	ext := filepath.Ext(handler.Filename)
	filename := str + ext

	savePath := filepath.Join(saveDir, filename)

	dst, err := os.Create(savePath)
	if err != nil {
		fmt.Print(err)
		http.Error(w, "Unable to save the file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	io.Copy(dst, file)

	url := fmt.Sprintf(os.Getenv("API_URL")+"/uploads/%s", filename)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&Response{Success: true, Url: url})

}
