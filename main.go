package main

import (
	"fmt"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	country := os.Getenv("COUNTRY")
	secret := os.Getenv("SECRET")

	fmt.Fprintf(w, "hello %s, this is your secret: %s\n", country, secret)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
