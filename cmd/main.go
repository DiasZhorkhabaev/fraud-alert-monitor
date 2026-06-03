package main

import (
	"fmt"
	"net/http"

	"fraud-alert-monitor/database"
	"fraud-alert-monitor/internal"
)

func main() {
	fmt.Println("Fraud Alert Monitor started")

	db, err := database.Connect()

	if err != nil {
		fmt.Println("Database connection error:", err)
		return
	}

	app := &internal.App{
		DB: db,
	}

	fmt.Println("Server started on :8080")

	defer db.Close()

	http.HandleFunc("/health", internal.HealthHandler)
	http.HandleFunc("/alerts", app.GetAlertsHandler)
	http.HandleFunc("/stats", app.GetStatsHandler)

	http.HandleFunc("/calls", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			app.CreateCallHandler(w, r)
			return
		}

		if r.Method == http.MethodGet {
			app.GetCallsHandler(w, r)
			return
		}

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Server error:", err)
	}
}
