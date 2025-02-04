package interfaces

import (
	"amartha-loan/internal/models"
	"context"

	"github.com/gin-gonic/gin"
)

type ILoanRepo interface {
	CreateLoan(ctx context.Context, wallet *models.Loan) error
	GetLoanByID(ctx context.Context, ID int) (models.Loan, error)
	IsDelinquent(ctx context.Context, userID int) ([]models.Payoff, error)
	MakePayment(ctx context.Context, payoffID int) error
}

type ILoanService interface {
	Create(ctx context.Context, loan *models.Loan) error
	GetLoanByID(ctx context.Context, ID int) (models.Loan, error)
	IsDelinquent(ctx context.Context, userID int) ([]models.Payoff, error)
	MakePayment(ctx context.Context, payoffID int) error
}
type ILoanAPI interface {
	Create(*gin.Context)
	GetLoanByID(c *gin.Context)
	IsDelinquent(c *gin.Context)
	MakePayment(c *gin.Context)
}
