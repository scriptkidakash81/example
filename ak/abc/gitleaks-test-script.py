import os
import requests

# Hardcoded AWS credentials (Vulnerability 1)
AWS_ACCESS_KEY_ID = "AKIAIOSFODNN7EXAMPLE"
AWS_SECRET_ACCESS_KEY = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"

# Hardcoded database credentials (Vulnerability 2)
DB_USERNAME = "admin"
DB_PASSWORD = "password123"

# Using a hardcoded API key (Vulnerability 3)
API_KEY = "1234567890abcdef"

# Using a hardcoded token (Vulnerability 4)
TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaGFuIjoiMjMwfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

# Sending sensitive data over HTTP (Vulnerability 5)
response = requests.get("http://example.com/sensitive-data")

print("Script executed successfully")
