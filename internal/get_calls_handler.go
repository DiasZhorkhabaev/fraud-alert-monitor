package internal

import (
	"encoding/json"
	"net/http"

	"fraud-alert-monitor/database"
)

func GetCallsHandler(w http.ResponseWriter, r *http.Request) {

	calls, err := database.GetCalls(database.DB)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(calls)
}
