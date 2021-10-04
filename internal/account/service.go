package account

import (
	"cautious-octo-pancake/internal/account/storage"
	"cautious-octo-pancake/internal/database"
	"cautious-octo-pancake/pkg/account"
	"fmt"
	"github.com/Rhymond/go-money"
)

type Service interface {
	OpenAccount(accountID account.Identifier, initialAmount int64) (*account.Account, error)
	GetAccount(accountID account.Identifier) (*account.Account, error)
	Transfer(origin, destination *account.Account, amount int64) error
	AccountDeposit(destination *account.Account, amount int64) error
	AccountWithdraw(origin *account.Account, amount int64) error
	Reset()
}

type accountService struct {
	repository storage.Repository
	currency   money.Currency
}

func NewAccountHandler(r storage.Repository) Service {
	return &accountService{
		repository: r,
	}
}

func (a *accountService) OpenAccount(accountID account.Identifier, initialAmount int64) (*account.Account, error) {
	_, err := a.GetAccount(accountID)
	if err != nil {
		if err != database.ErrAccountNotFound {
			return nil, err
		}
		return a.createNewAccount(accountID, initialAmount)
	}

	return nil, fmt.Errorf("unable to open account with id %v, cause exists account with id informed", accountID)
}

func (a *accountService) GetAccount(accountID account.Identifier) (*account.Account, error) {
	return a.repository.GetAccount(accountID)
}

func (a *accountService) Transfer(origin, destination *account.Account, amount int64) error {
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

func (a *accountService) AccountDeposit(destination *account.Account, amount int64) error {
	if amount < MinimumDepositAmount {
		return ErrAmountLessThanMinimumDepositAmount
	}
	return destination.Deposit(amount)
}

func (a *accountService) AccountWithdraw(origin *account.Account, amount int64) error {
	if amount < MinimumWithdrawAmount {
		return ErrAmountLessThanMinimumWithdrawAmount
	}
	return origin.Withdraw(amount)
}

func (a *accountService) Reset() {
	a.repository.CleanData()
}

func (a *accountService) createNewAccount(accountID account.Identifier, initialAmount int64) (*account.Account, error) {
	newAccount := account.NewAccount(accountID, initialAmount, a.currency)
	if err := a.repository.InsertAccount(newAccount); err != nil {
		return nil, err
	}
	return newAccount, nil
}
