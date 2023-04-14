package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
	body := r.Body
	defer body.Close()
	fmt.Println(body)
}

func main() {
	http.HandleFunc("/webhook", handler)
	fmt.Println("server starting...")
	http.ListenAndServe(":8080", nil)
}
