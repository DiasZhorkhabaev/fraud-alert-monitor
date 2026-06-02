package internal

import (
	"encoding/json"
	"net/http"

	"fraud-alert-monitor/database"
)

func CreateCallHandler(w http.ResponseWriter, r *http.Request) {
	var call Call

	err := json.NewDecoder(r.Body).Decode(&call)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.InsertCall(
		database.DB,
		call.PhoneFrom,
		call.PhoneTo,
		call.Country,
		call.Duration,
		call.Status,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(call)
}
