#!/bin/bash

# Step 1: Git add and commit with auto message
timestamp=$(date +"%Y-%m-%d %T")
git add .
git commit -m "Auto commit at $timestamp"

# Step 2: Push to main repo
git push origin main