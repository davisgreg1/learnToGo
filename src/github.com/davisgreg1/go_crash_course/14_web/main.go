package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About ME</h1>")
}
func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>404</h1>")
}

func main() {
	http.HandleFunc("*", notFound)
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server Starting...")
	http.ListenAndServe(":7777", nil)
}
