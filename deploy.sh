#!/bin/bash

# Step 1: Git add and commit with auto message
timestamp=$(date +"%Y-%m-%d %T")
git add .
git commit -m "Auto commit at $timestamp"

# Step 2: Push to main repo
git push origin main

# Step 3: Connect SSH
sshpass -p "RLjh&%R2BdFS2geJkh5#" ssh root@77.68.21.74

# Step 4: Update the repo in the folder
ssh root@77.68.21.74 "cd realestate-2-backend && git pull"

# Step 5: Build the Golang code
ssh root@77.68.21.74 "cd realestate-2-backend && go build -v -o main ./cmd"

# Step 6: Move the binary file
ssh root@77.68.21.74 "mv realestate-2-backend/main go/realestate-backend"

# Step 7: Restart the service
ssh root@77.68.21.74 "systemctl restart realestate-backend.service"