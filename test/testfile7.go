package main

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Vulnerability: Hardcoded credentials - AWS API Key
	awsAPIKey := "AKIAIOSFODNN7EXAMPLE" // This is a hardcoded secret

	// Vulnerability: Hardcoded password
	password := "SuperSecretPassword123" // Hardcoded password

	// Vulnerability: Hardcoded API key for Stripe
	stripeAPIKey := "sk_live_4eC39HqLyjWDarjtT1zdp7dc" // Stripe API key

	// Log the values (insecure to print secrets)
	log.Println("AWS API Key:", awsAPIKey)
	log.Println("Password:", password)
	log.Println("Stripe API Key:", stripeAPIKey)

	// Performing some tasks with the credentials
	processAWSRequest(awsAPIKey)
	processStripePayment(stripeAPIKey)

	// Vulnerability: SQL injection
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userInput := "Robert'); DROP TABLE users; --"
	query := fmt.Sprintf("SELECT * FROM users WHERE name = '%s'", userInput)
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	// Vulnerability: Command injection
	cmd := exec.Command("sh", "-c", "echo "+userInput)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))

	// Vulnerability: Weak password hashing
	hash := md5.Sum([]byte(password))
	fmt.Println("Password hash:", hash)

	// Vulnerability: Insecure HTTP client
	resp, err := http.Get("http://example.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Vulnerability: Missing HTTPS verification
	httpClient := &http.Client{}
	httpClient.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
}

func processAWSRequest(apiKey string) {
	// Simulate AWS request processing
	fmt.Println("Processing AWS request with API key:", apiKey)
}

func processStripePayment(apiKey string) {
	// Simulate Stripe payment processing
	fmt.Println("Processing payment with Stripe API key:", apiKey)
}
