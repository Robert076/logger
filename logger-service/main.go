package main

import "net/http"

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST methods are allowed on this endpoint.", http.StatusMethodNotAllowed)
			return
		}

	})
}
