package storage

import (
	"cautious-octo-pancake/pkg/account"
)

type Repository interface {
	InsertAccount(a *account.Account) error
	GetAccounts() []*account.Account
	GetAccount(id account.Identifier) (*account.Account, error)
	CleanData()
}