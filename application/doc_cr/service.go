package doc_cr

import (
	"context"

	domainDocCr "itapp_ddd/domain/doc_cr"
)

// Service adalah application service untuk DocumentCr
type Service struct {
	repo   domainDocCr.DocumentCrRepository
	logger Logger
}

func NewService(repo domainDocCr.DocumentCrRepository, logger Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

// GetAll mengambil semua DocumentCr berdasarkan filter
func (s *Service) GetAll(ctx context.Context, req domainDocCr.GetListRequest) ([]*domainDocCr.DocumentCrListItem, error) {
	items, err := s.repo.GetAll(ctx, req)
	if err != nil {
		s.logger.Error("GetAll failed", err)
		return nil, err
	}
	return items, nil
}

func (s *Service) GetDetail(ctx context.Context, req domainDocCr.GetDetailRequest) (*domainDocCr.DocumentCrDetailItem, error) {

	item, err := s.repo.GetDetail(ctx, req)
	if err != nil {
		s.logger.Error("GetDetail failed", err)
		return nil, err
	}

	return item, nil
}
