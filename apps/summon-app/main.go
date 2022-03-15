package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	usernameValue := os.Getenv("APP_USERNAME")
	passwordValue := os.Getenv("APP_PASSWORD")
	fmt.Fprintf(w, "Summon Example\n==========\nUser: %s\nPassword: %s", usernameValue, passwordValue)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
