package application

import "context"
 
type SSOClient interface {
	FetchApplications(ctx context.Context, req GetListApplicationRequest) ([]SSOApplication, error)
}
