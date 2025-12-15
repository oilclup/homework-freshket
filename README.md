Git (optional): For cloning the repository
cURL or Postman: For testing API endpoints

Step-by-Step Installation
Option 1: Clone from Git
bash# Clone the repository
git clone https://github.com/oilclup/homework-freshket.git

# Navigate to project directory
cd homework-freshket

# Download dependencies
go mod download

# run program
go run main.go calculator.go
# or
go run .
# or
air

# run tests
go test -v

# run coverage
go test -cover