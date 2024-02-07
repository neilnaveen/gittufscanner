#!/bin/bash

# Rest of your script...

# Define the repository location
REPO_PATH="$(pwd)/$1/$2"

# Change to the repository directory
cd "$REPO_PATH"

# Check if the 'main' branch exists in the remote repository
if git ls-remote --heads origin main | grep -q 'refs/heads/main'; then
    BRANCH="main"
else
    BRANCH="master"
fi

# Fetch the latest changes from the remote repository
git pull origin "$BRANCH"

git fetch origin --tags --force

# Rebase the specified branch
git rebase "origin/$BRANCH"

cd ../../
