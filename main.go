package main

import (
	"html/template"
	"net/http"
	"log"
	"path/filepath"
	"sync"
	"github.com/gorilla/mux"
	"encoding/json"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r  *http.Request) {
	t.once.Do(func() {
		t.templ =  template.Must(template.ParseFiles(filepath.Join("templates",
			t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {

	NewClient()
	r := mux.NewRouter()
	r.Handle("/", &templateHandler{filename: "index.html"})
    r.HandleFunc("/containers", c)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":5005", r))
}

func c(w http.ResponseWriter, req *http.Request) {
	resp := controllers()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(resp)
}