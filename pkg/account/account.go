package account

import (
	"errors"
	"fmt"
	"github.com/Rhymond/go-money"
)

var ErrAccountWithoutBalance = errors.New("account without balance")

type Account interface {
	ID() int64
	Balance() int64
	Deposit(amount int64) error
	Withdraw(amount int64) error
	RollbackBalanceTo(amount int64)
}

type account struct {
	id       int64
	balance  *money.Money
	currency money.Currency
}

func NewAccount(destination, amount int64, currency money.Currency) Account {
	return &account{
		id:       destination,
		balance:  money.New(amount, currency.Code),
		currency: currency,
	}
}

func (a *account) ID() int64 {
	return a.id
}

func (a *account) Balance() int64 {
	return a.balance.Amount()
}

func (a *account) Deposit(amount int64) error {
	newMoney, err := a.balance.Add(money.New(amount, a.currency.Code))
	if err != nil {
		return fmt.Errorf("unable to deposit amount in account with id = %v", a.id)
	}

	a.balance = newMoney
	return nil
}

func (a *account) Withdraw(amount int64) error {
	gte, err := a.balance.GreaterThanOrEqual(money.New(amount, a.currency.Code))
	if err != nil {
		return fmt.Errorf("unable to check if the current balance is greater than or equal than requested "+
			"amount: account intentifier = %v", a.ID())
	}

	if !gte {
		return ErrAccountWithoutBalance
	}

	newMoney, err := a.balance.Subtract(money.New(amount, a.currency.Code))
	if err != nil {
		return fmt.Errorf("unable to withdraw money from the account with id = %v", a.id)
	}

	a.balance = newMoney
	return nil
}

func (a *account) RollbackBalanceTo(amount int64) {
	a.balance = money.New(amount, a.currency.Code)
}
