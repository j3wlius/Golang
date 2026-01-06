package main

import (
	"log"
	"net/http"
)

// home handler
func home(w http.ResponseWriter, r *http.Request ){
	w.Write([]byte("Hello, from snippetbox"))
}

// specific snippet
func specific_snippet(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Displaying a specific snippet..."))
}

func create_snippet(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Display for creating a snippet...."))
}

func applicationRoutes() http.Handler{

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/specific_snippet", specific_snippet)
	mux.HandleFunc("/create_snippet", create_snippet)
	return mux
}

func main(){
	log.Print("Starting server on 8080")

	err := http.ListenAndServe(":8080", applicationRoutes())
	log.Fatal(err)
}