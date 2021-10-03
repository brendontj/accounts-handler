package application

import (
	"cautious-octo-pancake/internal/bank/storage"
	"cautious-octo-pancake/pkg/account"
	"encoding/json"
	"net/http"
	"strconv"
)

type EventDto struct {
	EventType string `json:"type"`
	Destination *string `json:"destination,omitempty"`
	Origin *string `json:"origin,omitempty"`
	Amount int64 `json:"amount"`
}

type ResponseDto struct {
	Destination *AccountResponseDto `json:"destination,omitempty"`
	Origin *AccountResponseDto `json:"origin,omitempty"`
}

type AccountResponseDto struct {
	ID string `json:"id"`
	Balance int64 `json:"balance"`
}

func (a *Application) EventHandler(w http.ResponseWriter, r *http.Request){
	var e EventDto
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&e); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer func() {
		_ = r.Body.Close()
	}()

	switch e.EventType {
	case EventTypeDeposit:
		accountID, err := strconv.Atoi(*e.Destination)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "destination needs to be a number")
			return
		}
		acc, err := a.Bank.GetAccount(account.Identifier(accountID))
		if err != nil {
			if err == storage.ErrAccountNotFound {
				acc, err = a.Bank.OpenAccount(account.Identifier(accountID), e.Amount)
				if err != nil {
					respondWithError(w, http.StatusInternalServerError, "unable to open account")
					return
				}
				respondWithJSON(w, http.StatusCreated, ResponseDto{Destination: &AccountResponseDto{
					ID:      acc.ID().String(),
					Balance: acc.Balance(),
				},
				Origin: nil})
				return
			}
		}
		if err := a.Bank.AccountDeposit(acc,e.Amount); err != nil {
			respondWithError(w, http.StatusInternalServerError, "unable to deposit in the account")
			return
		}
		respondWithJSON(w, http.StatusCreated, ResponseDto{Destination: &AccountResponseDto{
			ID:      acc.ID().String(),
			Balance: acc.Balance(),
		},
			Origin: nil})
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