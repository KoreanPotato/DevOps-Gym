package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("PORT 80")
	http.ListenAndServe(":80", nil)
}
