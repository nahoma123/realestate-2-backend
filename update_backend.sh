# Step 4: Update the repo in the folder
cd /home/nahom/Desktop/upwork/real_estate_backend/realestate-2-backend
git pull

# Step 5: Build the Golang code
go build -v -o main ./cmd

# Step 6: Move the binary file
mv main go/realestate-backend

# Step 7: Restart the service
systemctl restart realestate-backend.service