package doc_cr

import (
	"context"

	domainDocCr "itapp_ddd/domain/doc_cr"
)

// Service adalah application service untuk DocumentCr
type Service struct {
	repo domainDocCr.DocumentCrRepository
}

func NewService(repo domainDocCr.DocumentCrRepository) *Service {
	return &Service{repo: repo}
}

// GetAll mengambil semua DocumentCr berdasarkan filter
func (s *Service) GetAll(ctx context.Context, req domainDocCr.GetListRequest) ([]*domainDocCr.DocumentCrListItem, error) {
	return s.repo.GetAll(ctx, req)
}

func (s *Service) GetPr(ctx context.Context, req domainDocCr.GetListRequest) ([]*domainDocCr.DocumentCrListItem, error) {
	return s.repo.GetPr(ctx, req)
}
