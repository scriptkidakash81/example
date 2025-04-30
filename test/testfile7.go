package main

import (
	"fmt"
	"log"
)

func main() {
	// Vulnerability: Hardcoded credentials - AWS API Key
	awsAPIKey : "AKIAIOSFODNN7EXAMPLE" // This is a hardcoded secret

	// Vulnerability: Hardcoded password
	password : "SuperSecretPassword123" // Hardcoded password

	// Vulnerability: Hardcoded API key for Stripe
	stripeAPIKey : "sk_live_4eC39HqLyjWDarjtT1zdp7dc" // Stripe API key

	// Log the values (insecure to print secrets)
	log.Println("AWS API Key:", awsAPIKey)
	log.Println("Password:", password)
	log.Println("Stripe API Key:", stripeAPIKey)

	// Performing some tasks with the credentials
	processAWSRequest(awsAPIKey)
	processStripePayment(stripeAPIKey)
}

func processAWSRequest(apiKey string) {
	// Simulate AWS request processing
	fmt.Println("Processing AWS request with API key:", apiKey)
}

func processStripePayment(apiKey string) {
	// Simulate Stripe payment processing
	fmt.Println("Processing payment with Stripe API key:", apiKey)
}

#test
#test3
##
##
package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // Hardcoded credentials (Vulnerability 1)
    dbUser := "admin"
    dbPassword := "mysecretpassword"
    dbHost := "localhost"
    dbName := "mydb"

    // Insecure database connection string (Vulnerability 2)
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbHost, dbName)

    // Potential SQL injection vulnerability (Vulnerability 3)
    userInput := "Robert'); DROP TABLE users; --"
    query := fmt.Sprintf("SELECT * FROM users WHERE name = '%s'", userInput)

    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Insecure logging of sensitive data (Vulnerability 4)
    log.Printf("Connecting to database with credentials: %s:%s", dbUser, dbPassword)

    // Potential credentials exposure through environment variables (Vulnerability 5)
    fmt.Println("DB_PASSWORD:", os.Getenv("DB_PASSWORD"))

    // Execute the query
    rows, err := db.Query(query)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // Print the results
    for rows.Next() {
        var name string
        err := rows.Scan(&name)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(name)
    }
}
