package bank_test

import (
	"cautious-octo-pancake/internal/bank"
	"cautious-octo-pancake/internal/bank/storage"
	"cautious-octo-pancake/pkg/account"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOpenAccountInABank(t *testing.T) {
	s:= storage.NewMemoryRepository()
	b:= bank.NewBank(s)
	accountIdentifier := account.Identifier(1)

	a, err := b.OpenAccount(accountIdentifier, 1000)
	require.NoError(t, err)

	givenAccount, err := b.GetAccount(accountIdentifier)
	require.NoError(t, err)
	assert.ObjectsAreEqualValues(a, givenAccount)
}

func TestBankTransferBetweenAccounts(t *testing.T) {
	s:= storage.NewMemoryRepository()
	b:= bank.NewBank(s)
	a1, err := b.OpenAccount(1, 1000)
	require.NoError(t, err)
	a2, err := b.OpenAccount(2, 2000)
	require.NoError(t, err)

	err = b.Transfer(a1, a2, 1000)
	require.NoError(t, err)

	assert.Equal(t, int64(0), a1.Balance())
	assert.Equal(t, int64(3000), a2.Balance())
}
