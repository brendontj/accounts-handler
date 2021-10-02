package account

import "errors"

const BrazilianCurrencyCode = "BRL"

var ErrAccountWithoutBalance = errors.New("account without balance")

type Identifier int