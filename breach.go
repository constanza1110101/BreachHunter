package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

const (
	hibpAPIURL = "https://api.pwnedpasswords.com/range/"
	userAgent  = "BreachHunter-OSINT-Tool"
)

type EmailResult struct {
	Email     string `json:"email"`
	Breached  bool   `json:"breached"`
	BreachNum int    `json:"breach_count"`
	Breaches  []string `json:"breaches,omitempty"`
}

type PasswordResult struct {
	Password string `json:"password"`
	Breached bool   `json:"breached"`
	Count    int    `json:"count"`
}

func main() {
	// Define command line flags
	emailFile := flag.String("email-file", "", "File containing emails to check")
	email := flag.String("email", "", "Single email to check")
	passwordFile := flag.String("password-file", "", "File containing passwords to check")
	password := flag.String("password", "", "Single password to check")
	outputFile := flag.String("output", "", "Output file for results (JSON format)")
	flag.Parse()

	// Print banner
	color.Cyan("╔═════════════════════════════════════╗")
	color.Cyan("║ BreachHunter - Data Breach Scanner  ║")
	color.Cyan("╚═════════════════════════════════════╝")

	if *email == "" && *emailFile == "" && *password == "" && *passwordFile == "" {
		color.Red("Error: You must provide at least one email or password to check")
		flag.Usage()
		os.Exit(1)
	}

	var emailResults []EmailResult
	var passwordResults []PasswordResult

	// Process single email
	if *email != "" {
		color.Yellow("[*] Checking email: %s", *email)
		result := checkEmail(*email)
		emailResults = append(emailResults, result)
		printEmailResult(result)
	}

	// Process email file
	if *emailFile != "" {
		color.Yellow("[*] Checking emails from file: %s", *emailFile)
		results, err := processEmailFile(*emailFile)
		if err != nil {
			color.Red("[-] Error processing email file: %s", err)
		} else {
			emailResults = append(emailResults, results...)
			for _, result := range results {
				printEmailResult(result)
			}
		}
	}

	// Process single password
	if *password != "" {
		color.Yellow("[*] Checking password security")
		result := checkPassword(*password)
		passwordResults = append(passwordResults, result)
		printPasswordResult(result)
	}

	// Process password file
	if *passwordFile != "" {
		color.Yellow("[*] Checking passwords from file: %s", *passwordFile)
		results, err := processPasswordFile(*passwordFile)
		if err != nil {
			color.Red("[-] Error processing password file: %s", err)
		} else {
			passwordResults = append(passwordResults, results...)
			for _, result := range results {
				printPasswordResult(result)
			}
		}
	}

	// Save results if output file specified
	if *outputFile != "" && (len(emailResults) > 0 || len(passwordResults) > 0) {
		saveResults(*outputFile, emailResults, passwordResults)
	}
}

func checkEmail(email string) EmailResult {
	result := EmailResult{
		Email:    email,
		Breached: false,
	}

	// This is a mock implementation
	// In a real tool, you would use an actual API to check for breaches
	
	// Simulate API call with delay
	time.Sleep(500 * time.Millisecond)
	
	// For demo purposes, we'll consider emails with "test" to be breached
	if strings.Contains(email, "test") {
		result.Breached = true
		result.BreachNum = 3
		result.Breaches = []string{"Adobe", "LinkedIn", "Dropbox"}
	}
	
	return result
}

func processEmailFile(filename string) ([]EmailResult, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var results []EmailResult
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		email := strings.TrimSpace(scanner.Text())
		if email != "" {
			color.Yellow("[*] Checking email: %s", email)
			result := checkEmail(email)
			results = append(results, result)
			// Add a small delay to avoid rate limiting
			time.Sleep(100 * time.Millisecond)
		}
	}

	if err := scanner.Err(); err != nil {
		return results, err
	}
	
	return results, nil
}

func checkPassword(password string) PasswordResult {
	result := PasswordResult{
		Password: maskPassword(password),
		Breached: false,
		Count:    0,
	}

	// Hash the password with SHA-1
	h := sha1.New()
	h.Write([]byte(password))
	hash := strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	prefix := hash[0:5]
	suffix := hash[5:]

	// Query the API with the prefix
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", hibpAPIURL+prefix, nil)
	if err != nil {
		return result
	}
	
	req.Header.Add("User-Agent", userAgent)
	resp, err := client.Do(req)
	if err != nil {
		return result
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return result
	}

	// Read and process the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result
	}

	// Check if the hash suffix is in the response
	lines := strings.Split(string(body), "\r\n")
	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) == 2 && parts[0] == suffix {
			count := 0
			fmt.Sscanf(parts[1], "%d", &count)
			result.Breached = true
			result.Count = count
			break
		}
	}

	return result
}

func processPasswordFile(filename string) ([]PasswordResult, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var results []PasswordResult
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		password := strings.TrimSpace(scanner.Text())
		if password != "" {
			color.Yellow("[*] Checking password")
			result := checkPassword(password)
			results = append(results, result)
			// Add a small delay to avoid rate limiting
			time.Sleep(1500 * time.Millisecond)
		}
	}

	if err := scanner.Err(); err != nil {
		return results, err
	}
	
	return results, nil
}

func printEmailResult(result EmailResult) {
	if result.Breached {
		color.Red("[!] %s was found in %d data breaches:", result.Email, result.BreachNum)
		for _, breach := range result.Breaches {
			color.Red("    - %s", breach)
		}
	} else {
		color.Green("[+] %s was not found in any known data breaches", result.Email)
	}
}

func printPasswordResult(result PasswordResult) {
	if result.Breached {
		color.Red("[!] Password was found in %d data breaches", result.Count)
	} else {
		color.Green("[+] Password was not found in any known data breaches")
	}
}

func maskPassword(password string) string {
	if len(password) <= 2 {
		return "***"
	}
	return password[:1] + strings.Repeat("*", len(password)-2) + password[len(password)-1:]
}

func saveResults(filename string, emailResults []EmailResult, passwordResults []PasswordResult) {
	results := map[string]interface{}{
		"scan_date": time.Now().Format(time.RFC3339),
	}
	
	if len(emailResults) > 0 {
		results["email_results"] = emailResults
	}
	
	if len(passwordResults) > 0 {
		results["password_results"] = passwordResults
	}
	
	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		color.Red("[-] Error serializing results: %s", err)
		return
	}
	
	err = ioutil.WriteFile(filename, jsonData, 0644)
	if err != nil {
		color.Red("[-] Error saving results to %s: %s", filename, err)
		return
	}
	
	color.Green("[+] Results saved to %s", filename)
}
