package bank

import (
	"cautious-octo-pancake/pkg/account"
	"fmt"
)

type Storage interface {
	InsertAccount(a account.Account) error
	GetAccounts() []account.Account
	GetAccount(id account.Identifier) (account.Account, error)
	Reset()
}

type storage struct {
	accounts []account.Account
}

func NewStorage() Storage {
	return &storage{accounts: make([]account.Account, 0)}
}

func (s *storage) InsertAccount(a account.Account) error {
	if s.existAccountWithID(a.ID()) {
		return fmt.Errorf("unable to store new account with id %v, cause: exists account with the same id " +
			"registered", a.ID())
	}

	s.accounts = append(s.accounts, a)
	return nil
}

func (s *storage) GetAccounts() []account.Account {
	return s.accounts
}

func (s *storage) GetAccount(id account.Identifier) (account.Account, error) {
	for _, a := range s.accounts {
		if a.ID() == id {
			return a, nil
		}
	}
	return nil, ErrAccountNotFound
}

func (s *storage) existAccountWithID(id account.Identifier) bool {
	for _, a := range s.accounts {
		if a.ID() == id {
			return true
		}
	}
	return false
}

func (s *storage) Reset() {
	s.accounts = make([]account.Account, 0)
}
