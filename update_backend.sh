# Step 4: Update the repo in the folder
cd ./realestate-2-backend
git pull
echo "Repo updated successfully"

# Step 5: Build the Golang code
go build -v -o main ./cmd
echo "Golang code built successfully"

# Step 6: Move the binary file
mv main ../go/realestate-backend
echo "Binary file moved successfully"

# Step 7: Restart the service
systemctl restart realestate-backend.service
echo "Service restarted successfully"

# Step 8: View service logs
journalctl -u realestate-backend.service -f
