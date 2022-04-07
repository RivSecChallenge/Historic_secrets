package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", root)

	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func root(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("static/index.html")
	check(err)
	err = t.Execute(w, template.HTML(""))
	check(err)
}
