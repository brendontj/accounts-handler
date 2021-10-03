package application

import (
	"cautious-octo-pancake/internal/account_handler/storage"
	"cautious-octo-pancake/pkg/account"
	"net/http"
)

func (a *Application) BalanceHandler(w http.ResponseWriter, r *http.Request){
	accountID := r.URL.Query().Get("account_id")
	if accountID == "" {
		respondWithError(w, http.StatusBadRequest, "account_id query parameter value not found")
		return
	}

	id, done := transformAccountIdentifier(w, accountID)
	if done {
		return
	}

	accountRequested, err := a.AccountHandler.GetAccount(account.Identifier(id))
	if err != nil {
		if err == storage.ErrAccountNotFound {
			respondWithTextValue(w, http.StatusNotFound, 0)
			return
		}
		respondWithError(w, http.StatusInternalServerError, "unable to get account with informed identifier")
		return
	}

	respondWithTextValue(w, http.StatusOK, accountRequested.Balance())
}