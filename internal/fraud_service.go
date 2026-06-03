package internal

import (
	"database/sql"
	"log"

	"fraud-alert-monitor/database"
)

func CheckFraud(db *sql.DB, call Call) {

	if call.Duration < 10 {

		database.InsertAlert(
			db,
			call.PhoneFrom,
			"Short call duration",
			"MEDIUM",
		)

		log.Printf(
			"Fraud alert generated: phone=%s reason=%s",
			call.PhoneFrom,
			"Short call duration",
		)
	}

	if call.Status == "FAILED" {

		database.InsertAlert(
			db,
			call.PhoneFrom,
			"Failed call",
			"LOW",
		)

		log.Printf(
			"Fraud alert generated: phone=%s reason=%s",
			call.PhoneFrom,
			"Short call duration",
		)
	}
}
