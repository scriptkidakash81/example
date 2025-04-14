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
#12
