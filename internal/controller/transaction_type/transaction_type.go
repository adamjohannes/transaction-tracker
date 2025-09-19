package transaction_type

import (
	"context"
	domain "monthly-expenses-handler/internal/domain/transaction_type"
	repository "monthly-expenses-handler/internal/repository/transaction_type"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TypeController struct {
	pool *pgxpool.Pool
	ctx  context.Context
}

func NewTypeController(pool *pgxpool.Pool, ctx context.Context) *TypeController {
	return &TypeController{pool: pool, ctx: ctx}
}

func (tc *TypeController) GetAllTransactionTypes() ([]*domain.TransactionType, error) {
	repo := repository.NewPostgresRepository(tc.pool)
	return repo.GetAll(tc.ctx)
}
