package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func LoadDB() {
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

	fmt.Println(connStr)
	// Define the connection parameters
	//connStr := "user=postgres dbname=csi_conversion sslmode=disable"
	// Replace "username" and "mydb" with your PostgreSQL username and database name

	//connStr := "user=postgres dbname=csi_conversion sslmode=disable port=5433"

	//connStr1 := "user=postgres password=root dbname=csi_conversion sslmode=disable port=5433"
	//fmt.Println(connStr1)

	fmt.Println("aaaaaaaaaaaaaaaaa")
	// Open a database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	DB = db
}
