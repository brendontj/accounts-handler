package account_service_test

import (
	account2 "cautious-octo-pancake/internal/account"
	"cautious-octo-pancake/internal/database"
	"cautious-octo-pancake/pkg/account"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOpenAccount(t *testing.T) {
	t.Parallel()

	s:= database.NewMemoryRepository()
	accountHandler := account2.NewAccountHandler(s)
	accountIdentifier := account.Identifier(1)

	openedAccount, err := accountHandler.OpenAccount(accountIdentifier, 1000)
	require.NoError(t, err)

	givenAccount, err := accountHandler.GetAccount(accountIdentifier)
	require.NoError(t, err)
	assert.ObjectsAreEqualValues(openedAccount, givenAccount)
}

func TestOpenAccountWithTheSameIDFromAnExistingAccount(t *testing.T) {
	t.Parallel()

	memoryRepository := database.NewMemoryRepository()
	accountHandler := account2.NewAccountHandler(memoryRepository)
	accountIdentifier := account.Identifier(1)
	_, err := accountHandler.OpenAccount(accountIdentifier, 1000)
	require.NoError(t, err)

	a, err := accountHandler.OpenAccount(accountIdentifier, 10)

	require.Nil(t, a)
	require.NotNil(t, err)
	require.EqualError(t, err, fmt.Sprintf("unable to open account with id %v, cause exists account with" +
		" id informed", accountIdentifier))
}

func TestTransferBetweenAccounts(t *testing.T) {
	t.Parallel()

	memoryRepository := database.NewMemoryRepository()
	accountHandler := account2.NewAccountHandler(memoryRepository)
	firstAccount, err := accountHandler.OpenAccount(1, 1000)
	require.NoError(t, err)
	secondAccount, err := accountHandler.OpenAccount(2, 2000)
	require.NoError(t, err)

	err = accountHandler.Transfer(firstAccount, secondAccount, 1000)
	require.NoError(t, err)

	assert.Equal(t, int64(0), firstAccount.Balance())
	assert.Equal(t, int64(3000), secondAccount.Balance())
}

func TestTransferNegativeAmountBetweenAccounts(t *testing.T) {
	t.Parallel()

	memoryRepository := database.NewMemoryRepository()
	accountHandler := account2.NewAccountHandler(memoryRepository)
	firstAccount, err := accountHandler.OpenAccount(1, 1000)
	require.NoError(t, err)
	secondAccount, err := accountHandler.OpenAccount(2, 2000)
	require.NoError(t, err)

	err = accountHandler.Transfer(firstAccount, secondAccount, -1)
	require.NotNil(t, err)
	require.ErrorIs(t, err, account2.ErrAmountLessThanMinimumWithdrawAmount)
}

func TestWithdrawFromAnExistingAccount(t *testing.T) {
	t.Parallel()

	memoryRepository := database.NewMemoryRepository()
	accountHandler := account2.NewAccountHandler(memoryRepository)
	existingAccount, err := accountHandler.OpenAccount(1, 1000)
	require.NoError(t, err)

	err = accountHandler.AccountWithdraw(existingAccount, 200)
	require.NoError(t, err)

	assert.Equal(t, int64(800), existingAccount.Balance())
}

func TestWithdrawLessThanMinimumAllowedAmount(t *testing.T) {
	t.Parallel()

	memoryRepository := database.NewMemoryRepository()
	accountHandler := account2.NewAccountHandler(memoryRepository)
	acc, err := accountHandler.OpenAccount(1, 1000)
	require.NoError(t, err)

	err = accountHandler.AccountWithdraw(acc, account2.MinimumWithdrawAmount- 1)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, account2.ErrAmountLessThanMinimumWithdrawAmount)
}

func TestDepositFromAnExistingAccount(t *testing.T) {
	t.Parallel()

	memoryRepository := database.NewMemoryRepository()
	accountHandler := account2.NewAccountHandler(memoryRepository)
	existingAccount, err := accountHandler.OpenAccount(1, 1000)
	require.NoError(t, err)

	err = accountHandler.AccountDeposit(existingAccount, 200)
	require.NoError(t, err)

	assert.Equal(t, int64(1200), existingAccount.Balance())
}

func TestDepositLessThanMinimumAllowedAmount(t *testing.T) {
	t.Parallel()

	memoryRepository := database.NewMemoryRepository()
	accountHandler := account2.NewAccountHandler(memoryRepository)
	acc, err := accountHandler.OpenAccount(1, 1000)
	require.NoError(t, err)

	err = accountHandler.AccountDeposit(acc, account2.MinimumDepositAmount- 1)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, account2.ErrAmountLessThanMinimumDepositAmount)
}