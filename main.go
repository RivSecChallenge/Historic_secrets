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
	if req.Method == http.MethodPost {
		rootPost(w, req)
	} else {
		rootGet(w, req)
	}
}

func rootGet(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("static/index.html")
	check(err)
	err = t.Execute(w, template.HTML(""))
	check(err)
}

func rootPost(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	//Consider handeling missing data
	value := req.Form["secret"][0]
	if value == "Spring2022" {
		fmt.Fprintln(w, "Put secret word here")
	}
}
