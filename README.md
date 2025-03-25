BreachHunter
Data Breach Scanner
BreachHunter is a Go-based cybersecurity tool that checks if emails and passwords have been compromised in known data breaches, helping organizations and individuals assess their exposure to security risks.

Features
Email breach checking
Password breach checking (using k-anonymity for security)
Batch processing via input files
JSON output for integration with other tools
Privacy-focused design
Requirements
Go 1.16+
Go packages: github.com/fatih/color
Installation
bash

Hide
# Clone the repository
git clone https://github.com/costanza1110101/BreachHunter.git
cd BreachHunter

# Build the tool
go build -o breachhunter

# Or install directly
go install github.com/costanza1110101/BreachHunter@latest
Usage
bash

Hide
./breachhunter [options]

# Examples:
./breachhunter --email user@example.com
./breachhunter --email-file emails.txt
./breachhunter --password "mysecretpassword"
./breachhunter --password-file passwords.txt --output results.json
Security Note
Passwords are never sent in plaintext to any API. Only the first 5 characters of the SHA-1 hash are transmitted, ensuring your passwords remain secure during the checking process.

License
This tool is provided for legitimate security assessment purposes only.
