#!/bin/bash
set -e

# Set the base branch and pull request branch
BASE_BRANCH="main"
PR_BRANCH="${GITHUB_REF##*/}"

# Get the list of changed files between the base branch and the pull request branch
git diff --name-only ${BASE_BRANCH}...${PR_BRANCH} > changed_files.txt

# Run Gosec on the changed files
gosec -include=changed_files.txt . > gosec_results.txt

# Parse the Gosec results to highlight new vulnerabilities introduced in the pull request
NEW_VULNS=()
while IFS= read -r line; do
  if [[ $line =~ "severity" ]]; then
    NEW_VULNS+=("$line")
  fi
done < gosec_results.txt

# Print the new vulnerabilities
if [ ${#NEW_VULNS[@]} -gt 0 ]; then
  echo "New vulnerabilities introduced in the pull request:"
  for vuln in "${NEW_VULNS[@]}"; do
    echo "$vuln"
  done
else
  echo "No new vulnerabilities introduced in the pull request."
fi

# Post a comment on the pull request with the Gosec results
COMMENT="Gosec Results:\n"
if [ ${#NEW_VULNS[@]} -gt 0 ]; then
  COMMENT+="New vulnerabilities introduced in the pull request:\n"
  for vuln in "${NEW_VULNS[@]}"; do
    COMMENT+="$vuln\n"
  done
else
  COMMENT+="No new vulnerabilities introduced in the pull request."
fi

curl -X POST \
  -H "Authorization: token ${GITHUB_TOKEN}" \
  -H "Accept: application/vnd.github.v3+json" \
  -d "{\"body\":\"$COMMENT\"}" \
  "https://api.github.com/repos/${GITHUB_REPOSITORY}/issues/${GITHUB_PR_NUMBER}/comments"
