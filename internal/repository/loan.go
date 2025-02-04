package repository

import (
	"amartha-loan/internal/models"
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type LoanRepo struct {
	DB *gorm.DB
}

func (r *LoanRepo) CreateLoan(ctx context.Context, loan *models.Loan) error {
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(loan).Error
		if err != nil {
			return err
		}
		if loan.ID == 0 {
			return errors.New("loan id 0")
		}
		startDate := time.Now()
		amount := loan.TotalAmount / float64(loan.WeeklyPayment)
		for i := 0; i < int(loan.WeeklyPayment); i++ {
			tanggal := startDate.Add(time.Duration(i*7*24) * time.Hour)
			obj := models.Payoff{
				LoanID: loan.ID,
				Amount: amount,
				PaidAt: tanggal,
			}
			err = r.CreatePayoff(ctx, &obj, tx)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

func (r *LoanRepo) CreatePayoff(ctx context.Context, payoff *models.Payoff, tx *gorm.DB) error {
	return r.DB.Create(payoff).Error
}

func (r *LoanRepo) GetLoanByID(ctx context.Context, ID int) (models.Loan, error) {
	var (
		resp models.Loan
	)
	err := r.DB.Where("id = ?", ID).Last(&resp).Error

	return resp, err
}

func (r *LoanRepo) IsDelinquent(ctx context.Context, userID int) ([]models.Payoff, error) {
	var (
		resp  = []models.Payoff{}
		loans = []models.Loan{}
	)
	if err := r.DB.Where("user_id = ?", userID).Find(&loans).Error; err != nil {
		return resp, err
	}
	if len(loans) == 0 {
		return resp, fmt.Errorf("no loans found for user with ID %d", userID)
	}
	if err := r.DB.Where("is_paid = 0 and paid_at < now() and loan_id in (?)", getLoanIDs(loans)).Find(&resp).Error; err != nil {
		return resp, err
	}
	return resp, nil
}

func getLoanIDs(loans []models.Loan) []int {
	var loanIDs []int
	for _, loan := range loans {
		loanIDs = append(loanIDs, loan.ID)
	}
	return loanIDs
}

func (r *LoanRepo) MakePayment(ctx context.Context, payoffID int) error {
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		var (
			obj  models.Payoff
			loan models.Loan
		)

		if err := tx.Where("id = ?", payoffID).First(&obj).Error; err != nil {
			return fmt.Errorf("failed to find payoff: %w", err)
		}

		obj.IsPaid = 1
		if err := tx.Save(&obj).Error; err != nil {
			return fmt.Errorf("failed to update payoff status: %w", err)
		}

		if err := tx.Where("id = ?", obj.LoanID).First(&loan).Error; err != nil {
			return fmt.Errorf("failed to find loan for payoff: %w", err)
		}

		if loan.RemainingAmount < obj.Amount {
			return fmt.Errorf("remaining amount is less than the payment amount: %v < %v", loan.RemainingAmount, obj.Amount)
		}

		loan.RemainingAmount -= obj.Amount
		if err := tx.Save(&loan).Error; err != nil {
			return fmt.Errorf("failed to update loan remaining amount: %w", err)
		}

		return nil
	})

	return err
}
