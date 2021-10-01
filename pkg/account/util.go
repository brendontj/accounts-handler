package account

import "errors"

const ZeroAmount = 0

var ErrAccountWithoutBalance = errors.New("account without balance")
