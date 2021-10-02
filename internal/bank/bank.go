package bank

import (
	"cautious-octo-pancake/internal/bank/storage"
	"cautious-octo-pancake/pkg/account"
	"fmt"
	"github.com/Rhymond/go-money"
)

type Bank interface {
	OpenAccount(accountID account.Identifier, initialAmount int64) (*account.Account, error)
	GetAccount(accountID account.Identifier) (*account.Account, error)
	Transfer(origin, destination *account.Account, amount int64) error
	AccountDeposit(destination *account.Account, amount int64) error
	AccountWithdraw(origin *account.Account, amount int64) error
	Reset()
}

type bank struct {
	repository storage.Repository
	currency   money.Currency
}

func NewBank(r storage.Repository) Bank {
	return &bank{
		repository: r,
	}
}

func (b *bank) OpenAccount(accountID account.Identifier, initialAmount int64) (*account.Account, error) {
	_, err := b.GetAccount(accountID)
	if err != nil {
		if err != storage.ErrAccountNotFound {
			return nil, err
		}
		return b.createNewAccount(accountID, initialAmount)
	}

	return nil, fmt.Errorf("unable to open account with id %v, cause exists account with id informed", accountID)
}

func (b *bank) GetAccount(accountID account.Identifier) (*account.Account, error) {
	return b.repository.GetAccount(accountID)
}

func (b *bank) Transfer(origin, destination *account.Account, amount int64) error {
	originInitialBalance := origin.Balance()
	if err := b.AccountWithdraw(origin, amount); err != nil {
		return err
	}

	if err := b.AccountDeposit(destination, amount);  err != nil {
		origin.RollbackBalanceTo(originInitialBalance)
		return err
	}

	return nil
}

func (b *bank) AccountDeposit(destination *account.Account, amount int64) error {
	if amount < MinimumDepositAmount {
		return ErrAmountLessThanMinimumDepositAmount
	}
	return destination.Deposit(amount)
}

func (b *bank) AccountWithdraw(origin *account.Account, amount int64) error {
	if amount < MinimumWithdrawAmount {
		return ErrAmountLessThanMinimumWithdrawAmount
	}
	return origin.Withdraw(amount)
}

func (b *bank) Reset() {
	b.repository.CleanData()
}

func (b *bank) createNewAccount(accountID account.Identifier, initialAmount int64) (*account.Account, error) {
	a := account.NewAccount(accountID, initialAmount, b.currency)
	if err := b.repository.InsertAccount(a); err != nil {
		return nil, err
	}
	return a, nil
}
