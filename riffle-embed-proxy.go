package main

import (
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hi there", r.URL.Path[1:])
	http.Redirect(w, r, "https://read.rifflebooks.com"+r.URL.Path, 301)
}

func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
