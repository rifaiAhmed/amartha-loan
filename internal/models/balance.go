package models

import "github.com/go-playground/validator/v10"

type TransactionRequest struct {
	Reference string  `json:"reference" valid:"required"`
	Amount    float64 `json:"amount" valid:"required"`
}

func (l TransactionRequest) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type BalanceResponse struct {
	Balance float64 `json:"balance"`
}
