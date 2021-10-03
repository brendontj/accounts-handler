package account_handler

import (
	"cautious-octo-pancake/internal/account_handler/storage"
	"cautious-octo-pancake/pkg/account"
	"fmt"
	"github.com/Rhymond/go-money"
)

type AccountHandler interface {
	OpenAccount(accountID account.Identifier, initialAmount int64) (*account.Account, error)
	GetAccount(accountID account.Identifier) (*account.Account, error)
	Transfer(origin, destination *account.Account, amount int64) error
	AccountDeposit(destination *account.Account, amount int64) error
	AccountWithdraw(origin *account.Account, amount int64) error
	Reset()
}

type accountHandler struct {
	repository storage.Repository
	currency   money.Currency
}

func NewAccountHandler(r storage.Repository) AccountHandler {
	return &accountHandler{
		repository: r,
	}
}

func (a *accountHandler) OpenAccount(accountID account.Identifier, initialAmount int64) (*account.Account, error) {
	_, err := a.GetAccount(accountID)
	if err != nil {
		if err != storage.ErrAccountNotFound {
			return nil, err
		}
		return a.createNewAccount(accountID, initialAmount)
	}

	return nil, fmt.Errorf("unable to open account with id %v, cause exists account with id informed", accountID)
}

func (a *accountHandler) GetAccount(accountID account.Identifier) (*account.Account, error) {
	return a.repository.GetAccount(accountID)
}

func (a *accountHandler) Transfer(origin, destination *account.Account, amount int64) error {
	originInitialBalance := origin.Balance()
	if err := a.AccountWithdraw(origin, amount); err != nil {
		return err
	}

	if err := a.AccountDeposit(destination, amount);  err != nil {
		origin.RollbackBalanceTo(originInitialBalance)
		return err
	}

	return nil
}

func (a *accountHandler) AccountDeposit(destination *account.Account, amount int64) error {
	if amount < MinimumDepositAmount {
		return ErrAmountLessThanMinimumDepositAmount
	}
	return destination.Deposit(amount)
}

func (a *accountHandler) AccountWithdraw(origin *account.Account, amount int64) error {
	if amount < MinimumWithdrawAmount {
		return ErrAmountLessThanMinimumWithdrawAmount
	}
	return origin.Withdraw(amount)
}

func (a *accountHandler) Reset() {
	a.repository.CleanData()
}

func (a *accountHandler) createNewAccount(accountID account.Identifier, initialAmount int64) (*account.Account, error) {
	newAccount := account.NewAccount(accountID, initialAmount, a.currency)
	if err := a.repository.InsertAccount(newAccount); err != nil {
		return nil, err
	}
	return newAccount, nil
}
