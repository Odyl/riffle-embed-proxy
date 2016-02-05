package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hi there", r.URL.Path[1:])
	http.Redirect(w, r, "http://read.rifflebooks.com"+r.URL.Path, 301)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":5000", nil)
}
