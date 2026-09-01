package persistence

import (
	"context"
	"fmt"
	"strings"
	"time"

	domainDocCr "itapp_ddd/domain/doc_cr"

	"gorm.io/gorm"
)

// DocumentCrRepositoryDB implementasi konkret dari DocumentCrRepository
type DocumentCrRepositoryDB struct {
	db *gorm.DB
}

func NewDocumentCrRepositoryDB(db *gorm.DB) *DocumentCrRepositoryDB {
	return &DocumentCrRepositoryDB{db: db}
}

// --- GORM Models ---

// DocumentCrModel adalah model GORM untuk tabel trs_document_cr
type DocumentCrModel struct {
	DocCrID             uint           `gorm:"primarykey;column:doc_cr_id"`
	DocCrName           string         `gorm:"column:doc_cr_name"`
	FormId              string         `gorm:"column:form_id"`
	DocReference        string         `gorm:"column:doc_reference"`
	RequesterNIK        string         `gorm:"column:requester_nik"`
	RequesterName       string         `gorm:"column:requester_name"`
	RequesterDepartment string         `gorm:"column:requester_department"`
	ReasonDetail        string         `gorm:"column:reason_detail"`
	ProblemDetail       string         `gorm:"column:problem_detail"`
	ServerType          string         `gorm:"column:server_type"`
	CrType              string         `gorm:"column:cr_type"`
	CrPriority          string         `gorm:"column:cr_priority"`
	GitMergeRequest     string         `gorm:"column:git_merge_request"`
	GitRepository       string         `gorm:"column:git_repository"`
	DocNotes            string         `gorm:"column:doc_notes"`
	DocAttachment       string         `gorm:"column:doc_attachment"`
	DocStatus           string         `gorm:"column:doc_status"`
	ApplicationName     string         `gorm:"column:application_name"`
	DeployedDate        *time.Time     `gorm:"column:deployed_date"`
	CreatedAt           time.Time      `gorm:"column:created_at"`
	UpdatedAt           time.Time      `gorm:"column:updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (DocumentCrModel) TableName() string {
	return "trs.trs_document_cr"
}

func (m *DocumentCrModel) toEntity() *domainDocCr.DocumentCr {
	return &domainDocCr.DocumentCr{
		DocCrID:             m.DocCrID,
		DocCrName:           m.DocCrName,
		FormId:              m.FormId,
		DocReference:        m.DocReference,
		RequesterNIK:        m.RequesterNIK,
		RequesterName:       m.RequesterName,
		RequesterDepartment: m.RequesterDepartment,
		ReasonDetail:        m.ReasonDetail,
		ProblemDetail:       m.ProblemDetail,
		ServerType:          m.ServerType,
		CrType:              m.CrType,
		CrPriority:          m.CrPriority,
		GitMergeRequest:     m.GitMergeRequest,
		GitRepository:       m.GitRepository,
		DocNotes:            m.DocNotes,
		DocAttachment:       m.DocAttachment,
		DocStatus:           m.DocStatus,
		ApplicationName:     m.ApplicationName,
		DeployedDate:        m.DeployedDate,
		CreatedAt:           m.CreatedAt,
	}
}

// ReviewModel adalah model GORM untuk tabel trs_review
type ReviewModel struct {
	ReviewID     uint      `gorm:"primarykey;column:review_id"`
	ReviewNik    string    `gorm:"column:review_nik"`
	ReviewName   string    `gorm:"column:review_name"`
	ReviewDate   *string   `gorm:"column:review_date"`
	ReviewStatus string    `gorm:"column:review_status"`
	ReviewNote   string    `gorm:"column:review_note"`
	DocCrID      uint      `gorm:"column:doc_cr_id"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

func (ReviewModel) TableName() string {
	return "trs.trs_review"
}

func (m *ReviewModel) toEntity() *domainDocCr.Review {
	return &domainDocCr.Review{
		ReviewID:     &m.ReviewID,
		ReviewNik:    &m.ReviewNik,
		ReviewName:   &m.ReviewName,
		ReviewDate:   m.ReviewDate,
		ReviewStatus: &m.ReviewStatus,
		ReviewNote:   &m.ReviewNote,
		DocCrID:      &m.DocCrID,
	}
}

// ApprovalModel adalah model GORM untuk tabel trs_approval
type ApprovalModel struct {
	ApprovalID         uint      `gorm:"primarykey;column:approval_id"`
	ApprovalNik        string    `gorm:"column:approval_nik"`
	ApprovalName       string    `gorm:"column:approval_name"`
	ApprovalDate       *string   `gorm:"column:approval_date"`
	ApprovalStatus     string    `gorm:"column:approval_status"`
	ApprovalNote       string    `gorm:"column:approval_note"`
	DocCrID            uint      `gorm:"column:doc_cr_id"`
	ApprovalDepartment string    `gorm:"column:approval_department"`
	CreatedAt          time.Time `gorm:"column:created_at"`
	UpdatedAt          time.Time `gorm:"column:updated_at"`
}

func (ApprovalModel) TableName() string {
	return "trs.trs_approval"
}

func (m *ApprovalModel) toEntity() *domainDocCr.Approval {
	return &domainDocCr.Approval{
		ApprovalID:         &m.ApprovalID,
		ApprovalNik:        &m.ApprovalNik,
		ApprovalName:       &m.ApprovalName,
		ApprovalDate:       m.ApprovalDate,
		ApprovalStatus:     &m.ApprovalStatus,
		ApprovalDepartment: &m.ApprovalDepartment,
		DocCrID:            &m.DocCrID,
	}
}

// PICModel adalah model GORM untuk tabel trs_pic
type PICModel struct {
	PICID         uint      `gorm:"primarykey;column:pic_id"`
	PICNik        string    `gorm:"column:pic_nik"`
	PICName       string    `gorm:"column:pic_name"`
	DocCrID       uint      `gorm:"column:doc_cr_id"`
	PICDepartment string    `gorm:"column:approval_department"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
}

func (PICModel) TableName() string {
	return "trs.trs_pic"
}

func (m *PICModel) toEntity() *domainDocCr.PIC {
	return &domainDocCr.PIC{
		PICNik:        m.PICNik,
		PICName:       m.PICName,
		PICDepartment: m.PICDepartment,
	}
}

// DocumentCrLogModel adalah model GORM untuk tabel trs_document_cr_log
type DocumentCrLogModel struct {
	LogID       uint           `gorm:"primarykey;column:log_id"`
	DocCrID     uint           `gorm:"column:doc_cr_id"`
	LogActivity string         `gorm:"column:log_activity"`
	NikLog      string         `gorm:"column:nik_log"`
	NameLog     string         `gorm:"column:name_log"`
	NoteLog     string         `gorm:"column:note_log"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (DocumentCrLogModel) TableName() string {
	return "trs.trs_document_cr_log"
}

func (m *DocumentCrLogModel) toEntity() *domainDocCr.Milestone {
	return &domainDocCr.Milestone{
		LogID:       m.LogID,
		DocCrID:     m.DocCrID,
		LogActivity: m.LogActivity,
		NikLog:      m.NikLog,
		NameLog:     m.NameLog,
		NoteLog:     m.NoteLog,
		CreatedAt:   m.CreatedAt,
	}
}

// GetAll mengambil semua DocumentCr berdasarkan filter
func (r *DocumentCrRepositoryDB) GetAll(ctx context.Context, req domainDocCr.GetListRequest) ([]*domainDocCr.DocumentCrListItem, error) {
	query := `
		SELECT * FROM (
			SELECT DISTINCT ON (tdc.doc_cr_id) tdc.*
			FROM trs.trs_document_cr tdc
			LEFT JOIN trs.trs_pic tp ON tdc.doc_cr_id = tp.doc_cr_id
			LEFT JOIN trs.trs_approval ta ON tdc.doc_cr_id = ta.doc_cr_id
			WHERE 1=1
	`
	var queryParams []interface{}

	if req.ApprovalNik != "" {
		query += " AND ta.approval_nik = ?"
		queryParams = append(queryParams, req.ApprovalNik)
	}

	if req.RequesterNik != "" {
		query += " AND (tp.pic_nik = ? OR tdc.requester_nik = ?)"
		queryParams = append(queryParams, req.RequesterNik, req.RequesterNik)
	}

	if req.StartDate != "" && req.EndDate != "" {
		query += " AND date_trunc('day', tdc.created_at) BETWEEN ? AND ?"
		queryParams = append(queryParams, req.StartDate, req.EndDate)
	}

	if len(req.DocStatus) > 0 {
		placeholders := make([]string, len(req.DocStatus))
		for i := range req.DocStatus {
			placeholders[i] = "?"
		}
		query += " AND tdc.doc_status IN (" + strings.Join(placeholders, ",") + ")"
		for _, status := range req.DocStatus {
			queryParams = append(queryParams, status)
		}
	}

	query += " ORDER BY tdc.doc_cr_id, tdc.created_at DESC ) AS subquery ORDER BY created_at DESC"

	rows, err := r.db.WithContext(ctx).Raw(query, queryParams...).Rows()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve DocumentCr data: %w", err)
	}
	defer rows.Close()

	var items []*domainDocCr.DocumentCrListItem

	for rows.Next() {
		var docModel DocumentCrModel
		if err := r.db.ScanRows(rows, &docModel); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		doc := docModel.toEntity()
		item := &domainDocCr.DocumentCrListItem{
			DocCr:          doc,
			Approvals:      make([]*domainDocCr.Approval, 0),
			PICDevelopers:  make([]*domainDocCr.PIC, 0),
			PICFunctionals: make([]*domainDocCr.PIC, 0),
		}

		// Fetch related data (errors logged but don't fail the request)
		r.fetchRelatedData(ctx, doc.DocCrID, item)

		items = append(items, item)
	}

	return items, nil
}

// GetDetail mengambil detail DocumentCr berdasarkan DocCrID
func (r *DocumentCrRepositoryDB) GetDetail(ctx context.Context, req domainDocCr.GetDetailRequest) (*domainDocCr.DocumentCrDetailItem, error) {
	var docModel DocumentCrModel
	err := r.db.WithContext(ctx).Raw(`
		SELECT tdc.*
		FROM trs.trs_document_cr tdc
		WHERE tdc.doc_cr_id = ?;
	`, req.DocCrID).Scan(&docModel).Error

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve DocumentCr data: %w", err)
	}

	doc := docModel.toEntity()

	item := &domainDocCr.DocumentCrDetailItem{
		DocCr:          doc,
		Approvals:      make([]*domainDocCr.Approval, 0),
		PICDevelopers:  make([]*domainDocCr.PIC, 0),
		PICFunctionals: make([]*domainDocCr.PIC, 0),
		PICViewers:     make([]*domainDocCr.PIC, 0),
		PICQAs:         make([]*domainDocCr.PIC, 0),
		Milestones:     make([]*domainDocCr.Milestone, 0),
	}

	// Fetch Approval
	if err := r.fetchApprovals(ctx, doc.DocCrID, item); err != nil {
		return nil, fmt.Errorf("failed to retrieve approval data: %w", err)
	}

	// Fetch PIC
	if err := r.fetchPICs(ctx, doc.DocCrID, item); err != nil {
		return nil, fmt.Errorf("failed to retrieve PIC data: %w", err)
	}

	// Fetch Review (optional - don't fail if missing)
	r.fetchReview(ctx, doc.DocCrID, item)

	// Fetch Milestones (optional)
	r.fetchMilestones(ctx, doc.DocCrID, item)

	return item, nil
}

// --- Private helper methods ---

// fetchApprovals mengambil data approval untuk DocumentCr tertentu
func (r *DocumentCrRepositoryDB) fetchApprovals(ctx context.Context, docCrID uint, item *domainDocCr.DocumentCrDetailItem) error {
	rows, err := r.db.WithContext(ctx).Raw(`
		SELECT *
		FROM trs.trs_approval
		WHERE doc_cr_id = ? AND approval_status != 'Cancel';
	`, docCrID).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var approvalModel ApprovalModel
		r.db.ScanRows(rows, &approvalModel)
		item.Approvals = append(item.Approvals, approvalModel.toEntity())
	}
	return nil
}

// fetchPICs mengambil data PIC dan mengkategorikan berdasarkan department
func (r *DocumentCrRepositoryDB) fetchPICs(ctx context.Context, docCrID uint, item *domainDocCr.DocumentCrDetailItem) error {
	rows, err := r.db.WithContext(ctx).Raw(`
		SELECT *
		FROM trs.trs_pic
		WHERE doc_cr_id = ?;
	`, docCrID).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var picModel PICModel
		r.db.ScanRows(rows, &picModel)
		pic := picModel.toEntity()

		switch pic.PICDepartment {
		case "IT FUNCTIONAL":
			item.PICFunctionals = append(item.PICFunctionals, pic)
		case "IT DEVELOPMENT":
			item.PICDevelopers = append(item.PICDevelopers, pic)
		case "IT SERVICES & IMPLEMENTATION":
			item.PICViewers = append(item.PICViewers, pic)
		case "QA":
			item.PICQAs = append(item.PICQAs, pic)
		}
	}
	return nil
}

// fetchReview mengambil data review (optional)
func (r *DocumentCrRepositoryDB) fetchReview(ctx context.Context, docCrID uint, item *domainDocCr.DocumentCrDetailItem) {
	var reviewModel ReviewModel
	err := r.db.WithContext(ctx).Raw(`
		SELECT *
		FROM trs.trs_review
		WHERE doc_cr_id = ?;
	`, docCrID).Scan(&reviewModel).Error
	if err == nil {
		item.Review = reviewModel.toEntity()
	}
}

// fetchMilestones mengambil data milestones (optional)
func (r *DocumentCrRepositoryDB) fetchMilestones(ctx context.Context, docCrID uint, item *domainDocCr.DocumentCrDetailItem) {
	rows, err := r.db.WithContext(ctx).Raw(`
		SELECT log_id, doc_cr_id, log_activity, nik_log, name_log, note_log, created_at
		FROM trs.trs_document_cr_log
		WHERE doc_cr_id = ?
		ORDER BY created_at ASC;
	`, docCrID).Rows()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var milestoneModel DocumentCrLogModel
		r.db.ScanRows(rows, &milestoneModel)
		item.Milestones = append(item.Milestones, milestoneModel.toEntity())
	}
}

// fetchRelatedData mengambil semua data terkait untuk list view (best effort)
func (r *DocumentCrRepositoryDB) fetchRelatedData(ctx context.Context, docCrID uint, item *domainDocCr.DocumentCrListItem) {
	// Review (optional)
	var reviewModel ReviewModel
	err := r.db.Raw(`SELECT * FROM trs.trs_review WHERE doc_cr_id = ?;`, docCrID).Scan(&reviewModel).Error
	if err == nil {
		item.Review = reviewModel.toEntity()
	}

	// Approvals (optional)
	approvalRows, err := r.db.Raw(`SELECT * FROM trs.trs_approval WHERE doc_cr_id = ? AND approval_status != 'Cancel';`, docCrID).Rows()
	if err == nil {
		defer approvalRows.Close()
		for approvalRows.Next() {
			var approvalModel ApprovalModel
			r.db.ScanRows(approvalRows, &approvalModel)
			item.Approvals = append(item.Approvals, approvalModel.toEntity())
		}
	}

	// PICs (optional)
	picRows, err := r.db.Raw(`SELECT * FROM trs.trs_pic WHERE doc_cr_id = ?;`, docCrID).Rows()
	if err == nil {
		defer picRows.Close()
		for picRows.Next() {
			var picModel PICModel
			r.db.ScanRows(picRows, &picModel)
			pic := picModel.toEntity()

			if pic.PICDepartment == "IT FUNCTIONAL" {
				item.PICFunctionals = append(item.PICFunctionals, pic)
			}
			if pic.PICDepartment == "IT DEVELOPMENT" {
				item.PICDevelopers = append(item.PICDevelopers, pic)
			}
		}
	}
}
