package application

import (
	"cautious-octo-pancake/internal/application/dto"
	"cautious-octo-pancake/internal/database"
	"cautious-octo-pancake/pkg/account"
	"net/http"
)

func (a *Api) EventHandler(w http.ResponseWriter, r *http.Request){
	defer func() {
		_ = r.Body.Close()
	}()

	e, done := decodeRequestBody(w, r)
	if done {
		return
	}

	switch e.EventType {
	case EventTypeDeposit:
		a.handleDepositEvent(w, e)
	case EventTypeTransfer:
		a.handleTransferEvent(w, e)
	case EventTypeWithdraw:
		a.handleWithdrawEvent(w, e)
	default:
		respondWithError(w, http.StatusBadRequest, "event type not implemented")
	}
}

func (a *Api) handleDepositEvent(w http.ResponseWriter, e *dto.Event) {
	accountID, done := transformAccountIdentifier(w, *e.Destination)
	if done {
		return
	}
	acc, err := a.Service.GetAccount(account.Identifier(accountID))
	if err != nil {
		if err == database.ErrAccountNotFound {
			a.openAccountWithInitialBalance(w, accountID, *e)
			return
		}
		respondWithError(w, http.StatusInternalServerError, "unable to get account with informed identifier")
		return
	}

	a.accountDeposit(w, acc, *e)
}

func (a *Api) handleWithdrawEvent(w http.ResponseWriter, e *dto.Event) {
	accountID, done := transformAccountIdentifier(w, *e.Origin)
	if done {
		return
	}
	acc, err := a.Service.GetAccount(account.Identifier(accountID))
	if err != nil {
		if err == database.ErrAccountNotFound {
			respondWithTextValue(w, http.StatusNotFound, 0)
			return
		}
		respondWithError(w, http.StatusInternalServerError, "unable to get account with informed identifier")
		return
	}

	a.accountWithdraw(w, acc, *e)
}

func (a *Api) handleTransferEvent(w http.ResponseWriter, e *dto.Event) {
	accountIdentifierOrigin, done := transformAccountIdentifier(w, *e.Origin)
	if done {
		return
	}

	accountIdentifierDestination, done := transformAccountIdentifier(w, *e.Destination)
	if done {
		return
	}
	originAccount, err := a.Service.GetAccount(account.Identifier(accountIdentifierOrigin))
	if err != nil {
		if err == database.ErrAccountNotFound {
			respondWithTextValue(w, http.StatusNotFound, 0)
			return
		}
		respondWithError(w, http.StatusInternalServerError, "unable to get account with informed identifier")
		return
	}

	destinationAccount, err := a.Service.GetAccount(account.Identifier(accountIdentifierDestination))
	if err != nil {
		if err == database.ErrAccountNotFound {
			destinationAccount, err = a.Service.OpenAccount(account.Identifier(accountIdentifierDestination), 0)
			if err != nil {
				respondWithError(w, http.StatusInternalServerError, "unable to open destination account")
				return
			}
		} else {
			respondWithError(w, http.StatusInternalServerError, "unable to get account with informed identifier")
			return
		}
	}

	if err := a.Service.Transfer(originAccount, destinationAccount, e.Amount); err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to transfer amount from existing account")
		return
	}

	respondWithJSON(w, http.StatusCreated,
		dto.Response{
			Origin: &dto.AccountResponse{
				ID:      originAccount.ID().String(),
				Balance: originAccount.Balance(),
			},
			Destination: &dto.AccountResponse{
				ID:      destinationAccount.ID().String(),
				Balance: destinationAccount.Balance(),
			},
		})
}

func (a *Api) openAccountWithInitialBalance(w http.ResponseWriter, accountID int, e dto.Event) {
	openedAccount, err := a.Service.OpenAccount(account.Identifier(accountID), e.Amount)
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

func (a *Api) accountDeposit(w http.ResponseWriter, acc *account.Account, e dto.Event) {
	if err := a.Service.AccountDeposit(acc, e.Amount); err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to deposit in the account")
		return
	}
	respondWithJSON(w, http.StatusCreated, dto.Response{Destination: &dto.AccountResponse{
		ID:      acc.ID().String(),
		Balance: acc.Balance(),
	},
		Origin: nil})
}

func (a *Api) accountWithdraw(w http.ResponseWriter, acc *account.Account, e dto.Event) {
	if err := a.Service.AccountWithdraw(acc, e.Amount); err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to withdraw from account")
		return
	}
	respondWithJSON(w, http.StatusCreated, dto.Response{Origin: &dto.AccountResponse{
		ID:      acc.ID().String(),
		Balance: acc.Balance(),
	},
		Destination: nil})
}