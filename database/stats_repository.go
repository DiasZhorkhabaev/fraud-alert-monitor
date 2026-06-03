package database

import "database/sql"

func GetStats(db *sql.DB) (map[string]interface{}, error) {

	var totalCalls int
	var totalAlerts int
	var failedCalls int

	err := db.QueryRow(
		"SELECT COUNT(*) FROM calls",
	).Scan(&totalCalls)

	if err != nil {
		return nil, err
	}

	err = db.QueryRow(
		"SELECT COUNT(*) FROM fraud_alerts",
	).Scan(&totalAlerts)

	if err != nil {
		return nil, err
	}

	err = db.QueryRow(
		"SELECT COUNT(*) FROM calls WHERE status = 'FAILED'",
	).Scan(&failedCalls)

	if err != nil {
		return nil, err
	}

	stats := map[string]interface{}{
		"total_calls":  totalCalls,
		"total_alerts": totalAlerts,
		"failed_calls": failedCalls,
	}

	return stats, nil
}
