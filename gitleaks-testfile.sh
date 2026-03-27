#!/bin/bash

# Slack webhook token (Vulnerability 1)
SLACK_WEBHOOK="https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX"

# GitHub personal access token (Vulnerability 2)
GITHUB_TOKEN="ghp_XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"

# Google OAuth client secret (Vulnerability 3)
GOOGLE_CLIENT_SECRET="XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"

# Hardcoded IP address of a sensitive server (Vulnerability 4)
SENSITIVE_SERVER="192.168.1.100"

# Using curl with insecure SSL verification (Vulnerability 5)
curl -k -X GET "https://example.com/api/data"

# Test script with potential secrets

# Hardcoded AWS credentials
AWS_ACCESS_KEY_ID="AKIAIOSFODNN7EXAMPLE"
AWS_SECRET_ACCESS_KEY="wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"

# Hardcoded Google API key
GOOGLE_API_KEY="AIzaSyBdG5TAXTY472625EXAMPLE"

# Hardcoded database credentials
DB_USERNAME="admin"
DB_PASSWORD="mysecretpassword"
DB_HOST="localhost"
DB_PORT=5432

# Hardcoded Slack token
SLACK_TOKEN="xoxb-123456789012-1234567890123-abc123def456"

# Hardcoded private key
PRIVATE_KEY="-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAn1pMVSEDO4EPzQxKgAakFxRgMGiewWZFAktenWo5aMt/OIso
...
-----END RSA PRIVATE KEY-----"

echo "This script contains potential secrets"
