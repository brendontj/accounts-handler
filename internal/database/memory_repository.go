package database

import (
	"cautious-octo-pancake/internal/account/storage"
	"cautious-octo-pancake/pkg/account"
	"fmt"
)

type memoryRepository struct {
	accounts []*account.Account
}

func NewMemoryRepository() storage.Repository {
	return &memoryRepository{accounts: make([]*account.Account, 0)}
}

func (r *memoryRepository) InsertAccount(a *account.Account) error {
	if r.existAccountWithID(a.ID()) {
		return fmt.Errorf("unable to store new account with id %v, cause: exists account with the same id " +
			"registered", a.ID())
	}

	r.accounts = append(r.accounts, a)
	return nil
}

func (r *memoryRepository) GetAccounts() []*account.Account {
	return r.accounts
}

func (r *memoryRepository) GetAccount(id account.Identifier) (*account.Account, error) {
	for _, a := range r.accounts {
		if a.ID() == id {
			return a, nil
		}
	}
	return nil, ErrAccountNotFound
}

func (r *memoryRepository) existAccountWithID(id account.Identifier) bool {
	for _, a := range r.accounts {
		if a.ID() == id {
			return true
		}
	}
	return false
}

func (r *memoryRepository) CleanData() {
	r.accounts = make([]*account.Account, 0)
}
