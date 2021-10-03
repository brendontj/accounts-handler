package application

import (
	"encoding/json"
	"net/http"
)

const (
	EventTypeDeposit = "deposit"
	EventTypeWithdraw = "withdraw"
	EventTypeTransfer = "transfer"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithTextValue(w http.ResponseWriter, code int, value interface{}) {
	w.Header().Set("Content-Type", "plain/text")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(value)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}
