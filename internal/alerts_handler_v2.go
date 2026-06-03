package internal

import (
	"encoding/json"
	"net/http"

	"fraud-alert-monitor/database"
)

func (app *App) GetAlertsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	alerts, err := database.GetAlerts(app.DB)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(alerts)
}
