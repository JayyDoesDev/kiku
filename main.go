package main

import (
	"fmt"
	"html/template"
	"kiku/main/lib"
	"kiku/main/middleware"
	"kiku/main/routes/photos"
	"kiku/main/routes/shortener"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var tmpl *template.Template

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load ENV")
	}
	lib.InitMongoDB()
	tmpl = template.Must(template.ParseGlob("public/**/*.html"))
}

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	r.PathPrefix("/files/").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir("./storage"))))
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if err := tmpl.ExecuteTemplate(w, "home.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}).Methods("GET")

	r.Handle("/upload", middleware.Auth(http.HandlerFunc(photos.UploadHandler))).Methods("POST")
	r.Handle("/shorten", middleware.Auth(http.HandlerFunc(shortener.ShortenHandler))).Methods("POST")
	r.HandleFunc("/uploads/{id}", photos.UploadsWithParametersHandler(tmpl)).Methods("GET")
	r.HandleFunc("/uploads", photos.UploadsHandler(tmpl)).Methods("GET")

	fmt.Println("Listening port 3000")
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), r))
}
