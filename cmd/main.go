package main

import (
	"fmt"

	"fraud-alert-monitor/database"
)

func main() {
	fmt.Println("Fraud Alert Monitor started")

	db, err := database.Connect()

	if err != nil {
		fmt.Println("Database connection error:", err)
		return
	}

	defer db.Close()

	fmt.Println("Application is ready")
}
