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
	DocCr             *DocumentCr
	Review            *Review
	Approvals         []*Approval
	PICDevelopers     []*PIC
	PICFunctionals    []*PIC
}

// DocumentCrRepository adalah kontrak untuk akses data DocumentCr
type DocumentCrRepository interface {
	GetAll(ctx context.Context, req GetListRequest) ([]*DocumentCrListItem, error)
	GetDetail(ctx context.Context, req GetDetailRequest) (*DocumentCrDetailItem, error)
}

// GetDetailRequest adalah parameter untuk mengambil detail DocumentCr
type GetDetailRequest struct {
	DocCrID uint
}
 
// DocumentCrDetailItem adalah DTO untuk detail response
type DocumentCrDetailItem struct {
	DocCr          *DocumentCr
	Review         *Review
	Approvals      []*Approval
	PICDevelopers  []*PIC
	PICFunctionals []*PIC
	PICViewers     []*PIC
	PICQAs         []*PIC
	Milestones     []*Milestone
}