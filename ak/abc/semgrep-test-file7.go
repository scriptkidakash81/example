package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func main() {
	// Critical severity vulnerability: Use of weak MD5 hash
	hash := md5.Sum([]byte("password123"))
	fmt.Printf("MD5 Hash: %x\n", hash)

	// Critical severity vulnerability: Use of insecure random number generator
	buf := make([]byte, 16)
	rand.Read(buf)
	fmt.Printf("Random bytes: %x\n", buf)

	// Critical severity vulnerability: Use of insecure deserialization
	jsonStr := `{"name":"John","admin":true}`
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deserialized data: %+v\n", data)

	// Critical severity vulnerability: Command injection
	cmd := fmt.Sprintf("ping -c 1 %s", "example.com")
	output, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Command output: %s\n", output)

	// Critical severity vulnerability: Insecure use of net/http
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))

	// Critical severity vulnerability: Use of insecure base64 encoding
	encodedStr := base64.StdEncoding.EncodeToString([]byte("secret data"))
	fmt.Printf("Base64 encoded string: %s\n", encodedStr)

	// Critical severity vulnerability: Use of insecure string concatenation
	userInput := "example.com"
	url := "http://" + userInput + "/api/data"
	fmt.Printf("Concatenated URL: %s\n", url)
}
