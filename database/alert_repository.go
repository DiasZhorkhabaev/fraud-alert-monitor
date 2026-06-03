package database

import "database/sql"

func InsertAlert(
	db *sql.DB,
	phoneNumber string,
	reason string,
	riskLevel string,
) error {

	query := `
	INSERT INTO fraud_alerts
	(phone_number, reason, risk_level)
	VALUES ($1, $2, $3)
	`

	_, err := db.Exec(
		query,
		phoneNumber,
		reason,
		riskLevel,
	)

	return err
}

func GetAlerts(db *sql.DB) ([]map[string]interface{}, error) {

	rows, err := db.Query(`
		SELECT id,
			   phone_number,
			   reason,
			   risk_level
		FROM fraud_alerts
		ORDER BY id DESC
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var alerts []map[string]interface{}

	for rows.Next() {

		var id int
		var phoneNumber string
		var reason string
		var riskLevel string

		err := rows.Scan(
			&id,
			&phoneNumber,
			&reason,
			&riskLevel,
		)

		if err != nil {
			return nil, err
		}

		alert := map[string]interface{}{
			"id":           id,
			"phone_number": phoneNumber,
			"reason":       reason,
			"risk_level":   riskLevel,
		}

		alerts = append(alerts, alert)
	}

	return alerts, nil
}
