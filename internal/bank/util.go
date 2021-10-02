package bank

import "errors"

const (
	MinimumDepositAmount = 1
	MinimumWithdrawAmount = 1
)

var (
	ErrAmountLessThanMinimumDepositAmount  = errors.New("amount requested is less than minimum deposit amount")
	ErrAmountLessThanMinimumWithdrawAmount = errors.New("amount requested is less than minimum withdraw amount")
)