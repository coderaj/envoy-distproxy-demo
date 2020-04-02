package main

import (
	"fmt"
	"os"
	"log"

	"net/http"
)

func hostname() (string) {

	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	return name
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, hostname())
}

func main() {
	http.HandleFunc("/hi", handler)

	log.Fatal(http.ListenAndServe(":8082", nil))
}