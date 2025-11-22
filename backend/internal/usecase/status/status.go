package status

import (
	"context"
	"monthly-expenses-handler/internal/domain/status"
	"monthly-expenses-handler/internal/infrastructure/logger"
	statusRepo "monthly-expenses-handler/internal/repository/status"
)

type UseCase struct {
	ctx        context.Context
	statusRepo statusRepo.Repository
	logger     *logger.Logger
}

func NewStatusService(ctx context.Context, statusRepo statusRepo.Repository, logger *logger.Logger) UseCase {
	return UseCase{
		ctx,
		statusRepo,
		logger,
	}
}

func (u *UseCase) List() ([]*status.Status, error) {
	statusList, err := u.statusRepo.GetAll(u.ctx)

	if err != nil {
		return nil, err
	}

	return statusList, nil
}
