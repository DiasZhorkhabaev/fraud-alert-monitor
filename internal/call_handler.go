package internal

import (
	"encoding/json"
	"log"
	"net/http"

	"fraud-alert-monitor/database"
)

func (app *App) CreateCallHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	var call Call

	err := json.NewDecoder(r.Body).Decode(&call)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	err = database.InsertCall(
		app.DB,
		call.PhoneFrom,
		call.PhoneTo,
		call.Country,
		call.Duration,
		call.Status,
	)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	log.Printf(
		"Call created: from=%s to=%s country=%s duration=%d status=%s",
		call.PhoneFrom,
		call.PhoneTo,
		call.Country,
		call.Duration,
		call.Status,
	)

	CheckFraud(app.DB, call)

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(call)
}
