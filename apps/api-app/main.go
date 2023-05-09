package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/cyberark/conjur-api-go/conjurapi"
	_ "github.com/lib/pq"
	"github.com/olekukonko/tablewriter"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Get variables from environment variable
	// CONJUR_SECRETS=ADDRESS;DATABASE_NAME;USERNAME;PASSWORD
	// Replace above with your own variables
	conjurSecrets := os.Getenv("CONJUR_SECRETS")
	if conjurSecrets == "" {
		log.Fatal("CONJUR_SECRETS environment variable not set")
	}

	// Split conjurSecrets on semi-colon into a slice of strings
	conjurSecretsSlice := strings.Split(conjurSecrets, ";")

	// Load Conjur configuration from environment variables
	config, err := conjurapi.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %s", err)
	}

	// Create a new Conjur client
	conjurClient, err := conjurapi.NewClientFromEnvironment(config)
	if err != nil {
		log.Fatalf("Error creating client: %s", err)
	}

	// Retrieve secret(s) from Conjur
	secretsValueMap, err := conjurClient.RetrieveBatchSecrets(conjurSecretsSlice)
	if err != nil {
		log.Fatalf("Error retrieving secret(s): %s", err)
	}

	fmt.Fprintf(w, "API Example\n==========\nHost: %s\nDatabase Name: %s\nUser: %s\nPassword: %s",
		string(secretsValueMap[conjurSecretsSlice[0]]), // Host
		string(secretsValueMap[conjurSecretsSlice[1]]), // Database Name
		string(secretsValueMap[conjurSecretsSlice[2]]), // Username
		string(secretsValueMap[conjurSecretsSlice[3]])) // Password

	// Connect to the PostgreSQL database
	connStr := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
		string(secretsValueMap[conjurSecretsSlice[0]]), // Host
		string(secretsValueMap[conjurSecretsSlice[1]]), // Database Name
		string(secretsValueMap[conjurSecretsSlice[2]]), // Username
		string(secretsValueMap[conjurSecretsSlice[3]])) // Password
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Fprintf(w, "\nFailed to connect to database: %v", err)
		return
	}
	defer db.Close()

	// Query the 'customers' table
	rows, err := db.Query("SELECT * FROM customers")
	if err != nil {
		fmt.Fprintf(w, "\nFailed to execute query: %v", err)
		return
	}
	defer rows.Close()

	// Get column descriptions
	columns, err := rows.ColumnTypes()
	if err != nil {
		fmt.Fprintf(w, "\nFailed to get column types: %v", err)
		return
	}

	// Initialize tablewriter
	table := tablewriter.NewWriter(w)
	header := make([]string, len(columns))
	for i, column := range columns {
		header[i] = column.Name()
	}
	table.SetHeader(header)
	table.SetRowLine(true)
	table.SetAutoWrapText(false)

	// Iterate through rows and write to table
	for rows.Next() {
		// Prepare a slice of interface{} to hold the scanned values
		values := make([]interface{}, len(columns))
		pointers := make([]interface{}, len(columns))

		for i := range values {
			pointers[i] = &values[i]
		}

		// Scan the row data into the values slice
		err := rows.Scan(pointers...)
		if err != nil {
			fmt.Fprintf(w, "\nFailed to scan row: %v", err)
			return
		}

		// Add the row to the table
		row := make([]string, len(columns))
		for i, val := range values {
			row[i] = fmt.Sprintf("%v", val)
		}
		table.Append(row)
	}

	// Check for any errors encountered during iteration
	err = rows.Err()
	if err != nil {
		fmt.Fprintf(w, "\nFailed to iterate over rows: %v", err)
		return
	}

	// Render the table
	fmt.Fprint(w, "\n\nResults from the database:\n")
	table.Render()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
