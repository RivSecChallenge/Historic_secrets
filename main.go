package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", root)

	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func root(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello world \n")
}
