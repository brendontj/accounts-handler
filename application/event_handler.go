package application

import (
	"cautious-octo-pancake/application/dto"
	"cautious-octo-pancake/internal/bank/storage"
	"cautious-octo-pancake/pkg/account"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (a *Application) EventHandler(w http.ResponseWriter, r *http.Request){
	defer func() {
		_ = r.Body.Close()
	}()

	e, done := decodeRequestBody(w, r)
	if done {
		return
	}

	switch e.EventType {
	case EventTypeDeposit:
		accountID, done := transformAccountIdentifier(w, e.Destination)
		if done {
			return
		}
		acc, err := a.Bank.GetAccount(account.Identifier(accountID))
		if err != nil {
			if err == storage.ErrAccountNotFound {
				a.openAccountWithInitialBalance(w, accountID, *e)
				return
			}
			respondWithError(w, http.StatusInternalServerError, "Unable to get account with informed identifier")
			return
		}
		a.accountDeposit(w, acc, *e)
		return

	case EventTypeTransfer:
		panic("implement me")
	case EventTypeWithdraw:
		panic("implement me")
	default:
		respondWithError(w, http.StatusBadRequest, "Event type not implemented")
		return
	}

}

func (a *Application) openAccountWithInitialBalance(w http.ResponseWriter, accountID int, e dto.Event) {
	openedAccount, err := a.Bank.OpenAccount(account.Identifier(accountID), e.Amount)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to open account")
		return
	}
	respondWithJSON(w, http.StatusCreated, dto.Response{Destination: &dto.AccountResponse{
		ID:      openedAccount.ID().String(),
		Balance: openedAccount.Balance(),
	},
		Origin: nil})
	return
}

func (a *Application) accountDeposit(w http.ResponseWriter, acc *account.Account, e dto.Event) {
	if err := a.Bank.AccountDeposit(acc, e.Amount); err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to deposit in the account")
		return
	}
	respondWithJSON(w, http.StatusCreated, dto.Response{Destination: &dto.AccountResponse{
		ID:      acc.ID().String(),
		Balance: acc.Balance(),
	},
		Origin: nil})
}

func transformAccountIdentifier(w http.ResponseWriter, identifier interface {}) (int, bool) {
	accountID, err := strconv.Atoi(fmt.Sprintf("%v", identifier))
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
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return nil, true
	}
	return &e, false
}