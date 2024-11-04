package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}
	t.ExecuteTemplate(w, "index.html", nil)
}

func handleFunc() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8082", nil)
}

func main() {
	fmt.Println("GGG")
}
