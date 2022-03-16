package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Get environment variables for secrets from K8s Secrets store
	userValue := os.Getenv("APP_USERNAME")
	passValue := os.Getenv("APP_PASSWORD")

	// Print the secrets
	fmt.Fprintf(w, "K8s Secrets Example\n==========\nUsername: %s\nPassword: %s", userValue, passValue)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
