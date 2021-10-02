package memory_storage_test

import (
	"cautious-octo-pancake/internal/bank/storage"
	"cautious-octo-pancake/pkg/account"
	"github.com/Rhymond/go-money"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInsertAccountInMemoryRepository(t *testing.T) {
	t.Parallel()

	s := storage.NewMemoryRepository()
	currency := *money.GetCurrency("BRL")
	a := account.NewAccount(1, 100, currency)

	err := s.InsertAccount(a)
	require.NoError(t, err)

	assert.Len(t, s.GetAccounts(), 1)
	expectedAccount, err := s.GetAccount(a.ID())
	require.NoError(t, err)
	assert.ObjectsAreEqualValues(expectedAccount, a)
}

func TestGetSpecificAccountByAccountIdentifier(t *testing.T) {
	t.Parallel()

	s := storage.NewMemoryRepository()
	currency := *money.GetCurrency("BRL")
	accountIdentifier := account.Identifier(2)
	a := account.NewAccount(accountIdentifier, 100, currency)
	err := s.InsertAccount(a)
	require.NoError(t, err)

	expectedAccount, err := s.GetAccount(accountIdentifier)
	require.NoError(t, err)

	assert.ObjectsAreEqualValues(expectedAccount, a)
}

func TestGetAllAccountsInsertedInMemoryRepository(t *testing.T) {
	t.Parallel()

	s := storage.NewMemoryRepository()
	currency := *money.GetCurrency("BRL")
	a1 := account.NewAccount(1, 100, currency)
	err := s.InsertAccount(a1)
	require.NoError(t, err)
	a2 := account.NewAccount(2, 999, currency)
	err = s.InsertAccount(a2)
	require.NoError(t, err)

	accounts := s.GetAccounts()

	assert.Len(t, accounts, 2)
	assert.ObjectsAreEqualValues(a1, accounts[0])
	assert.ObjectsAreEqualValues(a2, accounts[1])
}

func TestCleanDataUsingMemoryRepository(t *testing.T) {
	t.Parallel()

	s := storage.NewMemoryRepository()
	currency := *money.GetCurrency("BRL")
	a1 := account.NewAccount(1, 100, currency)
	err := s.InsertAccount(a1)
	require.NoError(t, err)

	s.CleanData()

	assert.Len(t, s.GetAccounts(), 0)
}
