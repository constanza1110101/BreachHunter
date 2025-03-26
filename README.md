# BreachHunter
## Data Breach Scanner

BreachHunter is a Go-based cybersecurity tool that checks if emails and passwords have been compromised in known data breaches, helping organizations and individuals assess their exposure to security risks.

## Features
- **Email breach checking**: Check if your email has been part of a known data breach.
- **Password breach checking**: Check if your passwords have been exposed using k-anonymity for added security.
- **Batch processing**: Check multiple emails or passwords via input files.
- **JSON output**: Supports JSON output for integration with other tools.
- **Privacy-focused design**: Ensures your credentials are never sent in plaintext.

## Requirements
- **Go 1.16+**
- **Go packages**: `github.com/fatih/color`

## Installation

1. **Clone the repository:**
    ```bash
    git clone https://github.com/costanza1110101/BreachHunter.git
    cd BreachHunter
    ```

2. **Build the tool:**
    ```bash
    go build -o breachhunter
    ```

    Alternatively, **install directly**:
    ```bash
    go install github.com/costanza1110101/BreachHunter@latest
    ```

## Usage

Run the tool with the desired options:

```bash
./breachhunter [options]
Examples:
Check a single email:

bash
Copiar código
./breachhunter --email user@example.com
Check emails from a file:

bash
Copiar código
./breachhunter --email-file emails.txt
Check a single password:

bash
Copiar código
./breachhunter --password "mysecretpassword"
Check passwords from a file and output the results to JSON:

bash
Copiar código
./breachhunter --password-file passwords.txt --output results.json
Security Note
Passwords are never sent in plaintext to any API. Only the first 5 characters of the SHA-1 hash are transmitted, ensuring your passwords remain secure during the checking process.

License
This tool is provided for legitimate security assessment purposes only.
