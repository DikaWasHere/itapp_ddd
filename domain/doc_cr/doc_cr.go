package doc_cr

import (
	"errors"
	"time"
)

// Value Object: DocCrID
type DocCrID uint

// Value Object: FormID
type FormID string

func NewFormID(id string) (FormID, error) {
	if id == "" {
		return "", errors.New("form ID cannot be empty")
	}
	return FormID(id), nil
}

// Value Object: Review
type Review struct {
	ReviewID     *uint
	ReviewNik    *string
	ReviewName   *string
	ReviewDate   *string
	ReviewStatus *string
	ReviewNote   *string
	DocCrID      *uint
}

// Value Object: Approval
type Approval struct {
	ApprovalID         *uint
	ApprovalNik        *string
	ApprovalName       *string
	ApprovalDate       *string
	ApprovalStatus     *string
	ApprovalDepartment *string
	DocCrID            *uint
}

// Value Object: PIC
type PIC struct {
	PICNik        string
	PICName       string
	PICDepartment string
}

type Milestone struct {
	LogID       uint
	DocCrID     uint
	LogActivity string
	NikLog      string
	NameLog     string
	NoteLog     string
	CreatedAt   time.Time
}

// Entity: DocumentCr (Aggregate Root)
type DocumentCr struct {
	DocCrID             uint
	DocCrName           string
	FormId              string
	DocReference        string
	RequesterNIK        string
	RequesterName       string
	RequesterDepartment string
	ReasonDetail        string
	ProblemDetail       string
	ServerType          string
	CrType              string
	CrPriority          string
	GitMergeRequest     string
	GitRepository       string
	DocNotes            string
	DocAttachment       string
	DocStatus           string
	ApplicationName     string
	DeployedDate        *time.Time
	CreatedAt           time.Time
}

// NewDocumentCr constructor
func NewDocumentCr(
	docCrID uint,
	docCrName string,
	formId string,
	docReference string,
	requesterNIK string,
	requesterName string,
	requesterDepartment string,
	reasonDetail string,
	problemDetail string,
	serverType string,
	crType string,
	crPriority string,
	gitMergeRequest string,
	gitRepository string,
	docNotes string,
	docAttachment string,
	docStatus string,
	applicationName string,
	deployedDate *time.Time,
	createdAt time.Time,
) *DocumentCr {
	return &DocumentCr{
		DocCrID:             docCrID,
		DocCrName:           docCrName,
		FormId:              formId,
		DocReference:        docReference,
		RequesterNIK:        requesterNIK,
		RequesterName:       requesterName,
		RequesterDepartment: requesterDepartment,
		ReasonDetail:        reasonDetail,
		ProblemDetail:       problemDetail,
		ServerType:          serverType,
		CrType:              crType,
		CrPriority:          crPriority,
		GitMergeRequest:     gitMergeRequest,
		GitRepository:       gitRepository,
		DocNotes:            docNotes,
		DocAttachment:       docAttachment,
		DocStatus:           docStatus,
		ApplicationName:     applicationName,
		DeployedDate:        deployedDate,
		CreatedAt:           createdAt,
	}
}
