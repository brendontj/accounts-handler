package bank

import (
	"cautious-octo-pancake/pkg/account"
)

type Bank interface {
	Reset()
	GetAccount(accountID int64) (account.Account, error)
	Transfer(origin, destination account.Account, amount int64) error
	AccountDeposit(destination account.Account, amount int64) error
	AccountWithdraw(origin account.Account, amount int64) error
}

type bank struct {
	repo Storage
}

func NewBank(s Storage) Bank {
	return &bank{
		repo: s,
	}
}

func (b *bank) GetAccount(accountID int64) (account.Account, error) {
	panic("implement me")
}

func (b *bank) Transfer(origin, destination account.Account, amount int64) error {
	panic("implement me")
}

func (b *bank) AccountDeposit(destination account.Account, amount int64) error {
	panic("implement me")
}

func (b *bank) AccountWithdraw(origin account.Account, amount int64) error {
	panic("implement me")
}

func (b *bank) Reset() {
	b.repo.Reset()
}
