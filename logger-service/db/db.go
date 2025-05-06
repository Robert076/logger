package db

import (
	"database/sql"
	"fmt"
	"os"

	model "github.com/Robert076/logger/logger-service/ping"
	_ "github.com/lib/pq"
)

func GetPingsFromDatabase() []model.Ping {
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
		panic(err)
	}

	// nolint:errcheck
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT Id, Message FROM Ping")

	if err != nil {
		panic(err)
	}

	// nolint:errcheck
	defer rows.Close()

	var pings []model.Ping

	for rows.Next() {
		var newPing model.Ping
		err := rows.Scan(&newPing.Id, &newPing.Message)
		if err != nil {
			panic(err)
		}
		pings = append(pings, newPing)
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	return pings
}
