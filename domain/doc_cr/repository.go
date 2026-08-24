package doc_cr

import "context"

// GetListRequest adalah parameter untuk filter data
type GetListRequest struct {
	StartDate    string
	EndDate      string
	RequesterNik string
	ApprovalNik  string
	DocStatus    []string
}

// DocumentCrListItem adalah DTO untuk list response
type DocumentCrListItem struct {
	DocCr          *DocumentCr
	Review         *Review
	Approvals      []*Approval
	PICDevelopers  []*PIC
	PICFunctionals []*PIC
}

// DocumentCrRepository adalah kontrak untuk akses data DocumentCr
type DocumentCrRepository interface {
	GetAll(ctx context.Context, req GetListRequest) ([]*DocumentCrListItem, error)
	GetAndika(ctx context.Context, req GetListRequest) ([]*DocumentCrListItem, error)
	GetOneMine(ctx context.Context, req GetListRequest) ([]*DocumentCrListItem, error)
}
