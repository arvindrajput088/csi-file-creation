package main

import (
	"csi-conversion/processline"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// LoadEnvironment loads environment variables from the .env file.
func LoadEnvironment() error {
	err := godotenv.Load()
	return err
}

// CreateDBConnection creates and returns a database connection.
func CreateDBConnection() (*sql.DB, error) {
	// Retrieve configuration data from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSL := os.Getenv("DB_SSL")
	dbPort := os.Getenv("DB_PORT")

	// Construct the connection string
	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=%s port=%s",
		dbUser, dbPassword, dbName, dbSSL, dbPort,
	)

	// Open a database connection
	db, err := sql.Open("postgres", connStr)
	return db, err
}

// Main function
func main() {
	//id := 1 // Insert a row into hot_file table and get the id
	fmt.Println("Hello, this is my new project CSN Conversion")

	// Load environment variables
	if err := LoadEnvironment(); err != nil {
		log.Fatal(err)
	}

	// Create a database connection
	db, err := CreateDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Open the file for reading
	/*file, err := os.Open("hotfiles/1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		processline.ProcessLine(db, id, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}*/
	processline.CSIFileCreation()
}
