package db

import (
	"database/sql"
	"fmt"
	"os"

	model "github.com/Robert076/logger/logger-service/ping"
	_ "github.com/lib/pq"
)

func GetPingsFromDatabase() ([]model.Ping, error) {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging the database: %v", err)
	}

	rows, err := db.Query("SELECT Id, Message, CreatedAt FROM Ping")
	if err != nil {
		return nil, fmt.Errorf("error querying database: %v", err)
	}
	defer rows.Close()

	var pings []model.Ping
	for rows.Next() {
		var newPing model.Ping
		err := rows.Scan(&newPing.Id, &newPing.Message, &newPing.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		pings = append(pings, newPing)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return pings, nil
}

func InsertPing(message string) (int, error) {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return 0, fmt.Errorf("error opening database connection: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return 0, fmt.Errorf("error pinging the database: %v", err)
	}

	query := "INSERT INTO Ping (Message) VALUES ($1) RETURNING Id"
	var id int
	err = db.QueryRow(query, message).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error inserting ping: %v", err)
	}

	return id, nil
}
