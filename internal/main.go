package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)


func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello .., %q", html.EscapeString(r.URL.Path))
	})

	log.Println("Listening on localhost:222")
	log.Fatal(http.ListenAndServe(":2222", nil))
}