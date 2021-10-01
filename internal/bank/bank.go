package bank

import (
	"cautious-octo-pancake/internal/bank/storage"
	"cautious-octo-pancake/pkg/account"
	"context"
	"fmt"
	"github.com/Rhymond/go-money"
)

type Bank interface {
	OpenAccount(accountID account.Identifier, initialAmount int64) (account.Account, error)
	GetAccount(accountID account.Identifier) (account.Account, error)
	Transfer(origin, destination account.Account, amount int64) error
	AccountDeposit(destination account.Account, amount int64) error
	AccountWithdraw(origin account.Account, amount int64) error
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

func (b *bank) OpenAccount(accountID account.Identifier, initialAmount int64) (account.Account, error) {
	_, err := b.GetAccount(accountID)
	if err != nil {
		if err != ErrAccountNotFound {
			return nil, err
		}

		a := account.NewAccount(accountID, initialAmount, b.currency)
		return a, nil
	}

	return nil, fmt.Errorf("unable to open account with id %v, cause exists account with id informed", accountID)
}

func (b *bank) GetAccount(accountID account.Identifier) (account.Account, error) {
	return b.repository.GetAccount(accountID)
}

func (b *bank) Transfer(origin, destination account.Account, amount int64) error {
	ctx := context.WithValue(context.Background(), "originOriginalBalance", origin.Balance())

	if err := b.AccountWithdraw(origin, amount); err != nil {
		return err
	}

	if err := b.AccountDeposit(destination, amount);  err != nil {
		origin.RollbackBalanceTo(ctx.Value("originOriginalBalance").(int64))
		return err
	}

	return nil
}

func (b *bank) AccountDeposit(destination account.Account, amount int64) error {
	if amount < MinimumDepositAmount {
		return ErrAmountLessThanMinimumDepositAmount
	}
	return destination.Deposit(amount)
}

func (b *bank) AccountWithdraw(origin account.Account, amount int64) error {
	if amount < MinimumWithdrawAmount {
		return ErrAmountLessThanMinimumWithdrawAmount
	}
	return origin.Withdraw(amount)
}

func (b *bank) Reset() {
	b.repository.Reset()
}
