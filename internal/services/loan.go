package services

import (
	"amartha-loan/internal/interfaces"
	"amartha-loan/internal/models"
	"context"
	"time"
)

type LoanService struct {
	LoanRepo interfaces.ILoanRepo
}

func (s *LoanService) Create(ctx context.Context, loan *models.Loan) error {
	convertToYear := loan.WeeklyPayment / 52
	amountInterestRate := ((loan.Principal * loan.InterestRate) / 100) * convertToYear
	amountInterestRate += loan.Principal
	loan.RemainingAmount = amountInterestRate
	loan.TotalAmount = amountInterestRate
	loan.PaymentCompleteDate = estimate(loan.WeeklyPayment)
	return s.LoanRepo.CreateLoan(ctx, loan)
}

func estimate(WeeklyPayment float64) time.Time {
	durasi := time.Duration(WeeklyPayment*7*24) * time.Hour
	return time.Now().Add(durasi)
}

func (s *LoanService) GetLoanByID(ctx context.Context, ID int) (models.Loan, error) {
	obj, err := s.LoanRepo.GetLoanByID(ctx, ID)
	if err != nil {
		return obj, err
	}
	return obj, nil
}

func (s *LoanService) IsDelinquent(ctx context.Context, userID int) ([]models.Payoff, error) {
	resp, err := s.LoanRepo.IsDelinquent(ctx, userID)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *LoanService) MakePayment(ctx context.Context, payoffID int) error {
	return s.LoanRepo.MakePayment(ctx, payoffID)
}
