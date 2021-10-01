package storage

import (
	"cautious-octo-pancake/internal/bank"
	"cautious-octo-pancake/pkg/account"
	"fmt"
)

type Repository interface {
	InsertAccount(a account.Account) error
	GetAccounts() []account.Account
	GetAccount(id account.Identifier) (account.Account, error)
	Reset()
}

type repository struct {
	accounts []account.Account
}

func NewRepository() Repository {
	return &repository{accounts: make([]account.Account, 0)}
}

func (r *repository) InsertAccount(a account.Account) error {
	if r.existAccountWithID(a.ID()) {
		return fmt.Errorf("unable to store new account with id %v, cause: exists account with the same id " +
			"registered", a.ID())
	}

	r.accounts = append(r.accounts, a)
	return nil
}

func (r *repository) GetAccounts() []account.Account {
	return r.accounts
}

func (r *repository) GetAccount(id account.Identifier) (account.Account, error) {
	for _, a := range r.accounts {
		if a.ID() == id {
			return a, nil
		}
	}
	return nil, bank.ErrAccountNotFound
}

func (r *repository) existAccountWithID(id account.Identifier) bool {
	for _, a := range r.accounts {
		if a.ID() == id {
			return true
		}
	}
	return false
}

func (r *repository) Reset() {
	r.accounts = make([]account.Account, 0)
}
