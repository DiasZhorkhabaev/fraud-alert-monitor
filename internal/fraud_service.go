package internal

import "fraud-alert-monitor/database"

func CheckFraud(call Call) {

	if call.Duration < 10 {

		database.InsertAlert(
			database.DB,
			call.PhoneFrom,
			"Short call duration",
			"MEDIUM",
		)
	}

	if call.Status == "FAILED" {

		database.InsertAlert(
			database.DB,
			call.PhoneFrom,
			"Failed call",
			"LOW",
		)
	}
}
