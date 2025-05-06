package main

import (
	"encoding/json"
	"log"
	"net/http"

	model "github.com/Robert076/logger/logger-service/ping"
)

func main() {
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

	if err := http.ListenAndServe(":0607", nil); err != nil {
		log.Fatalf("Cannot start server: %v", err)
		return
	}
}
