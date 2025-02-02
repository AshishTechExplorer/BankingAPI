package dto

import (
	"strings"
)

type NewAccountRequest struct {
	CustomerId  string  `db:"customer_id"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
}

func (r NewAccountRequest) Validate() error {
	if r.Amount < 5000 {
		//	return "To open a new account you need to deposit atleast 5000.00"
	}
	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		//	return errs.NewValidationError("Account type should be checking or saving")
	}
	return nil
}
