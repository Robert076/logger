package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Robert076/logger/logger-service/db"
	model "github.com/Robert076/logger/logger-service/ping"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed on this endpoint.", http.StatusMethodNotAllowed)
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

	http.HandleFunc("/pings", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Only GET method is allowed on this endpoint", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		pings := db.GetPingsFromDatabase()

		if err := json.NewEncoder(w).Encode(pings); err != nil {
			http.Error(w, "Failed to encode pings", http.StatusInternalServerError)
			return
		}
	})

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatalf("Cannot start server: %v", err)
		return
	}
}
