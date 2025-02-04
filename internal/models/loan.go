package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Loan struct {
	ID                  int       `json:"id"`
	UserID              int       `json:"user_id" gorm:"column:user_id;" validate:"required"`
	Principal           float64   `json:"principal" gorm:"column:principal;type:decimal(15, 2)" validate:"required"`
	InterestRate        float64   `json:"interest_rate" gorm:"column:interest_rate;type:decimal(15, 2)" validate:"required"`
	TotalAmount         float64   `json:"total_amount" gorm:"column:total_amount;type:decimal(15, 2)"`
	RemainingAmount     float64   `json:"remaining_amount" gorm:"column:remaining_amount;type:decimal(15, 2)"`
	WeeklyPayment       float64   `json:"weekly_payment" gorm:"column:weekly_payment;type:decimal(15, 2)" validate:"required"`
	PaymentCompleteDate time.Time `json:"payment_complete_date" gorm:"column:payment_complete_date;"`
	CreatedAt           time.Time `json:"-"`
	UpdatedAt           time.Time `json:"-"`
}

func (*Loan) TableName() string {
	return "loans"
}

func (l Loan) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type Payoff struct {
	ID        int       `json:"id"`
	LoanID    int       `json:"loan_id" gorm:"column:loan_id"`
	Amount    float64   `json:"amount" gorm:"column:amount;type:decimal(15, 2)"`
	PaidAt    time.Time `json:"paid_at" gorm:"column:paid_at;not null"`
	IsPaid    int       `json:"is_paid" gorm:"column:is_paid"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (*Payoff) TableName() string {
	return "payoffs"
}

type WalletHistoryParam struct {
	Page                  int    `form:"page"`
	Limit                 int    `form:"limit"`
	WalletTransactionType string `form:"wallet_transaction_type"`
}
