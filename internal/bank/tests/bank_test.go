package bank_test

import (
	"cautious-octo-pancake/internal/bank"
	"cautious-octo-pancake/internal/bank/storage"
	"cautious-octo-pancake/pkg/account"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOpenAccountInABank(t *testing.T) {
	t.Parallel()

	s:= storage.NewMemoryRepository()
	b:= bank.NewBank(s)
	accountIdentifier := account.Identifier(1)

	a, err := b.OpenAccount(accountIdentifier, 1000)
	require.NoError(t, err)

	givenAccount, err := b.GetAccount(accountIdentifier)
	require.NoError(t, err)
	assert.ObjectsAreEqualValues(a, givenAccount)
}

func TestOpenAccountWithTheSameIDFromAnExistingAccount(t *testing.T) {
	t.Parallel()

	s:= storage.NewMemoryRepository()
	b:= bank.NewBank(s)
	accountIdentifier := account.Identifier(1)
	_, err := b.OpenAccount(accountIdentifier, 1000)
	require.NoError(t, err)

	a, err := b.OpenAccount(accountIdentifier, 10)

	require.Nil(t, a)
	require.NotNil(t, err)
	require.EqualError(t, err, fmt.Sprintf("unable to open account with id %v, cause exists account with" +
		" id informed", accountIdentifier))
}

func TestBankTransferBetweenAccounts(t *testing.T) {
	t.Parallel()

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

func TestTransferNegativeAmountBetweenAccounts(t *testing.T) {
	t.Parallel()

	s:= storage.NewMemoryRepository()
	b:= bank.NewBank(s)
	a1, err := b.OpenAccount(1, 1000)
	require.NoError(t, err)
	a2, err := b.OpenAccount(2, 2000)
	require.NoError(t, err)

	err = b.Transfer(a1, a2, -1)
	require.NotNil(t, err)
	require.ErrorIs(t, err, bank.ErrAmountLessThanMinimumWithdrawAmount)
}

func TestWithdrawFromAnExistingAccount(t *testing.T) {
	t.Parallel()

	s:= storage.NewMemoryRepository()
	b:= bank.NewBank(s)
	a, err := b.OpenAccount(1, 1000)
	require.NoError(t, err)

	err = b.AccountWithdraw(a, 200)
	require.NoError(t, err)

	assert.Equal(t, int64(800), a.Balance())
}

func TestWithdrawLessThanMinimumAllowedAmount(t *testing.T) {
	t.Parallel()

	s:= storage.NewMemoryRepository()
	b:= bank.NewBank(s)
	a, err := b.OpenAccount(1, 1000)
	require.NoError(t, err)

	err = b.AccountWithdraw(a, bank.MinimumWithdrawAmount - 1)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, bank.ErrAmountLessThanMinimumWithdrawAmount)
}

func TestDepositFromAnExistingAccount(t *testing.T) {
	t.Parallel()

	s:= storage.NewMemoryRepository()
	b:= bank.NewBank(s)
	a, err := b.OpenAccount(1, 1000)
	require.NoError(t, err)

	err = b.AccountDeposit(a, 200)
	require.NoError(t, err)

	assert.Equal(t, int64(1200), a.Balance())
}

func TestDepositLessThanMinimumAllowedAmount(t *testing.T) {
	t.Parallel()

	s:= storage.NewMemoryRepository()
	b:= bank.NewBank(s)
	a, err := b.OpenAccount(1, 1000)
	require.NoError(t, err)

	err = b.AccountDeposit(a, bank.MinimumDepositAmount - 1)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, bank.ErrAmountLessThanMinimumDepositAmount)
}