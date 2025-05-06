package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	model "github.com/Robert076/logger/logger-service/ping"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Cannot load env file: %v", err)
		return
	}

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST methods are allowed on this endpoint.", http.StatusMethodNotAllowed)
			return
		}

		var newPing model.Ping
		if err := json.NewDecoder(r.Body).Decode(&newPing); err != nil {
			http.Error(w, "Invalid JSON format.", http.StatusBadRequest)
			return
		}

		if newPing.Message == "" {
			http.Error(w, "Ping message cannot be empty.", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(newPing); err != nil {
			http.Error(w, "Error encoding message.", http.StatusInternalServerError)
		}
	})

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatalf("Cannot start server: %v", err)
		return
	}
}
