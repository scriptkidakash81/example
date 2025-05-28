import pandas as pd
import json
import requests
import os
from tabulate import tabulate

# Load JSON data
with open('gitleaks-results.json') as f:
    data = json.load(f)

# Define the fields to extract
fields = ['Match', 'Secret', 'RuleID', 'Entropy', 'File', 'StartLine', 'Commit', 'Author', 'Email', 'Date', 'Fingerprint', 'Link']

# Rename the fields to match the desired output
field_mapping = {
    'Match': 'Finding',
    'StartLine': 'Line'
}

# Extract the fields and rename them
extracted_data = [{field_mapping.get(field, field): item[field_mapping.get(field, field)] for field in fields} for item in data]

# Convert the extracted data to a Pandas DataFrame
df = pd.DataFrame(extracted_data)

# Save to CSV file
df.to_csv('gitleaks-results.csv', index=False)

# Convert the DataFrame to a table format
table = tabulate(df, headers='keys', tablefmt='psql')

# Get environment variables
token = os.environ['GITHUB_TOKEN']
repo_owner = os.environ['GITHUB_REPOSITORY'].split('/')[0]
repo_name = os.environ['GITHUB_REPOSITORY'].split('/')[1]
pr_number = os.environ['PR_NUMBER']

# Comment on PR using GitHub API
url = f'https://api.github.com/repos/{repo_owner}/{repo_name}/issues/{pr_number}/comments'
headers = {'Authorization': f'token {token}'}
data = {'body': f'### Gitleaks Results\n```\n{table}\n```'}
response = requests.post(url, headers=headers, json=data)

if response.status_code == 201:
    print('Comment posted successfully!')
else:
    print('Failed to post comment:', response.text)
