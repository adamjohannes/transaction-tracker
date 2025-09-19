package status

import (
	"context"
	domain "monthly-expenses-handler/internal/domain/status"
	repository "monthly-expenses-handler/internal/repository/status"

	"github.com/jackc/pgx/v5/pgxpool"
)

type StatusController struct {
	pool *pgxpool.Pool
	ctx  context.Context
}

func NewStatusController(pool *pgxpool.Pool, ctx context.Context) *StatusController {
	return &StatusController{pool: pool, ctx: ctx}
}

func (sc *StatusController) GetAllStatus() ([]*domain.Status, error) {
	repo := repository.NewPostgresRepository(sc.pool)
	return repo.GetAll(sc.ctx)
}
