package bank

import "errors"

const (
	MinimumDepositAmount = 0
	MinimumWithdrawAmount = 0
)

var (
	ErrAmountLessThanMinimumDepositAmount  = errors.New("amount requested is less than minimum deposit amount")
	ErrAmountLessThanMinimumWithdrawAmount = errors.New("amount requested is less than minimum withdraw amount")
)