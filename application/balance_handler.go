package application

import (
	"cautious-octo-pancake/internal/bank/storage"
	"cautious-octo-pancake/pkg/account"
	"net/http"
	"strconv"
)

func (a *Application) BalanceHandler(w http.ResponseWriter, r *http.Request){
	accountID := r.URL.Query().Get("account_id")
	if accountID == "" {
		respondWithError(w, http.StatusBadRequest, "account_id query parameter value not found")
		return
	}

	id, err := strconv.Atoi(accountID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "account_id needs to be an integer number")
		return
	}

	accountRequested, err := a.Bank.GetAccount(account.Identifier(id))
	if err != nil {
		if err == storage.ErrAccountNotFound {
			respondWithTextValue(w, http.StatusNotFound, 0)
			return
		}
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithTextValue(w, http.StatusOK, accountRequested.Balance())
}