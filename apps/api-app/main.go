package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cyberark/conjur-api-go/conjurapi"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Get variable path from environment variable
	conjurUserObject := os.Getenv("CONJUR_USER_OBJECT")
	conjurPassObject := os.Getenv("CONJUR_PASS_OBJECT")

	config, err := conjurapi.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %s", err)
	}

	conjurClient, err := conjurapi.NewClientFromEnvironment(config)
	if err != nil {
		log.Fatalf("Error creating client: %s", err)
	}

	userObjectValue, err := conjurClient.RetrieveSecret(conjurUserObject)
	if err != nil {
		log.Fatalf("Error retrieving %s: %s", userObjectValue, err)
	}

	passObjectValue, err := conjurClient.RetrieveSecret(conjurPassObject)
	if err != nil {
		log.Fatalf("Error retrieving %s: %s", passObjectValue, err)
	}

	fmt.Fprintf(w, "API Example\n==========\nUser: %s\nPassword: %s", string(userObjectValue), string(passObjectValue))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
