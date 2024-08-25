package main

import (
	"fmt"
	"net/http"
)

func home_page(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello epta")
}

func about(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Omagad")
}

func main(){
	http.HandleFunc("/", home_page)
	http.HandleFunc("/about/", about)
	http.ListenAndServe(":8080", nil)
}