package application
import (
	"context"
	"errors"
	"strings"
 
	domainApp "itapp_ddd/domain/application"
)
 
type Logger interface {
	Info(msg string, args ...interface{})
	Error(msg string, err error, args ...interface{})
	Warn(msg string, args ...interface{})
}
 
// Service berisi use case Application
type Service struct {
	sso    domainApp.SSOClient
	logger Logger
}
 
// NewService constructor
func NewService(sso domainApp.SSOClient, logger Logger) *Service {
	return &Service{sso: sso, logger: logger}
}
 

func (s *Service) GetAll(ctx context.Context, req domainApp.GetListApplicationRequest) ([]*domainApp.Application, error) {
	if strings.TrimSpace(req.SessionToken) == "" {
		return nil, errors.New("session_token are required")
	}
 
	items, err := s.sso.FetchApplications(ctx, req)
	if err != nil {
		s.logger.Error("GetAll from SSO failed", err)
		return nil, err
	}
 
	apps := make([]*domainApp.Application, 0, len(items))
	for _, item := range items {
		apps = append(apps, &domainApp.Application{
			ApplicationID:   item.ID,
			ApplicationName: item.Module,
		})
	}
	return apps, nil
}
