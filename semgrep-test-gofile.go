package main

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	_ "github.com/go-sql-driver/mysql"
	"github.com/dgrijalva/jwt-go"
)

func main() {
	// Vulnerability 1: Hardcoded credentials
	username := "admin"
	password := "mysecretpassword"

	// Vulnerability 2: Insecure password hashing
	hash := md5.Sum([]byte(password))
	fmt.Println("Password hash:", hash)

	// Vulnerability 3: SQL injection
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

	// Vulnerability 4: Command injection
	cmd := exec.Command("sh", "-c", "echo "+userInput)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))

	// Vulnerability 5: Insecure JWT token signing
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})
	tokenString, err := token.SignedString([]byte("secretkey"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("JWT token:", tokenString)

	// Vulnerability 6: Insecure HTTP client
	resp, err := http.Get("http://example.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Vulnerability 7: Missing HTTPS verification
	httpClient := &http.Client{}
	httpClient.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	// Vulnerability 8: Using vulnerable dependency
	// github.com/dgrijalva/jwt-go is vulnerable to CVE-2020-26160
}

func processRequest(w http.ResponseWriter, r *http.Request) {
	// Vulnerability 9: Potential cross-site scripting (XSS) vulnerability
	fmt.Fprintf(w, "<script>alert('%s')</script>", r.URL.Query().Get("input"))
}

func insecureFileHandling() {
	// Vulnerability 10: Potential path traversal vulnerability
	filePath := "../sensitive_data.txt"
	_, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
}
