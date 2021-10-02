package account

import (
	"fmt"
	"github.com/Rhymond/go-money"
)

type Account struct {
	id       Identifier
	balance  *money.Money
	currency money.Currency
}

func NewAccount(destination Identifier, amount int64, currency money.Currency) *Account {
	return &Account{
		id:       destination,
		balance:  money.New(amount, currency.Code),
		currency: currency,
	}
}

func (a *Account) ID() Identifier {
	return a.id
}

func (a *Account) Balance() int64 {
	return a.balance.Amount()
}

func (a *Account) Currency() money.Currency {
	return a.currency
}

func (a *Account) Deposit(amount int64) error {
	newMoney, err := a.balance.Add(money.New(amount, a.currency.Code))
	if err != nil {
		return fmt.Errorf("unable to deposit amount in Account with id = %v", a.id)
	}

	a.balance = newMoney
	return nil
}

func (a *Account) Withdraw(amount int64) error {
	gte, err := a.balance.GreaterThanOrEqual(money.New(amount, a.currency.Code))
	if err != nil {
		return fmt.Errorf("unable to check if the current balance is greater than or equal than requested "+
			"amount: Account intentifier = %v", a.ID())
	}

	if !gte {
		return ErrAccountWithoutBalance
	}

	newMoney, err := a.balance.Subtract(money.New(amount, a.currency.Code))
	if err != nil {
		return fmt.Errorf("unable to withdraw money from the Account with id = %v", a.id)
	}

	a.balance = newMoney
	return nil
}

func (a *Account) RollbackBalanceTo(amount int64) {
	a.balance = money.New(amount, a.currency.Code)
}
