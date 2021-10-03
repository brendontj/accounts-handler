package application

import (
	"cautious-octo-pancake/application/dto"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const (
	EventTypeDeposit = "deposit"
	EventTypeWithdraw = "withdraw"
	EventTypeTransfer = "transfer"
)

func transformAccountIdentifier(w http.ResponseWriter, identifier string) (int, bool) {
	accountID, err := strconv.Atoi(identifier)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "account identifier needs to be a number")
		return 0, true
	}
	return accountID, false
}

func decodeRequestBody(w http.ResponseWriter, r *http.Request) (*dto.Event, bool) {
	var e dto.Event
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&e); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request payload")
		return nil, true
	}
	return &e, false
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithTextValue(w http.ResponseWriter, code int, value interface{}) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(code)
	_ , _ = w.Write([]byte(fmt.Sprintf("%v", value)))
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}
