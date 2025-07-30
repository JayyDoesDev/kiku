package main

import (
	"encoding/json"
	"fmt"
	"kiku/main/middleware"
	"kiku/main/routes/sharex"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load ENV")
	}
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		response := map[string]string{
			"hello": "world",
		}

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, "Error", http.StatusInternalServerError)
			return
		}
	}).Methods("GET")
	r.Handle("/upload", middleware.Auth(http.HandlerFunc(sharex.UploadHandler))).Methods("POST")
	r.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./storage"))))
	r.HandleFunc("/uploads/{id}", sharex.UploadsHandler).Methods("GET")

	fmt.Println("Listening port 3000")
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), r))
}
