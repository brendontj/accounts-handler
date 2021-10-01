package account_test

import (
	"cautious-octo-pancake/pkg/account"
	"github.com/Rhymond/go-money"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Withdraw_From_An_Account_With_Enough_Amount(t *testing.T) {
	t.Parallel()

	testCurrency := *money.GetCurrency("BRL")

	type testCase struct {
		currency money.Currency
		account account.Account
		amountToWithdraw int64
		balanceExpected int64
	}

	testCases := []testCase{
		{
			currency:         testCurrency,
			account:          account.NewAccount(1, 100, testCurrency),
			amountToWithdraw: 50,
			balanceExpected:  int64(50),
		},
		{
			currency:         testCurrency,
			account:          account.NewAccount(1, 100, testCurrency),
			amountToWithdraw: 100,
			balanceExpected:  int64(0),
		},
		{
			currency:         testCurrency,
			account:          account.NewAccount(1, 100, testCurrency),
			amountToWithdraw: 0,
			balanceExpected:  int64(100),
		},
		{
			currency:         testCurrency,
			account:          account.NewAccount(1, 100, testCurrency),
			amountToWithdraw: 100,
			balanceExpected:  int64(0),
		},
	}

	for _, tc := range testCases {
		err := tc.account.Withdraw(tc.amountToWithdraw)
		require.NoError(t, err)

		assert.Equal(t, tc.balanceExpected, tc.account.Balance())
	}
}

func Test_Withdraw_From_An_Account_Without_Enough_Amount(t *testing.T) {
	t.Parallel()

	testCurrency := *money.GetCurrency("BRL")
	a := account.NewAccount(1, 100, testCurrency)

	err := a.Withdraw(200)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, account.ErrAccountWithoutBalance)
}

func Test_Deposit_With_Success(t *testing.T) {
	t.Parallel()

	testCurrency := *money.GetCurrency("BRL")

	type testCase struct {
		currency money.Currency
		account account.Account
		amountToDeposit int64
		balanceExpected int64
	}

	testCases := []testCase{
		{
			currency:         testCurrency,
			account:          account.NewAccount(1, 0, testCurrency),
			amountToDeposit: 50,
			balanceExpected:  int64(50),
		},
		{
			currency:         testCurrency,
			account:          account.NewAccount(1, 1, testCurrency),
			amountToDeposit: 133,
			balanceExpected:  int64(134),
		},
		{
			currency:         testCurrency,
			account:          account.NewAccount(1, 101, testCurrency),
			amountToDeposit: 99,
			balanceExpected:  int64(200),
		},
		{
			currency:         testCurrency,
			account:          account.NewAccount(1, 100, testCurrency),
			amountToDeposit: 0,
			balanceExpected:  int64(100),
		},
	}

	for _, tc := range testCases {
		err := tc.account.Deposit(tc.amountToDeposit)
		require.NoError(t, err)

		assert.Equal(t, tc.balanceExpected, tc.account.Balance())
	}
}

func Test_account_RollbackBalanceTo(t *testing.T) {
	t.Parallel()

	testCurrency := *money.GetCurrency("BRL")
	a := account.NewAccount(1, 200, testCurrency)

	a.RollbackBalanceTo(100)

	assert.Equal(t, int64(100), a.Balance())
}
