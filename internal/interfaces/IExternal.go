package interfaces

import (
	"amartha-loan/internal/models"
	"context"
)

type IExternal interface {
	ValidateToken(ctx context.Context, token string) (models.TokenData, error)
}
