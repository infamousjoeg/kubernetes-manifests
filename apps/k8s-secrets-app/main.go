package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/olekukonko/tablewriter"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Get environment variables for secrets from K8s Secrets store
	hostValue := os.Getenv("DB_HOSTNAME")
	nameValue := os.Getenv("DB_NAME")
	userValue := os.Getenv("DB_USERNAME")
	passValue := os.Getenv("DB_PASSWORD")

	// Print the secrets
	fmt.Fprintf(w, "K8s Secrets Example\n==========\nHostname: %s\nDatabase Name: %s\nUsername: %s\nPassword: %s", hostValue, nameValue, userValue, passValue)

	// Connect to the PostgreSQL database
	connStr := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable", hostValue, nameValue, userValue, passValue)
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
