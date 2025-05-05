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
