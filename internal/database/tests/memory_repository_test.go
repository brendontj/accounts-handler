package memory_storage_test

import (
	"cautious-octo-pancake/internal/database"
	"cautious-octo-pancake/pkg/account"
	"fmt"
	"github.com/Rhymond/go-money"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInsertAccountInMemoryRepository(t *testing.T) {
	t.Parallel()

	s := database.NewMemoryRepository()
	currency := *money.GetCurrency(database.BrazilianCurrencyCode)
	a := account.NewAccount(1, 100, currency)

	err := s.InsertAccount(a)
	require.NoError(t, err)

	assert.Len(t, s.GetAccounts(), 1)
	expectedAccount, err := s.GetAccount(a.ID())
	require.NoError(t, err)
	assert.ObjectsAreEqualValues(expectedAccount, a)
}

func TestInsertAccountDuplicatedInMemoryRepository(t *testing.T) {
	t.Parallel()

	s := database.NewMemoryRepository()
	currency := *money.GetCurrency(database.BrazilianCurrencyCode)
	a1 := account.NewAccount(1, 100, currency)
	a2 := account.NewAccount(1, 250, currency)
	err := s.InsertAccount(a1)
	require.NoError(t, err)

	err = s.InsertAccount(a2)

	assert.NotNil(t, err)
	assert.EqualError(t, err, fmt.Sprintf("unable to store new account with id %v, cause: exists account " +
		"with the same id registered", a2.ID()))
}

func TestGetSpecificAccountByAccountIdentifier(t *testing.T) {
	t.Parallel()

	s := database.NewMemoryRepository()
	currency := *money.GetCurrency(database.BrazilianCurrencyCode)
	accountIdentifier := account.Identifier(2)
	a := account.NewAccount(accountIdentifier, 100, currency)
	err := s.InsertAccount(a)
	require.NoError(t, err)

	expectedAccount, err := s.GetAccount(accountIdentifier)
	require.NoError(t, err)

	assert.ObjectsAreEqualValues(expectedAccount, a)
}

func TestGetAccountAndGetErrorAccountNotFound(t *testing.T) {
	t.Parallel()

	s := database.NewMemoryRepository()

	a, err := s.GetAccount(account.Identifier(1))

	require.Nil(t, a)
	require.NotNil(t, err)
	require.EqualError(t, err, database.ErrAccountNotFound.Error())
}

func TestGetAllAccountsInsertedInMemoryRepository(t *testing.T) {
	t.Parallel()

	s := database.NewMemoryRepository()
	currency := *money.GetCurrency(database.BrazilianCurrencyCode)
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

	s := database.NewMemoryRepository()
	currency := *money.GetCurrency(database.BrazilianCurrencyCode)
	a1 := account.NewAccount(1, 100, currency)
	err := s.InsertAccount(a1)
	require.NoError(t, err)

	s.CleanData()

	assert.Len(t, s.GetAccounts(), 0)
}
