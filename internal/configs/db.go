package configs

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func GetDB() (*sql.DB, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Create table and index if not exists
	err = createTableAndIndex(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createTableAndIndex(db *sql.DB) error {
	// Create employees table if not exists
	createTableQuery := `CREATE TABLE IF NOT EXISTS Employee (
		id INT AUTO_INCREMENT PRIMARY KEY,
		employee_name VARCHAR(255) NOT NULL,
		employee_salary INT NOT NULL,
		employee_age INT NOT NULL,
		profile_image VARCHAR(255) NOT NULL
	);`
	_, err := db.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("error creating table: %v", err)
	}

	// Add full-text index on employee_name if not exists
	createIndexQuery := `ALTER TABLE Employee ADD FULLTEXT(employee_name);`
	_, err = db.Exec(createIndexQuery)
	if err != nil {
		return fmt.Errorf("error creating full-text index: %v", err)
	}

	return nil
}
