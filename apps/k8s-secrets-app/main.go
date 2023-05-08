package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Get environment variables for secrets from K8s Secrets store
	hostValue := os.Getenv("DB_HOSTNAME")
	nameValue := os.Getenv("DB_NAME")
	userValue := os.Getenv("DB_USERNAME")
	passValue := os.Getenv("DB_PASSWORD")

	// Print the secrets
	fmt.Fprintf(w, "K8s Secrets Example\n==========\nHostname: %s\nDatabase Name: %s\nUsername: %s\nPassword: %s", hostValue, nameValue, userValue, passValue)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
