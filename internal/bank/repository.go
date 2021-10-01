package bank

import "cautious-octo-pancake/pkg/account"

type Storage interface {
	InsertAccount(a account.Account) error
	Reset()
}

type storage struct {
	accounts []account.Account
}

func NewStorage() Storage {
	return &storage{accounts: make([]account.Account, 0)}
}

func (s *storage) InsertAccount(a account.Account) error {
	panic("implement me")
}

func (s *storage) Reset() {
	s.accounts = make([]account.Account, 0)
}
