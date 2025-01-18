package infraestructure_services

import (
	"context"
	infraestructure_repository "my_wallet/api/respository/healtcheck"

	"github.com/sirupsen/logrus"
)

type HealtcheckService interface {
	GetHealtcheck(ctx context.Context) (bool, error)
}

type healtcheckService struct {
	ctx        context.Context
	repository infraestructure_repository.HealtCheckRepository
	logger     logrus.FieldLogger
}

func NewHealtcheckService(ctx context.Context, repo infraestructure_repository.HealtCheckRepository, logger logrus.FieldLogger) *healtcheckService {
	return &healtcheckService{
		ctx:        ctx,
		repository: repo,
		logger:     logger,
	}
}

func (s *healtcheckService) GetHealtcheck(ctx context.Context) (bool, error) {
	return s.repository.GetHealtcheck(ctx)
}
