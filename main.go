package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func absolutePath(path string, handler http.HandlerFunc ) (string, http.HandlerFunc) {
	return path, func (w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != path {
			http.NotFound(w, r)
			return 
		}
		handler(w, r);
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home page testing"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if(err != nil || id < 1) {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if(r.Method != http.MethodPost) {
		w.Header().Set("Allow", "POST")
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet"))
}


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(absolutePath("/", home))
	mux.HandleFunc("/snippets/view", snippetView)
	mux.HandleFunc("/snippets/create", snippetCreate)
	mux.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Static"))
	})

	log.Println("Starting a webserver on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}