package database

import (
	"database/sql"
)

func InsertCall(
	db *sql.DB,
	phoneFrom string,
	phoneTo string,
	country string,
	duration int,
	status string,
) error {

	query := `
	INSERT INTO calls
	(phone_from, phone_to, country, duration, status)
	VALUES ($1, $2, $3, $4, $5)
	`

	_, err := db.Exec(
		query,
		phoneFrom,
		phoneTo,
		country,
		duration,
		status,
	)

	return err
}

func GetCalls(db *sql.DB) ([]map[string]interface{}, error) {

	rows, err := db.Query(`
		SELECT id,
			   phone_from,
			   phone_to,
			   country,
			   duration,
			   status
		FROM calls
		ORDER BY id DESC
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var calls []map[string]interface{}

	for rows.Next() {

		var id int
		var phoneFrom string
		var phoneTo string
		var country string
		var duration int
		var status string

		err := rows.Scan(
			&id,
			&phoneFrom,
			&phoneTo,
			&country,
			&duration,
			&status,
		)

		if err != nil {
			return nil, err
		}

		call := map[string]interface{}{
			"id":         id,
			"phone_from": phoneFrom,
			"phone_to":   phoneTo,
			"country":    country,
			"duration":   duration,
			"status":     status,
		}

		calls = append(calls, call)
	}

	return calls, nil
}
