package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting container")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	log.Fatal(http.ListenAndServe((":8081"), nil))
}
