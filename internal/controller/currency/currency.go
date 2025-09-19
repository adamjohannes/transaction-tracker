package currency

import (
	"context"
	domain "monthly-expenses-handler/internal/domain/currency"
	repository "monthly-expenses-handler/internal/repository/currency"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CurrencyController struct {
	pool *pgxpool.Pool
	ctx  context.Context
}

func NewCurrencyController(pool *pgxpool.Pool, ctx context.Context) *CurrencyController {
	return &CurrencyController{pool: pool, ctx: ctx}
}

func (cc *CurrencyController) GetAllCurrencies() ([]*domain.Currency, error) {
	repo := repository.NewPostgresRepository(cc.pool)
	return repo.GetAll(cc.ctx)
}
