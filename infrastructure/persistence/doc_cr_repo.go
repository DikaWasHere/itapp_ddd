package persistence

import (
	"context"
	"fmt"
	"strconv"
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
	CreatedAt          *time.Time `gorm:"column:created_at"`
	UpdatedAt          *time.Time `gorm:"column:updated_at"`
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
	PICDepartment string    `gorm:"column:pic_department"`
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

	var docIDs []uint

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
		//r.fetchRelatedData(ctx, doc.DocCrID, item)
		docIDs = append(docIDs, doc.DocCrID)

		items = append(items, item)
	}
	// go
	// Batch fetch semua review sekaligus
	reviewMap := make(map[uint]*domainDocCr.Review)
	if len(docIDs) > 0 {
		var reviewModels []ReviewModel
		r.db.Where("doc_cr_id IN ?", docIDs).Find(&reviewModels)
		for _, rev := range reviewModels {
			reviewMap[rev.DocCrID] = rev.toEntity()
		}
	}

	// Batch fetch semua approval sekaligus
	approvalMap := make(map[uint][]*domainDocCr.Approval)
	if len(docIDs) > 0 {
		var approvalModels []ApprovalModel
		r.db.Where("doc_cr_id IN ? AND approval_status != ?", docIDs, "Cancel").Find(&approvalModels)
		for i := range approvalModels {
			appr := approvalModels[i].toEntity()
			if appr.DocCrID != nil { 
				approvalMap[*appr.DocCrID] = append(approvalMap[*appr.DocCrID], appr) 
			}
		}
	}

	// Batch fetch semua PIC sekaligus
	picMap := make(map[uint][]*domainDocCr.PIC)
	if len(docIDs) > 0 {
		var picModels []PICModel
		r.db.Where("doc_cr_id IN ?", docIDs).Find(&picModels)
		for _, picM := range picModels {
			pic := picM.toEntity()
			switch pic.PICDepartment {
			case "IT FUNCTIONAL":
				picMap[picM.DocCrID] = append(picMap[picM.DocCrID], pic)
			case "IT DEVELOPMENT":
				picMap[picM.DocCrID] = append(picMap[picM.DocCrID], pic)
			}
		}
	}

	// Cocokkan data dari map ke item
	for _, item := range items {
		docID := item.DocCr.DocCrID

		if review, ok := reviewMap[docID]; ok {
			item.Review = review
		}

		if approvals, ok := approvalMap[docID]; ok {
			item.Approvals = approvals
		}

		if pics, ok := picMap[docID]; ok {
			for _, pic := range pics {
				switch pic.PICDepartment {
				case "IT FUNCTIONAL":
					item.PICFunctionals = append(item.PICFunctionals, pic)
				case "IT DEVELOPMENT":
					item.PICDevelopers = append(item.PICDevelopers, pic)
				}
			}
		}
	}

	return items, nil
}

// GetDetail mengambil detail DocumentCr berdasarkan DocCrID (optimized with JOIN)
func (r *DocumentCrRepositoryDB) GetDetail(ctx context.Context, req domainDocCr.GetDetailRequest) (*domainDocCr.DocumentCrDetailItem, error) {
	query := `
        SELECT
            -- Document fields
            tdc.doc_cr_id, tdc.doc_cr_name, tdc.form_id, tdc.doc_reference,
            tdc.requester_nik, tdc.requester_name, tdc.requester_department,
            tdc.reason_detail, tdc.problem_detail, tdc.server_type, tdc.cr_type,
            tdc.cr_priority, tdc.git_merge_request, tdc.git_repository,
            tdc.doc_notes, tdc.doc_attachment, tdc.doc_status, tdc.application_name,
            tdc.deployed_date, tdc.created_at,
            -- Approval fields (nullable)
            ta.approval_id, ta.approval_nik, ta.approval_name,
            ta.approval_date, ta.approval_status, ta.approval_note,
            ta.approval_department,
            -- PIC fields (nullable)
            tp.pic_id, tp.pic_nik, tp.pic_name, tp.pic_department AS pic_department,
            -- Review fields (nullable)
            tr.review_id, tr.review_nik, tr.review_name,
            tr.review_date, tr.review_status, tr.review_note
        FROM trs.trs_document_cr tdc
        LEFT JOIN trs.trs_approval ta ON tdc.doc_cr_id = ta.doc_cr_id AND ta.approval_status != 'Cancel'
        LEFT JOIN trs.trs_pic tp ON tdc.doc_cr_id = tp.doc_cr_id
        LEFT JOIN trs.trs_review tr ON tdc.doc_cr_id = tr.doc_cr_id
        WHERE tdc.doc_cr_id = ?
        ORDER BY tdc.created_at DESC;
    `

	var joinRows []DocumentCrDetailJoinModel
	err := r.db.WithContext(ctx).Raw(query, req.DocCrID).Scan(&joinRows).Error
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve DocumentCr detail: %w", err)
	}

	// Step 1: Convert JOIN rows to detail item (without milestones)
	item, err := toDocumentCrDetailItem(joinRows)
	if err != nil {
		return nil, err
	}

	// Step 2: Fetch milestones separately (already exists!)
	r.fetchMilestones(ctx, item.DocCr.DocCrID, item)

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
// func (r *DocumentCrRepositoryDB) fetchPICs(ctx context.Context, docCrID uint, item *domainDocCr.DocumentCrDetailItem) error {
// 	rows, err := r.db.WithContext(ctx).Raw(`
// 		SELECT *
// 		FROM trs.trs_pic
// 		WHERE doc_cr_id = ?;
// 	`, docCrID).Rows()
// 	if err != nil {
// 		return err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var picModel PICModel
// 		r.db.ScanRows(rows, &picModel)
// 		pic := picModel.toEntity()

// 		switch pic.PICDepartment {
// 		case "IT FUNCTIONAL":
// 			item.PICFunctionals = append(item.PICFunctionals, pic)
// 		case "IT DEVELOPMENT":
// 			item.PICDevelopers = append(item.PICDevelopers, pic)
// 		case "IT SERVICES & IMPLEMENTATION":
// 			item.PICViewers = append(item.PICViewers, pic)
// 		case "QA":
// 			item.PICQAs = append(item.PICQAs, pic)
// 		}
// 	}
// 	return nil
// }

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

// DocumentCrDetailJoinModel is a temporary model for JOIN query results
type DocumentCrDetailJoinModel struct {
	// Data from trs_document_cr
	DocCrID             uint       `gorm:"column:doc_cr_id"`
	DocCrName           string     `gorm:"column:doc_cr_name"`
	FormId              string     `gorm:"column:form_id"`
	DocReference        string     `gorm:"column:doc_reference"`
	RequesterNIK        string     `gorm:"column:requester_nik"`
	RequesterName       string     `gorm:"column:requester_name"`
	RequesterDepartment string     `gorm:"column:requester_department"`
	ReasonDetail        string     `gorm:"column:reason_detail"`
	ProblemDetail       string     `gorm:"column:problem_detail"`
	ServerType          string     `gorm:"column:server_type"`
	CrType              string     `gorm:"column:cr_type"`
	CrPriority          string     `gorm:"column:cr_priority"`
	GitMergeRequest     string     `gorm:"column:git_merge_request"`
	GitRepository       string     `gorm:"column:git_repository"`
	DocNotes            string     `gorm:"column:doc_notes"`
	DocAttachment       string     `gorm:"column:doc_attachment"`
	DocStatus           string     `gorm:"column:doc_status"`
	ApplicationName     string     `gorm:"column:application_name"`
	DeployedDate        *time.Time `gorm:"column:deployed_date"`
	CreatedAt           time.Time  `gorm:"column:created_at"`

	// Data from trs_approval (nullable because LEFT JOIN)
	ApprovalID         *uint   `gorm:"column:approval_id"`
	ApprovalNik        *string `gorm:"column:approval_nik"`
	ApprovalName       *string `gorm:"column:approval_name"`
	ApprovalDate       *string `gorm:"column:approval_date"`
	ApprovalStatus     *string `gorm:"column:approval_status"`
	ApprovalNote       *string `gorm:"column:approval_note"`
	ApprovalDepartment *string `gorm:"column:approval_department"`

	// Data from trs_pic (nullable because LEFT JOIN)
	PICID         *uint   `gorm:"column:pic_id"`
	PICNik        *string `gorm:"column:pic_nik"`
	PICName       *string `gorm:"column:pic_name"`
	PICDepartment *string `gorm:"column:pic_department"`

	// Data from trs_review (nullable because LEFT JOIN)
	ReviewID     *uint   `gorm:"column:review_id"`
	ReviewNik    *string `gorm:"column:review_nik"`
	ReviewName   *string `gorm:"column:review_name"`
	ReviewDate   *string `gorm:"column:review_date"`
	ReviewStatus *string `gorm:"column:review_status"`
	ReviewNote   *string `gorm:"column:review_note"`
}

// toDocumentCrDetailItem converts JOIN model rows to domain entity
func toDocumentCrDetailItem(rows []DocumentCrDetailJoinModel) (*domainDocCr.DocumentCrDetailItem, error) {
	if len(rows) == 0 {
		return nil, fmt.Errorf("document not found")
	}

	first := rows[0]

	// Build DocumentCr entity
	doc := &domainDocCr.DocumentCr{
		DocCrID:             first.DocCrID,
		DocCrName:           first.DocCrName,
		FormId:              first.FormId,
		DocReference:        first.DocReference,
		RequesterNIK:        first.RequesterNIK,
		RequesterName:       first.RequesterName,
		RequesterDepartment: first.RequesterDepartment,
		ReasonDetail:        first.ReasonDetail,
		ProblemDetail:       first.ProblemDetail,
		ServerType:          first.ServerType,
		CrType:              first.CrType,
		CrPriority:          first.CrPriority,
		GitMergeRequest:     first.GitMergeRequest,
		GitRepository:       first.GitRepository,
		DocNotes:            first.DocNotes,
		DocAttachment:       first.DocAttachment,
		DocStatus:           first.DocStatus,
		ApplicationName:     first.ApplicationName,
		DeployedDate:        first.DeployedDate,
		CreatedAt:           first.CreatedAt,
	}

	item := &domainDocCr.DocumentCrDetailItem{
		DocCr:          doc,
		Approvals:      make([]*domainDocCr.Approval, 0),
		PICDevelopers:  make([]*domainDocCr.PIC, 0),
		PICFunctionals: make([]*domainDocCr.PIC, 0),
		PICViewers:     make([]*domainDocCr.PIC, 0),
		PICQAs:         make([]*domainDocCr.PIC, 0),
		Milestones:     make([]*domainDocCr.Milestone, 0),
	}

	// Deduplication maps
	seenApprovals := make(map[uint]bool)
	seenPICs := make(map[uint]bool)

	for _, row := range rows {
		// Collect Approvals
		if row.ApprovalID != nil && !seenApprovals[*row.ApprovalID] {
			seenApprovals[*row.ApprovalID] = true
			item.Approvals = append(item.Approvals, &domainDocCr.Approval{
				ApprovalID:         row.ApprovalID,
				ApprovalNik:        row.ApprovalNik,
				ApprovalName:       row.ApprovalName,
				ApprovalDate:       row.ApprovalDate,
				ApprovalStatus:     row.ApprovalStatus,
				ApprovalDepartment: row.ApprovalDepartment,
				DocCrID:            &first.DocCrID,
			})
		}

		// Collect PICs
		if row.PICID != nil && row.PICNik != nil && row.PICName != nil && row.PICDepartment != nil && !seenPICs[*row.PICID] {
			seenPICs[*row.PICID] = true
			pic := &domainDocCr.PIC{
				PICNik:        *row.PICNik,
				PICName:       *row.PICName,
				PICDepartment: *row.PICDepartment,
			}
			switch *row.PICDepartment {
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

		// Collect Review (only first one)
		if row.ReviewID != nil && item.Review == nil {
			item.Review = &domainDocCr.Review{
				ReviewID:     row.ReviewID,
				ReviewNik:    row.ReviewNik,
				ReviewName:   row.ReviewName,
				ReviewDate:   row.ReviewDate,
				ReviewStatus: row.ReviewStatus,
				ReviewNote:   row.ReviewNote,
				DocCrID:      &first.DocCrID,
			}
		}

	}

	return item, nil
}

// ini N+1 karena melakukan 3 query yaitu review, approval, dan pic. Namun, ini dilakukan untuk menghindari query yang terlalu kompleks dan tetap menjaga performa. Jika diperlukan, bisa dioptimalkan dengan join query, tetapi harus dipastikan hasilnya tetap sesuai dengan kebutuhan.
// func (r *DocumentCrRepositoryDB) fetchRelatedData(ctx context.Context, docCrID uint, item *domainDocCr.DocumentCrListItem) {
// 	// Review
// 	var reviewModel ReviewModel
// 	err := r.db.Raw(`SELECT * FROM trs.trs_review WHERE doc_cr_id = ?;`, docCrID).Scan(&reviewModel).Error
// 	if err == nil {
// 		item.Review = reviewModel.toEntity()
// 	}

// 	// Approvals
// 	approvalRows, err := r.db.Raw(`SELECT * FROM trs.trs_approval WHERE doc_cr_id = ? AND approval_status != 'Cancel';`, docCrID).Rows()
// 	if err == nil {
// 		defer approvalRows.Close()
// 		for approvalRows.Next() {
// 			var approvalModel ApprovalModel
// 			r.db.ScanRows(approvalRows, &approvalModel)
// 			item.Approvals = append(item.Approvals, approvalModel.toEntity())
// 		}
// 	}

// 	// PICs
// 	picRows, err := r.db.Raw(`SELECT * FROM trs.trs_pic WHERE doc_cr_id = ?;`, docCrID).Rows()
// 	if err == nil {
// 		defer picRows.Close()
// 		for picRows.Next() {
// 			var picModel PICModel
// 			r.db.ScanRows(picRows, &picModel)
// 			pic := picModel.toEntity()

// 			if pic.PICDepartment == "IT FUNCTIONAL" {
// 				item.PICFunctionals = append(item.PICFunctionals, pic)
// 			}
// 			if pic.PICDepartment == "IT DEVELOPMENT" {
// 				item.PICDevelopers = append(item.PICDevelopers, pic)
// 			}
// 		}
// 	}
// }

// Create membuat DocumentCr baru
func (r *DocumentCrRepositoryDB) Create(ctx context.Context, req domainDocCr.CreateDocCrRequest) (*domainDocCr.DocumentCr, error) {
    // Konversi DocAttachment ([]string) ke string dengan join "|"
    var docAttachmentStr string
    if len(req.DocAttachment) > 0 {
        docAttachmentStr = strings.Join(req.DocAttachment, "|")
    }
 
    docModel := DocumentCrModel{
        DocCrName:           req.DocCrName,
        FormId:              req.FormId,
        DocReference:        req.DocReference,
        RequesterNIK:        req.RequesterNIK,
        RequesterName:       req.RequesterName,
        RequesterDepartment: req.RequesterDepartment,
        ReasonDetail:        req.ReasonDetail,
        ProblemDetail:       req.ProblemDetail,
        ServerType:          req.ServerType,
        CrType:              req.CrType,
        CrPriority:          req.CrPriority,
        GitMergeRequest:     req.GitMergeRequest,
        GitRepository:       req.GitRepository,
        DocNotes:            req.DocNotes,
        DocAttachment:       docAttachmentStr,  
        DocStatus:           req.DocStatus,
        ApplicationName:     req.ApplicationName,
        DeployedDate:        req.DeployedDate,
        CreatedAt:           time.Now(),
        UpdatedAt:           time.Now(),
    }
 
    if err := r.db.WithContext(ctx).Create(&docModel).Error; err != nil {
        return nil, fmt.Errorf("failed to create DocumentCr: %w", err)
    }
 
    return docModel.toEntity(), nil
}

// CountByYear menghitung jumlah DocumentCr berdasarkan tahun (range query berbasis indeks)
func (r *DocumentCrRepositoryDB) CountByYear(ctx context.Context, year string) (int64, error) {
	var count int64

	// Pembatas rentang waktu: awal tahun sampai awal tahun berikutnya
	startOfYear := fmt.Sprintf("%s-01-01 00:00:00", year)

	yearInt, err := strconv.Atoi(year)
	if err != nil {
		return 0, fmt.Errorf("invalid year format: %w", err)
	}
	endOfYear := fmt.Sprintf("%d-01-01 00:00:00", yearInt+1)

	err = r.db.WithContext(ctx).
		Model(&DocumentCrModel{}).
		Where("created_at >= ? AND created_at < ?", startOfYear, endOfYear).
		Count(&count).Error
	if err != nil {
		return 0, fmt.Errorf("failed to count DocumentCr by year: %w", err)
	}
	return count, nil
}

type SendReviewUpdateModel struct {
	DocCrName       string
	FormId          string
	DocReference    string
	ReasonDetail    string
	ProblemDetail   string
	ServerType      string
	CrType          string
	CrPriority      string
	GitMergeRequest string
	GitRepository   string
	DocNotes        string
	DocAttachment   string
	DocStatus       string
	ApplicationName string
}

// SendReview implementasi persistensi untuk review CR
func (r *DocumentCrRepositoryDB) SendReview(ctx context.Context, req domainDocCr.SendReviewRequest) (*domainDocCr.SendReviewResult, error) {
	// 1. Update document CR (partial update)
	updateDoc := SendReviewUpdateModel{
		DocCrName:       req.DocCrName,
		FormId:          req.FormId,
		DocReference:    req.DocReference,
		ReasonDetail:    req.ReasonDetail,
		ProblemDetail:   req.ProblemDetail,
		ServerType:      req.ServerType,
		CrType:          req.CrType,
		CrPriority:      req.CrPriority,
		GitMergeRequest: req.GitMergeRequest,
		GitRepository:   req.GitRepository,
		DocNotes:        req.DocNotes,
		DocAttachment:   strings.Join(req.DocAttachment, "|"),
		DocStatus:       req.DocStatus,
		ApplicationName: req.ApplicationName,
	}

	// guaranteed baris ini dijalankan walau nanti ada error di steps berikutnya
	if err := r.db.WithContext(ctx).Model(&DocumentCrModel{}).Where("doc_cr_id = ?", req.DocCrID).Updates(updateDoc).Error; err != nil {
		return nil, fmt.Errorf("failed to update document cr: %w", err)
	}

	// 2. Log aktivitas
	logModel := DocumentCrLogModel{
		DocCrID:     req.DocCrID,
		LogActivity: "Reviewed",
		NikLog:      req.NikLog,
		NameLog:     req.NameLog,
		NoteLog:     req.NoteLog,
	}
	if err := r.db.WithContext(ctx).Create(&logModel).Error; err != nil {
		return nil, fmt.Errorf("failed to create log: %w", err)
	}

	// 3. Reset PIC (hapus lama, insert baru)
	if err := r.db.WithContext(ctx).Delete(&PICModel{}, "doc_cr_id = ?", req.DocCrID).Error; err != nil {
		return nil, fmt.Errorf("failed to delete old pics: %w", err)
	}

	pics := make([]PICModel, 0, len(req.PICDevelopers)+len(req.PICFunctionals)+len(req.PICViewers)+len(req.PICQAs))
	for i := range req.PICDevelopers {
		pics = append(pics, PICModel{DocCrID: req.DocCrID, PICNik: req.PICDevelopers[i].PICNik, PICName: req.PICDevelopers[i].PICName, PICDepartment: req.PICDevelopers[i].PICDepartment})
	}
	for i := range req.PICFunctionals {
		pics = append(pics, PICModel{DocCrID: req.DocCrID, PICNik: req.PICFunctionals[i].PICNik, PICName: req.PICFunctionals[i].PICName, PICDepartment: req.PICFunctionals[i].PICDepartment})
	}
	for i := range req.PICViewers {
		pics = append(pics, PICModel{DocCrID: req.DocCrID, PICNik: req.PICViewers[i].PICNik, PICName: req.PICViewers[i].PICName, PICDepartment: req.PICViewers[i].PICDepartment})
	}
	for i := range req.PICQAs {
		pics = append(pics, PICModel{DocCrID: req.DocCrID, PICNik: req.PICQAs[i].PICNik, PICName: req.PICQAs[i].PICName, PICDepartment: "QA"})
	}

	if len(pics) > 0 {
		if err := r.db.WithContext(ctx).Create(&pics).Error; err != nil {
			return nil, fmt.Errorf("failed to create new pics: %w", err)
		}
	}

	// 4. Insert approvals
	approvals := make([]ApprovalModel, 0, len(req.Approvals))
	for _, a := range req.Approvals {
		if a == nil || a.ApprovalNik == nil || a.ApprovalName == nil || a.ApprovalDepartment == nil {
			continue
		}
		approvals = append(approvals, ApprovalModel{
			ApprovalNik:        *a.ApprovalNik,
			ApprovalName:       *a.ApprovalName,
			ApprovalDepartment: *a.ApprovalDepartment,
			ApprovalStatus:     "Waiting",
			DocCrID:            req.DocCrID,
		})
	}

	if len(approvals) > 0 {
		if err := r.db.WithContext(ctx).Create(&approvals).Error; err != nil {
			return nil, fmt.Errorf("failed to create approvals: %w", err)
		}
	}

	// 5. Update review
	now := time.Now()
	reviewUpdate := ReviewModel{
		ReviewDate:   func() *string { s := now.Format(time.RFC3339); return &s }(),
		ReviewStatus: "Reviewed",
		ReviewNote:   req.ReviewNote,
	}
	if err := r.db.WithContext(ctx).Model(&ReviewModel{}).Where("doc_cr_id = ?", req.DocCrID).Updates(reviewUpdate).Error; err != nil {
		return nil, fmt.Errorf("failed to update review: %w", err)
	}

	return &domainDocCr.SendReviewResult{DocCrID: req.DocCrID}, nil
}

// SendReviewReject implementasi persistensi untuk reject review CR
func (r *DocumentCrRepositoryDB) SendReviewReject(ctx context.Context, req domainDocCr.SendReviewRejectRequest) (*domainDocCr.SendReviewRejectResult, error) {
	// 1. Update status dokumen ke "Waiting Revision"
	newDocStatus := struct {
		DocStatus string
	}{
		DocStatus: "Waiting Revision",
	}
 
	if err := r.db.WithContext(ctx).Model(&DocumentCrModel{}).Where("doc_cr_id = ?", req.DocCrID).Updates(newDocStatus).Error; err != nil {
		return nil, fmt.Errorf("failed to update document status: %w", err)
	}
 
	// 2. Insert log aktivitas "Rejected"
	logModel := DocumentCrLogModel{
		DocCrID:     req.DocCrID,
		LogActivity: "Rejected",
		NikLog:      req.NikLog,
		NameLog:     req.NameLog,
		NoteLog:     req.NoteLog,
	}
	if err := r.db.WithContext(ctx).Create(&logModel).Error; err != nil {
		return nil, fmt.Errorf("failed to create log: %w", err)
	}
 
	// 3. Update data review
	now := time.Now()
	reviewUpdate := ReviewModel{
		ReviewDate:   func() *string { s := now.Format(time.RFC3339); return &s }(),
		ReviewStatus: "Waiting Revision",
		ReviewNote:   req.ReviewNote,
	}
	if err := r.db.WithContext(ctx).Model(&ReviewModel{}).Where("doc_cr_id = ?", req.DocCrID).Updates(reviewUpdate).Error; err != nil {
		return nil, fmt.Errorf("failed to update review: %w", err)
	}
 
	return &domainDocCr.SendReviewRejectResult{DocCrID: req.DocCrID}, nil
}

// SendRevisionUpdateModel adalah model untuk partial update document CR saat revision
type SendRevisionUpdateModel struct {
	DocCrName       string
	DocReference    string
	ReasonDetail    string
	ProblemDetail   string
	ServerType      string
	CrType          string
	CrPriority      string
	GitMergeRequest string
	GitRepository   string
	DocNotes        string
	DocAttachment   string
	DocStatus       string
	ApplicationName string
}
 
// SendRevision implementasi persistensi untuk merevisi CR
func (r *DocumentCrRepositoryDB) SendRevision(ctx context.Context, req domainDocCr.SendRevisionRequest) (*domainDocCr.SendRevisionResult, error) {
	// 1. Update document CR
	updateDoc := SendRevisionUpdateModel{
		DocCrName:       req.DocCrName,
		DocReference:    req.DocReference,
		ReasonDetail:    req.ReasonDetail,
		ProblemDetail:   req.ProblemDetail,
		ServerType:      req.ServerType,
		CrType:          req.CrType,
		CrPriority:      req.CrPriority,
		GitMergeRequest: req.GitMergeRequest,
		GitRepository:   req.GitRepository,
		DocNotes:        req.DocNotes,
		DocAttachment:   req.DocAttachment,
		DocStatus:       "Waiting Review",
		ApplicationName: req.ApplicationName,
	}
 
	if err := r.db.WithContext(ctx).Model(&DocumentCrModel{}).Where("doc_cr_id = ?", req.DocCrID).Updates(updateDoc).Error; err != nil {
		return nil, fmt.Errorf("failed to update document cr: %w", err)
	}
 
	// 2. Log aktivitas "Revised"
	logModel := DocumentCrLogModel{
		DocCrID:     req.DocCrID,
		LogActivity: "Revised",
		NikLog:      req.NikLog,
		NameLog:     req.NameLog,
		NoteLog:     req.NoteLog,
	}
	if err := r.db.WithContext(ctx).Create(&logModel).Error; err != nil {
		return nil, fmt.Errorf("failed to create log: %w", err)
	}
 
	return &domainDocCr.SendRevisionResult{DocCrID: req.DocCrID}, nil
}

// SendApprove implementasi persistensi untuk approve CR
func (r *DocumentCrRepositoryDB) SendApprove(ctx context.Context, req domainDocCr.SendApproveRequest) (*domainDocCr.SendApproveResult, error) {
 
	// 1. Update approval record (status, tanggal, catatan)
	now := time.Now()
	nowStr := now.Format(time.RFC3339)
	approvalUpdate := struct {
		ApprovalDate   *string
		ApprovalStatus string
		ApprovalNote   string
	}{
		ApprovalDate:   &nowStr,
		ApprovalStatus: req.ApprovalStatus,
		ApprovalNote:   req.ApprovalNote,
	}
 
	if err := r.db.WithContext(ctx).Model(&ApprovalModel{}).Where("approval_id = ?", req.ApprovalID).Updates(approvalUpdate).Error; err != nil {
		return nil, fmt.Errorf("failed to update approval: %w", err)
	}
 
	// 2. Ambil semua approval yang sudah "Approve" untuk CR ini
	var approvedApprovals []ApprovalModel
	if err := r.db.WithContext(ctx).
		Where("doc_cr_id = ? AND approval_status = ?", req.DocCrID, "Approve").
		Find(&approvedApprovals).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch approvals: %w", err)
	}
 
	// 3. Cek business rule
	hasSameDepartment := false
	hasITServices := false
 
	for _, approval := range approvedApprovals {
		if approval.ApprovalDepartment == req.RequesterDepartment {
			hasSameDepartment = true
		}
		if approval.ApprovalDepartment == "IT SERVICES & IMPLEMENTATION" {
			hasITServices = true
		}
	}
 
	// 4. Update status document jika syarat terpenuhi
	if hasSameDepartment {
		shouldDeploy := false
 
		if req.CrPriority == "Not Urgent" {
			// Not Urgent: butuh IT SERVICES juga approve
			if hasITServices {
				shouldDeploy = true
			}
		} else {
			// Urgent: langsung deploy
			shouldDeploy = true
		}
 
		if shouldDeploy {
			docUpdate := struct {
				DocStatus string
			}{
				DocStatus: "Waiting Deploy",
			}
 
			if err := r.db.WithContext(ctx).Model(&DocumentCrModel{}).Where("doc_cr_id = ?", req.DocCrID).Updates(docUpdate).Error; err != nil {
				return nil, fmt.Errorf("failed to update document status: %w", err)
			}
 
			// Buat log "Approved"
			logModel := DocumentCrLogModel{
				DocCrID:     req.DocCrID,
				LogActivity: "Approved",
				NikLog:      req.NikLog,
				NameLog:     req.NameLog,
				NoteLog:     req.NoteLog,
			}
			if err := r.db.WithContext(ctx).Create(&logModel).Error; err != nil {
				return nil, fmt.Errorf("failed to create log: %w", err)
			}
		}
	}
 
	// 5. Map hasil ke domain DTO
	var approveItems []*domainDocCr.ApproveItem
	for i := range approvedApprovals {
		item := &domainDocCr.ApproveItem{
			ApprovalID:         &approvedApprovals[i].ApprovalID,
			ApprovalNik:        &approvedApprovals[i].ApprovalNik,
			ApprovalName:       &approvedApprovals[i].ApprovalName,
			ApprovalDate:       approvedApprovals[i].ApprovalDate,
			ApprovalStatus:     &approvedApprovals[i].ApprovalStatus,
			ApprovalDepartment: &approvedApprovals[i].ApprovalDepartment,
			ApprovalNote:       &approvedApprovals[i].ApprovalNote,
			DocCrID:            &approvedApprovals[i].DocCrID,
			CreatedAt:          approvedApprovals[i].CreatedAt,
			UpdatedAt:          approvedApprovals[i].UpdatedAt,
		}
		approveItems = append(approveItems, item)
	}
 
	return &domainDocCr.SendApproveResult{Approvals: approveItems}, nil
}

// SendApproveReject menolak hasil approval CR
func (r *DocumentCrRepositoryDB) SendApproveReject(ctx context.Context, req domainDocCr.SendApproveRejectRequest) (*domainDocCr.SendApproveRejectResult, error) {
	// 1. Update status dokumen ke "Waiting Review"
	docUpdate := struct {
		DocStatus string
	}{
		DocStatus: "Waiting Review",
	}
 
	if err := r.db.WithContext(ctx).Model(&DocumentCrModel{}).Where("doc_cr_id = ?", req.DocCrID).Updates(docUpdate).Error; err != nil {
		return nil, fmt.Errorf("failed to update document status: %w", err)
	}
 
	// 2. Cancel semua approval_status pada CR tersebut
	if err := r.db.WithContext(ctx).Model(&ApprovalModel{}).Where("doc_cr_id = ?", req.DocCrID).Update("approval_status", "Cancel").Error; err != nil {
		return nil, fmt.Errorf("failed to cancel approvals: %w", err)
	}
 
	// 3. Update data review
	now := time.Now()
	reviewUpdate := ReviewModel{
		ReviewDate:   func() *string { s := now.Format(time.RFC3339); return &s }(),
		ReviewStatus: "Waiting Review",
		ReviewNote:   "Reject Approve",
	}
	if err := r.db.WithContext(ctx).Model(&ReviewModel{}).Where("doc_cr_id = ?", req.DocCrID).Updates(reviewUpdate).Error; err != nil {
		return nil, fmt.Errorf("failed to update review: %w", err)
	}
 
	// 4. Buat log aktivitas "Rejected"
	logModel := DocumentCrLogModel{
		DocCrID:     req.DocCrID,
		LogActivity: "Rejected",
		NikLog:      req.NikLog,
		NameLog:     req.NameLog,
		NoteLog:     req.NoteLog,
	}
	if err := r.db.WithContext(ctx).Create(&logModel).Error; err != nil {
		return nil, fmt.Errorf("failed to create log: %w", err)
	}
 
	return &domainDocCr.SendApproveRejectResult{DocCrID: req.DocCrID}, nil
}

// UpdateStatus mengubah status dokumen CR
func (r *DocumentCrRepositoryDB) UpdateStatus(ctx context.Context, req domainDocCr.UpdateStatusRequest) (*domainDocCr.UpdateStatusResult, error) {
	// 1. Update status dokumen
	docUpdate := struct {
		DocStatus    string
		DeployedDate *time.Time
	}{
		DocStatus: req.DocStatus,
	}
 
	// 2. Jika status "Deployed", set deployed_date dan cancel approval yang masih "Waiting"
	if req.DocStatus == "Deployed" {
		now := time.Now()
		docUpdate.DeployedDate = &now
 
		// Cari semua approval yang masih "Waiting"
		var waitingApprovals []ApprovalModel
		if err := r.db.WithContext(ctx).
			Where("doc_cr_id = ? AND approval_status = ?", req.DocCrID, "Waiting").
			Find(&waitingApprovals).Error; err != nil {
			return nil, fmt.Errorf("failed to fetch waiting approvals: %w", err)
		}
 
		// Update approval yang masih "Waiting" → "No Action"
		for _, approval := range waitingApprovals {
			approvalUpdate := struct {
				ApprovalDate   *string
				ApprovalStatus string
				ApprovalNote   string
			}{
				ApprovalStatus: "No Action",
				ApprovalNote:   "",
			}
			nowStr := time.Now().Format(time.RFC3339)
			approvalUpdate.ApprovalDate = &nowStr
 
			if err := r.db.WithContext(ctx).Model(&ApprovalModel{}).Where("approval_id = ?", approval.ApprovalID).Updates(approvalUpdate).Error; err != nil {
				return nil, fmt.Errorf("failed to update approval status: %w", err)
			}
		}
	}
 
	// 3. Simpan perubahan status dokumen
	if err := r.db.WithContext(ctx).Model(&DocumentCrModel{}).Where("doc_cr_id = ?", req.DocCrID).Updates(docUpdate).Error; err != nil {
		return nil, fmt.Errorf("failed to update document status: %w", err)
	}
 
	// 4. Buat log aktivitas
	logModel := DocumentCrLogModel{
		DocCrID:     req.DocCrID,
		LogActivity: req.DocStatus,
		NikLog:      req.NikLog,
		NameLog:     req.NameLog,
		NoteLog:     req.NoteLog,
	}
	if err := r.db.WithContext(ctx).Create(&logModel).Error; err != nil {
		return nil, fmt.Errorf("failed to create log: %w", err)
	}
 
	return &domainDocCr.UpdateStatusResult{
		DocCrID:   req.DocCrID,
		DocStatus: req.DocStatus,
	}, nil
}


func (r *DocumentCrRepositoryDB) Delete (ctx context.Context,req domainDocCr.DeleteRequest) (*domainDocCr.DeleteResult,error){
	if err := r.db.WithContext(ctx).Delete(&DocumentCrModel{}, "doc_cr_id = ?", req.DocCrID).Error; err != nil {
		return nil, fmt.Errorf("failed to delete document: %w", err)
	}
	return &domainDocCr.DeleteResult{
		DocCrID: req.DocCrID,
		ApplicationName: req.ApplicationName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()}, nil
}