package doc_cr

import (
	"context"
	"time"
)

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
	GetDetail(ctx context.Context, req GetDetailRequest) (*DocumentCrDetailItem, error)
	Create(ctx context.Context, req CreateDocCrRequest) (*DocumentCr, error)
	CountByYear(ctx context.Context, year string) (int64, error)
	SendReview(ctx context.Context, req SendReviewRequest) (*SendReviewResult, error)
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
type CreateDocCrRequest struct {
	DocCrID                uint
	DocCrName              string
	FormId                 string
	DocReference           string
	RequesterNIK           string
	RequesterName          string
	RequesterDepartment    string
	ReasonDetail           string
	ProblemDetail          string
	ServerType             string
	CrType                 string
	CrPriority             string
	GitMergeRequest        string
	GitRepository          string
	DocNotes               string
	DocAttachment          []string
	DocAttachmentExtension []string
	DocStatus              string
	ApplicationName        string
	DeployedDate           *time.Time
}

// SendReviewRequest adalah request untuk me-review CR (DDD-simplified)
type SendReviewRequest struct {
	DocCrID             uint
	DocCrName           string
	FormId              string
	DocReference        string
	RequesterNIK        string
	RequesterName       string
	RequesterDepartment string
	PICDevelopers       []PIC
	PICFunctionals      []PIC
	PICViewers          []PIC
	PICQAs              []PIC
	ReasonDetail        string
	ProblemDetail       string
	ServerType          string
	CrType              string
	CrPriority          string
	GitMergeRequest     string
	GitRepository       string
	DocNotes            string
	DocAttachment       []string
	DocStatus           string
	ApplicationName     string
	ReviewNote          string
	NikLog              string
	NameLog             string
	NoteLog             string
	Approvals           []*Approval
}

// SendReviewResult adalah hasil dari operasi send review
type SendReviewResult struct {
	DocCrID uint
}
