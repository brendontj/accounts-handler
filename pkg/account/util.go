package account

import (
	"errors"
	"strconv"
)

const BrazilianCurrencyCode = "BRL"

var ErrAccountWithoutBalance = errors.New("account without balance")

type Identifier int

func (i Identifier) String() string{
	return strconv.Itoa(int(i))
}