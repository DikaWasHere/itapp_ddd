package handler

import (
	"time"

	docCrApp "itapp_ddd/application/doc_cr"
	domainDocCr "itapp_ddd/domain/doc_cr"
	"itapp_ddd/pkg"

	"github.com/gofiber/fiber/v2"
)

// DocumentCrHandler adalah handler untuk DocumentCr
type DocumentCrHandler struct {
	service *docCrApp.Service
	logger  docCrApp.Logger
}

// NewDocumentCrHandler constructor
func NewDocumentCrHandler(service *docCrApp.Service, logger docCrApp.Logger) *DocumentCrHandler {
	return &DocumentCrHandler{service: service, logger: logger}
}

// Request DTO untuk GetAll
type docCrGetListRequest struct {
	StartDate    string   `json:"start_date"`
	EndDate      string   `json:"end_date"`
	RequesterNik string   `json:"requester_nik"`
	ApprovalNik  string   `json:"approval_nik"`
	DocStatus    []string `json:"doc_status"`
}

// Request DTO untuk GetDetail
type docCrGetDetailRequest struct {
	DocCrID uint `json:"doc_cr_id"`
}

// Response DTO untuk Review
type reviewResponse struct {
	ReviewID     *uint   `json:"review_id"`
	ReviewNik    *string `json:"review_nik"`
	ReviewName   *string `json:"review_name"`
	ReviewDate   *string `json:"review_date"`
	ReviewStatus *string `json:"review_status"`
	ReviewNote   *string `json:"review_note"`
	DocCrID      *uint   `json:"doc_cr_id"`
}

// Response DTO untuk Approval
type approvalResponse struct {
	ApprovalID         *uint   `json:"approval_id"`
	ApprovalNik        *string `json:"approval_nik"`
	ApprovalName       *string `json:"approval_name"`
	ApprovalDate       *string `json:"approval_date"`
	ApprovalStatus     *string `json:"approval_status"`
	ApprovalDepartment *string `json:"approval_department"`
	DocCrID            *uint   `json:"doc_cr_id"`
}

// Response DTO untuk PIC
type picResponse struct {
	PICNik        string `json:"pic_nik"`
	PICName       string `json:"pic_name"`
	PICDepartment string `json:"pic_department"`
}

// Response DTO untuk DocumentCr List
type docCrListResponse struct {
	DocCrID             uint               `json:"doc_cr_id"`
	DocCrName           string             `json:"doc_cr_name"`
	FormId              string             `json:"form_id"`
	DocReference        string             `json:"doc_reference"`
	RequesterNIK        string             `json:"requester_nik"`
	RequesterName       string             `json:"requester_name"`
	RequesterDepartment string             `json:"requester_department"`
	ReasonDetail        string             `json:"reason_detail"`
	ProblemDetail       string             `json:"problem_detail"`
	ServerType          string             `json:"server_type"`
	CrType              string             `json:"cr_type"`
	CrPriority          string             `json:"cr_priority"`
	GitMergeRequest     string             `json:"git_merge_request"`
	GitRepository       string             `json:"git_repository"`
	DocNotes            string             `json:"doc_notes"`
	DocAttachment       string             `json:"doc_attachment"`
	DocStatus           string             `json:"doc_status"`
	ApplicationName     string             `json:"application_name"`
	DeployedDate        *time.Time         `json:"deployed_date"`
	CreatedAt           time.Time          `json:"created_at"`
	PICDevelopers       []picResponse      `json:"pic_developer"`
	PICFunctionals      []picResponse      `json:"pic_functional"`
	Approval            []approvalResponse `json:"approval"`
	Review              *reviewResponse    `json:"review"`
}

// Response DTO untuk Milestone
type milestoneResponse struct {
	LogID       uint      `json:"log_id"`
	DocCrID     uint      `json:"doc_cr_id"`
	LogActivity string    `json:"log_activity"`
	NikLog      string    `json:"nik_log"`
	NameLog     string    `json:"name_log"`
	NoteLog     string    `json:"note_log"`
	CreatedAt   time.Time `json:"created_at"`
}

// Response DTO untuk DocumentCr Detail
type docCrDetailResponse struct {
	DocCrID             uint                `json:"doc_cr_id"`
	DocCrName           string              `json:"doc_cr_name"`
	FormId              string              `json:"form_id"`
	DocReference        string              `json:"doc_reference"`
	RequesterNIK        string              `json:"requester_nik"`
	RequesterName       string              `json:"requester_name"`
	RequesterDepartment string              `json:"requester_department"`
	ReasonDetail        string              `json:"reason_detail"`
	ProblemDetail       string              `json:"problem_detail"`
	ServerType          string              `json:"server_type"`
	CrType              string              `json:"cr_type"`
	CrPriority          string              `json:"cr_priority"`
	GitMergeRequest     string              `json:"git_merge_request"`
	GitRepository       string              `json:"git_repository"`
	DocNotes            string              `json:"doc_notes"`
	DocAttachment       string              `json:"doc_attachment"`
	DocStatus           string              `json:"doc_status"`
	ApplicationName     string              `json:"application_name"`
	DeployedDate        *time.Time          `json:"deployed_date"`
	CreatedAt           time.Time           `json:"created_at"`
	PICDevelopers       []picResponse       `json:"pic_developer"`
	PICFunctionals      []picResponse       `json:"pic_functional"`
	PICViewer           []picResponse       `json:"pic_viewer"`
	PICQAs              []picResponse       `json:"pic_qa"`
	Approval            []approvalResponse  `json:"approval"`
	Review              *reviewResponse     `json:"review"`
	Milestones          []milestoneResponse `json:"milestones"`
}

// GetAll menangani request untuk mengambil semua DocumentCr
func (h *DocumentCrHandler) GetAll(c *fiber.Ctx) error {
	req := new(docCrGetListRequest)

	if err := c.BodyParser(req); err != nil {
		return pkg.BadRequestResponse(c, "Failed to parse request body")
	}

	// Map request ke domain request
	domainReq := domainDocCr.GetListRequest{
		StartDate:    req.StartDate,
		EndDate:      req.EndDate,
		RequesterNik: req.RequesterNik,
		ApprovalNik:  req.ApprovalNik,
		DocStatus:    req.DocStatus,
	}

	// Panggil service
	items, err := h.service.GetAll(c.Context(), domainReq)
	if err != nil {
		return pkg.BadRequestResponse(c, "Failed to retrieve DocumentCr data")
	}

	// Map domain response ke response DTO
	var data []docCrListResponse
	for _, item := range items {
		doc := docCrListResponse{
			DocCrID:             item.DocCr.DocCrID,
			DocCrName:           item.DocCr.DocCrName,
			FormId:              item.DocCr.FormId,
			DocReference:        item.DocCr.DocReference,
			RequesterNIK:        item.DocCr.RequesterNIK,
			RequesterName:       item.DocCr.RequesterName,
			RequesterDepartment: item.DocCr.RequesterDepartment,
			ReasonDetail:        item.DocCr.ReasonDetail,
			ProblemDetail:       item.DocCr.ProblemDetail,
			ServerType:          item.DocCr.ServerType,
			CrType:              item.DocCr.CrType,
			CrPriority:          item.DocCr.CrPriority,
			GitMergeRequest:     item.DocCr.GitMergeRequest,
			GitRepository:       item.DocCr.GitRepository,
			DocNotes:            item.DocCr.DocNotes,
			DocAttachment:       item.DocCr.DocAttachment,
			DocStatus:           item.DocCr.DocStatus,
			ApplicationName:     item.DocCr.ApplicationName,
			DeployedDate:        item.DocCr.DeployedDate,
			CreatedAt:           item.DocCr.CreatedAt,
			PICDevelopers:       make([]picResponse, 0),
			PICFunctionals:      make([]picResponse, 0),
			Approval:            make([]approvalResponse, 0),
		}

		// Map Review
		if item.Review != nil {
			doc.Review = &reviewResponse{
				ReviewID:     item.Review.ReviewID,
				ReviewNik:    item.Review.ReviewNik,
				ReviewName:   item.Review.ReviewName,
				ReviewDate:   item.Review.ReviewDate,
				ReviewStatus: item.Review.ReviewStatus,
				ReviewNote:   item.Review.ReviewNote,
				DocCrID:      item.Review.DocCrID,
			}
		}

		// Map Approval
		for _, a := range item.Approvals {
			doc.Approval = append(doc.Approval, approvalResponse{
				ApprovalID:         a.ApprovalID,
				ApprovalNik:        a.ApprovalNik,
				ApprovalName:       a.ApprovalName,
				ApprovalDate:       a.ApprovalDate,
				ApprovalStatus:     a.ApprovalStatus,
				ApprovalDepartment: a.ApprovalDepartment,
				DocCrID:            a.DocCrID,
			})
		}

		// Map PIC Developer
		for _, p := range item.PICDevelopers {
			doc.PICDevelopers = append(doc.PICDevelopers, picResponse{
				PICNik:        p.PICNik,
				PICName:       p.PICName,
				PICDepartment: p.PICDepartment,
			})
		}

		// Map PIC Functional
		for _, p := range item.PICFunctionals {
			doc.PICFunctionals = append(doc.PICFunctionals, picResponse{
				PICNik:        p.PICNik,
				PICName:       p.PICName,
				PICDepartment: p.PICDepartment,
			})
		}

		data = append(data, doc)
	}

	h.logger.Info("[%s] %s | Status: 200 | Berhasil mengambil semua data", c.Method(), c.Path())
	return pkg.SuccessResponse(c, data)
}

// GetDetail menangani request untuk mengambil detail DocumentCr
func (h *DocumentCrHandler) GetDetail(c *fiber.Ctx) error {
	req := new(docCrGetDetailRequest)

	if err := c.BodyParser(req); err != nil {
		h.logger.Error("Failed to parse request body", err)
		return pkg.BadRequestResponse(c, "Failed to parse request body")
	}

	// Map request ke domain request
	domainReq := domainDocCr.GetDetailRequest{
		DocCrID: req.DocCrID,
	}

	// Panggil service
	item, err := h.service.GetDetail(c.Context(), domainReq)
	if err != nil {
		h.logger.Error("Failed to parse request body", err)
		return pkg.BadRequestResponse(c, "Failed to retrieve DocumentCr data")
	}

	// Map domain response ke response DTO
	doc := docCrDetailResponse{
		DocCrID:             item.DocCr.DocCrID,
		DocCrName:           item.DocCr.DocCrName,
		FormId:              item.DocCr.FormId,
		DocReference:        item.DocCr.DocReference,
		RequesterNIK:        item.DocCr.RequesterNIK,
		RequesterName:       item.DocCr.RequesterName,
		RequesterDepartment: item.DocCr.RequesterDepartment,
		ReasonDetail:        item.DocCr.ReasonDetail,
		ProblemDetail:       item.DocCr.ProblemDetail,
		ServerType:          item.DocCr.ServerType,
		CrType:              item.DocCr.CrType,
		CrPriority:          item.DocCr.CrPriority,
		GitMergeRequest:     item.DocCr.GitMergeRequest,
		GitRepository:       item.DocCr.GitRepository,
		DocNotes:            item.DocCr.DocNotes,
		DocAttachment:       item.DocCr.DocAttachment,
		DocStatus:           item.DocCr.DocStatus,
		ApplicationName:     item.DocCr.ApplicationName,
		DeployedDate:        item.DocCr.DeployedDate,
		CreatedAt:           item.DocCr.CreatedAt,
		PICDevelopers:       make([]picResponse, 0),
		PICFunctionals:      make([]picResponse, 0),
		PICViewer:           make([]picResponse, 0),
		PICQAs:              make([]picResponse, 0),
		Approval:            make([]approvalResponse, 0),
		Milestones:          make([]milestoneResponse, 0),
	}

	// Map Review
	if item.Review != nil {
		doc.Review = &reviewResponse{
			ReviewID:     item.Review.ReviewID,
			ReviewNik:    item.Review.ReviewNik,
			ReviewName:   item.Review.ReviewName,
			ReviewDate:   item.Review.ReviewDate,
			ReviewStatus: item.Review.ReviewStatus,
			ReviewNote:   item.Review.ReviewNote,
			DocCrID:      item.Review.DocCrID,
		}
	}

	// Map Approval
	for _, a := range item.Approvals {
		doc.Approval = append(doc.Approval, approvalResponse{
			ApprovalID:         a.ApprovalID,
			ApprovalNik:        a.ApprovalNik,
			ApprovalName:       a.ApprovalName,
			ApprovalDate:       a.ApprovalDate,
			ApprovalStatus:     a.ApprovalStatus,
			ApprovalDepartment: a.ApprovalDepartment,
			DocCrID:            a.DocCrID,
		})
	}

	// Map PIC Developer
	for _, p := range item.PICDevelopers {
		doc.PICDevelopers = append(doc.PICDevelopers, picResponse{
			PICNik:        p.PICNik,
			PICName:       p.PICName,
			PICDepartment: p.PICDepartment,
		})
	}

	// Map PIC Functional
	for _, p := range item.PICFunctionals {
		doc.PICFunctionals = append(doc.PICFunctionals, picResponse{
			PICNik:        p.PICNik,
			PICName:       p.PICName,
			PICDepartment: p.PICDepartment,
		})
	}

	// Map PIC Viewer
	for _, p := range item.PICViewers {
		doc.PICViewer = append(doc.PICViewer, picResponse{
			PICNik:        p.PICNik,
			PICName:       p.PICName,
			PICDepartment: p.PICDepartment,
		})
	}

	// Map PIC QA
	for _, p := range item.PICQAs {
		doc.PICQAs = append(doc.PICQAs, picResponse{
			PICNik:        p.PICNik,
			PICName:       p.PICName,
			PICDepartment: p.PICDepartment,
		})
	}

	// Map Milestones
	for _, m := range item.Milestones {
		doc.Milestones = append(doc.Milestones, milestoneResponse{
			LogID:       m.LogID,
			DocCrID:     m.DocCrID,
			LogActivity: m.LogActivity,
			NikLog:      m.NikLog,
			NameLog:     m.NameLog,
			NoteLog:     m.NoteLog,
			CreatedAt:   m.CreatedAt,
		})
	}

	h.logger.Info("[%s] %s | Status: 200 | Berhasil mengambil semua data", c.Method(), c.Path())
	return pkg.SuccessResponse(c, doc)
}

// --- Request DTO untuk Create ---
type docCrCreateRequest struct {
	DocCrName              string     `json:"doc_cr_name"`
	FormId                 string     `json:"form_id"`
	DocReference           string     `json:"doc_reference"`
	RequesterNIK           string     `json:"requester_nik"`
	RequesterName          string     `json:"requester_name"`
	RequesterDepartment    string     `json:"requester_department"`
	ReasonDetail           string     `json:"reason_detail"`
	ProblemDetail          string     `json:"problem_detail"`
	ServerType             string     `json:"server_type"`
	CrType                 string     `json:"cr_type"`
	CrPriority             string     `json:"cr_priority"`
	GitMergeRequest        string     `json:"git_merge_request"`
	GitRepository          string     `json:"git_repository"`
	DocNotes               string     `json:"doc_notes"`
	DocAttachment          []string   `json:"doc_attachment"`
	DocAttachmentExtension []string   `json:"doc_attachment_extension"`
	DocStatus              string     `json:"doc_status"`
	ApplicationName        string     `json:"application_name"`
	DeployedDate           *time.Time `json:"deployed_date"`
}

// --- Response DTO untuk Create ---
type docCrCreateResponse struct {
	DocCrID             uint       `json:"doc_cr_id"`
	DocCrName           string     `json:"doc_cr_name"`
	FormId              string     `json:"form_id"`
	DocReference        string     `json:"doc_reference"`
	RequesterNIK        string     `json:"requester_nik"`
	RequesterName       string     `json:"requester_name"`
	RequesterDepartment string     `json:"requester_department"`
	ReasonDetail        string     `json:"reason_detail"`
	ProblemDetail       string     `json:"problem_detail"`
	ServerType          string     `json:"server_type"`
	CrType              string     `json:"cr_type"`
	CrPriority          string     `json:"cr_priority"`
	GitMergeRequest     string     `json:"git_merge_request"`
	GitRepository       string     `json:"git_repository"`
	DocNotes            string     `json:"doc_notes"`
	DocAttachment       string     `json:"doc_attachment"`
	DocStatus           string     `json:"doc_status"`
	ApplicationName     string     `json:"application_name"`
	DeployedDate        *time.Time `json:"deployed_date"`
	CreatedAt           time.Time  `json:"created_at"`
}

// go
// Create menangani request untuk membuat DocumentCr baru
func (h *DocumentCrHandler) Create(c *fiber.Ctx) error {
	req := new(docCrCreateRequest)

	if err := c.BodyParser(req); err != nil {
		h.logger.Error("Failed to parse request body", err)
		return pkg.BadRequestResponse(c, "Failed to parse request body")
	}

	// Map request ke domain request
	domainReq := domainDocCr.CreateDocCrRequest{
		DocCrName:              req.DocCrName,
		FormId:                 req.FormId,
		DocReference:           req.DocReference,
		RequesterNIK:           req.RequesterNIK,
		RequesterName:          req.RequesterName,
		RequesterDepartment:    req.RequesterDepartment,
		ReasonDetail:           req.ReasonDetail,
		ProblemDetail:          req.ProblemDetail,
		ServerType:             req.ServerType,
		CrType:                 req.CrType,
		CrPriority:             req.CrPriority,
		GitMergeRequest:        req.GitMergeRequest,
		GitRepository:          req.GitRepository,
		DocNotes:               req.DocNotes,
		DocAttachment:          req.DocAttachment,          // Now []string
		DocAttachmentExtension: req.DocAttachmentExtension, // NEW
		DocStatus:              req.DocStatus,
		ApplicationName:        req.ApplicationName,
		DeployedDate:           req.DeployedDate,
	}

	// Panggil service
	doc, err := h.service.Create(c.Context(), domainReq)
	if err != nil {
		h.logger.Error("Failed to create DocumentCr", err)
		return pkg.BadRequestResponse(c, "Failed to create DocumentCr: "+err.Error())
	}

	// Map domain response ke response DTO
	data := docCrCreateResponse{
		DocCrID:             doc.DocCrID,
		DocCrName:           doc.DocCrName,
		FormId:              doc.FormId,
		DocReference:        doc.DocReference,
		RequesterNIK:        doc.RequesterNIK,
		RequesterName:       doc.RequesterName,
		RequesterDepartment: doc.RequesterDepartment,
		ReasonDetail:        doc.ReasonDetail,
		ProblemDetail:       doc.ProblemDetail,
		ServerType:          doc.ServerType,
		CrType:              doc.CrType,
		CrPriority:          doc.CrPriority,
		GitMergeRequest:     doc.GitMergeRequest,
		GitRepository:       doc.GitRepository,
		DocNotes:            doc.DocNotes,
		DocAttachment:       doc.DocAttachment,
		DocStatus:           doc.DocStatus,
		ApplicationName:     doc.ApplicationName,
		DeployedDate:        doc.DeployedDate,
		CreatedAt:           doc.CreatedAt,
	}

	h.logger.Info("[%s] %s | Status: 201 | DocumentCr created successfully", c.Method(), c.Path())
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  fiber.StatusCreated,
		"message": "DocumentCr created successfully",
		"data":    data,
	})
}

// --- Request DTO untuk SendReview (simplified) ---
type docCrSendReviewRequest struct {
	DocCrID             uint                 `json:"doc_cr_id"`
	DocCrName           string               `json:"doc_cr_name"`
	FormId              string               `json:"form_id"`
	DocReference        string               `json:"doc_reference"`
	RequesterNIK        string               `json:"requester_nik"`
	RequesterName       string               `json:"requester_name"`
	RequesterDepartment string               `json:"requester_department"`
	PICDevelopers       []picRequestDTO      `json:"pic_developer"`
	PICFunctionals      []picRequestDTO      `json:"pic_functional"`
	PICViewers          []picRequestDTO      `json:"pic_viewer"`
	PICQAs              []picRequestDTO      `json:"pic_qa"`
	ReasonDetail        string               `json:"reason_detail"`
	ProblemDetail       string               `json:"problem_detail"`
	ServerType          string               `json:"server_type"`
	CrType              string               `json:"cr_type"`
	CrPriority          string               `json:"cr_priority"`
	GitMergeRequest     string               `json:"git_merge_request"`
	GitRepository       string               `json:"git_repository"`
	DocNotes            string               `json:"doc_notes"`
	DocAttachment       []string             `json:"doc_attachment"`
	DocStatus           string               `json:"doc_status"`
	ApplicationName     string               `json:"application_name"`
	ReviewNote          string               `json:"review_note"`
	NikLog              string               `json:"nik_log"`
	NameLog             string               `json:"name_log"`
	NoteLog             string               `json:"note_log"`
	Approvals           []approvalRequestDTO `json:"approval"`
}

type picRequestDTO struct {
	PICNik        string `json:"pic_nik"`
	PICName       string `json:"pic_name"`
	PICDepartment string `json:"pic_department"`
}

type approvalRequestDTO struct {
	ApprovalNik        string `json:"approval_nik"`
	ApprovalName       string `json:"approval_name"`
	ApprovalDepartment string `json:"approval_department"`
}

// SendReview menangani request untuk submit review CR
func (h *DocumentCrHandler) SendReview(c *fiber.Ctx) error {
	req := new(docCrSendReviewRequest)
	if err := c.BodyParser(req); err != nil {
		h.logger.Error("Failed to parse request body", err)
		return pkg.BadRequestResponse(c, "Failed to parse request body")
	}

	// Map ke domain request
	picsDev := make([]domainDocCr.PIC, 0, len(req.PICDevelopers))
	for _, p := range req.PICDevelopers {
		picsDev = append(picsDev, domainDocCr.PIC{PICNik: p.PICNik, PICName: p.PICName, PICDepartment: p.PICDepartment})
	}

	picsFunc := make([]domainDocCr.PIC, 0, len(req.PICFunctionals))
	for _, p := range req.PICFunctionals {
		picsFunc = append(picsFunc, domainDocCr.PIC{PICNik: p.PICNik, PICName: p.PICName, PICDepartment: p.PICDepartment})
	}

	picsView := make([]domainDocCr.PIC, 0, len(req.PICViewers))
	for _, p := range req.PICViewers {
		picsView = append(picsView, domainDocCr.PIC{PICNik: p.PICNik, PICName: p.PICName, PICDepartment: p.PICDepartment})
	}

	picsQA := make([]domainDocCr.PIC, 0, len(req.PICQAs))
	for _, p := range req.PICQAs {
		picsQA = append(picsQA, domainDocCr.PIC{PICNik: p.PICNik, PICName: p.PICName, PICDepartment: p.PICDepartment})
	}

	approvals := make([]*domainDocCr.Approval, 0, len(req.Approvals))
	for _, a := range req.Approvals {
		aid := a.ApprovalNik
		and := a.ApprovalName
		adept := a.ApprovalDepartment
		approvals = append(approvals, &domainDocCr.Approval{
			ApprovalNik:        &aid,
			ApprovalName:       &and,
			ApprovalDepartment: &adept,
		})
	}

	domainReq := domainDocCr.SendReviewRequest{
		DocCrID:             req.DocCrID,
		DocCrName:           req.DocCrName,
		FormId:              req.FormId,
		DocReference:        req.DocReference,
		RequesterNIK:        req.RequesterNIK,
		RequesterName:       req.RequesterName,
		RequesterDepartment: req.RequesterDepartment,
		PICDevelopers:       picsDev,
		PICFunctionals:      picsFunc,
		PICViewers:          picsView,
		PICQAs:              picsQA,
		ReasonDetail:        req.ReasonDetail,
		ProblemDetail:       req.ProblemDetail,
		ServerType:          req.ServerType,
		CrType:              req.CrType,
		CrPriority:          req.CrPriority,
		GitMergeRequest:     req.GitMergeRequest,
		GitRepository:       req.GitRepository,
		DocNotes:            req.DocNotes,
		DocAttachment:       req.DocAttachment,
		DocStatus:           req.DocStatus,
		ApplicationName:     req.ApplicationName,
		ReviewNote:          req.ReviewNote,
		NikLog:              req.NikLog,
		NameLog:             req.NameLog,
		NoteLog:             req.NoteLog,
		Approvals:           approvals,
	}

	result, err := h.service.SendReview(c.Context(), domainReq)
	if err != nil {
		h.logger.Error("Failed to send review", err)
		return pkg.BadRequestResponse(c, "Failed to send review: "+err.Error())
	}

	h.logger.Info("[%s] %s | Status: 200 | Review submitted for DocCrID: %d", c.Method(), c.Path(), result.DocCrID)
	return pkg.SuccessResponse(c, fiber.Map{
		"doc_cr_id": result.DocCrID,
	})
}

// --- Request DTO untuk SendReviewReject ---
type docCrSendReviewRejectRequest struct {
	DocCrID      uint   `json:"doc_cr_id"`
	RequesterNIK string `json:"requester_nik"`
	FormId       string `json:"form_id"`
	ReviewNote   string `json:"review_note"`
	NikLog       string `json:"nik_log"`
	NameLog      string `json:"name_log"`
	NoteLog      string `json:"note_log"`
}

// SendReviewReject menangani request untuk menolak CR dari tahap review
func (h *DocumentCrHandler) SendReviewReject(c *fiber.Ctx) error {
	req := new(docCrSendReviewRejectRequest)
	if err := c.BodyParser(req); err != nil {
		h.logger.Error("Failed to parse request body", err)
		return pkg.BadRequestResponse(c, "Failed to parse request body")
	}

	// Validasi input
	if req.DocCrID == 0 {
		return pkg.BadRequestResponse(c, "doc_cr_id is required")
	}

	// Map ke domain request
	domainReq := domainDocCr.SendReviewRejectRequest{
		DocCrID:      req.DocCrID,
		RequesterNIK: req.RequesterNIK,
		FormId:       req.FormId,
		ReviewNote:   req.ReviewNote,
		NikLog:       req.NikLog,
		NameLog:      req.NameLog,
		NoteLog:      req.NoteLog,
	}

	// Panggil service
	result, err := h.service.SendReviewReject(c.Context(), domainReq)
	if err != nil {
		h.logger.Error("Failed to reject review", err)
		return pkg.BadRequestResponse(c, "Failed to reject review: "+err.Error())
	}

	h.logger.Info("[%s] %s | Status: 200 | Review rejected for DocCrID: %d", c.Method(), c.Path(), result.DocCrID)
	return pkg.SuccessResponse(c, fiber.Map{
		"doc_cr_id": result.DocCrID,
	})
}

// --- Request DTO untuk SendRevision ---
type docCrSendRevisionRequest struct {
	DocCrID             uint             `json:"doc_cr_id"`
	DocCrName           string           `json:"doc_cr_name"`
	FormId              string           `json:"form_id"`
	DocReference        string           `json:"doc_reference"`
	RequesterNIK        string           `json:"requester_nik"`
	RequesterName       string           `json:"requester_name"`
	RequesterDepartment string           `json:"requester_department"`
	PICDevelopers       []picRequestDTO  `json:"pic_developer"`
	PICFunctionals      []picRequestDTO  `json:"pic_functional"`
	PICViewers          []picRequestDTO  `json:"pic_viewer"`
	PICQAs              []picRequestDTO  `json:"pic_qa"`
	ReasonDetail        string           `json:"reason_detail"`
	ProblemDetail       string           `json:"problem_detail"`
	ServerType          string           `json:"server_type"`
	CrType              string           `json:"cr_type"`
	CrPriority          string           `json:"cr_priority"`
	GitMergeRequest     string           `json:"git_merge_request"`
	GitRepository       string           `json:"git_repository"`
	DocNotes            string           `json:"doc_notes"`
	DocAttachment       string           `json:"doc_attachment"`
	ApplicationName     string           `json:"application_name"`
	ReviewNote          string           `json:"review_note"`
	NikLog              string           `json:"nik_log"`
	NameLog             string           `json:"name_log"`
	NoteLog             string           `json:"note_log"`
	Review              reviewRequestDTO `json:"review"`
}

type reviewRequestDTO struct {
	ReviewNik  string `json:"review_nik"`
	ReviewName string `json:"review_name"`
}

// SendRevision menangani request untuk merevisi CR
func (h *DocumentCrHandler) SendRevision(c *fiber.Ctx) error {
	req := new(docCrSendRevisionRequest)
	if err := c.BodyParser(req); err != nil {
		h.logger.Error("Failed to parse request body", err)
		return pkg.BadRequestResponse(c, "Failed to parse request body")
	}

	// Validasi input
	if req.DocCrID == 0 {
		return pkg.BadRequestResponse(c, "doc_cr_id is required")
	}

	// Map ke domain request
	domainReq := domainDocCr.SendRevisionRequest{
		DocCrID:             req.DocCrID,
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
		DocAttachment:       req.DocAttachment,
		ApplicationName:     req.ApplicationName,
		ReviewNote:          req.ReviewNote,
		NikLog:              req.NikLog,
		NameLog:             req.NameLog,
		NoteLog:             req.NoteLog,
		Review: domainDocCr.ReviewCreateRequest{
			ReviewNik:  req.Review.ReviewNik,
			ReviewName: req.Review.ReviewName,
		},
	}

	// Panggil service
	result, err := h.service.SendRevision(c.Context(), domainReq)
	if err != nil {
		h.logger.Error("Failed to send revision", err)
		return pkg.BadRequestResponse(c, "Failed to send revision: "+err.Error())
	}

	h.logger.Info("[%s] %s | Status: 200 | Revision submitted for DocCrID: %d", c.Method(), c.Path(), result.DocCrID)
	return pkg.SuccessResponse(c, fiber.Map{
		"doc_cr_id": result.DocCrID,
	})
}

// --- Request DTO untuk SendApprove ---
type docCrSendApproveRequest struct {
	ApprovalID          uint   `json:"approval_id"`
	ApprovalStatus      string `json:"approval_status"`
	ApprovalNote        string `json:"approval_note"`
	NikLog              string `json:"nik_log"`
	NameLog             string `json:"name_log"`
	NoteLog             string `json:"note_log"`
	DocCrID             uint   `json:"doc_cr_id"`
	CrPriority          string `json:"cr_priority"`
	RequesterDepartment string `json:"requester_department"`
}

// --- Response DTO untuk SendApprove (termasuk ApprovalNote) ---
type approvalApproveResponse struct {
	ApprovalID         *uint   `json:"approval_id"`
	ApprovalNik        *string `json:"approval_nik"`
	ApprovalName       *string `json:"approval_name"`
	ApprovalDate       *string `json:"approval_date"`
	ApprovalStatus     *string `json:"approval_status"`
	ApprovalDepartment *string `json:"approval_department"`
	ApprovalNote       *string `json:"approval_note"`
	DocCrID            *uint   `json:"doc_cr_id"`
}

// SendApprove menangani request untuk approve CR
func (h *DocumentCrHandler) SendApprove(c *fiber.Ctx) error {
	// 1. Parse JSON
	req := new(docCrSendApproveRequest)
	if err := c.BodyParser(req); err != nil {
		h.logger.Error("Failed to parse request body", err)
		return pkg.BadRequestResponse(c, "Failed to parse request body")
	}

	// 2. Validasi input
	if req.ApprovalID == 0 {
		return pkg.BadRequestResponse(c, "approval_id is required")
	}
	if req.DocCrID == 0 {
		return pkg.BadRequestResponse(c, "doc_cr_id is required")
	}

	// 3. Map ke domain request
	domainReq := domainDocCr.SendApproveRequest{
		ApprovalID:          req.ApprovalID,
		ApprovalStatus:      req.ApprovalStatus,
		ApprovalNote:        req.ApprovalNote,
		NikLog:              req.NikLog,
		NameLog:             req.NameLog,
		NoteLog:             req.NoteLog,
		DocCrID:             req.DocCrID,
		CrPriority:          req.CrPriority,
		RequesterDepartment: req.RequesterDepartment,
	}

	// 4. Panggil service
	result, err := h.service.SendApprove(c.Context(), domainReq)
	if err != nil {
		h.logger.Error("Failed to send approve", err)
		return pkg.BadRequestResponse(c, "Failed to send approve: "+err.Error())
	}

	// 5. Map result ke response DTO
	var approvals []approvalApproveResponse
	for _, a := range result.Approvals {
		approvals = append(approvals, approvalApproveResponse{
			ApprovalID:         a.ApprovalID,
			ApprovalNik:        a.ApprovalNik,
			ApprovalName:       a.ApprovalName,
			ApprovalDate:       a.ApprovalDate,
			ApprovalStatus:     a.ApprovalStatus,
			ApprovalDepartment: a.ApprovalDepartment,
			ApprovalNote:       a.ApprovalNote,
			DocCrID:            a.DocCrID,
		})
	}

	// 6. Return response
	h.logger.Info("[%s] %s | Status: 200 | Approval submitted for DocCrID: %d", c.Method(), c.Path(), req.DocCrID)
	return pkg.SuccessResponse(c, approvals)
}

// --- Request DTO untuk SendApproveReject ---
type docCrSendApproveRejectRequest struct {
	DocCrID       uint   `json:"doc_cr_id"`
	RequesterName string `json:"requester_name"`
	NikLog        string `json:"nik_log"`
	NameLog       string `json:"name_log"`
	NoteLog       string `json:"note_log"`
}

// SendApproveReject menangani request untuk menolak hasil approval CR
func (h *DocumentCrHandler) SendApproveReject(c *fiber.Ctx) error {
	req := new(docCrSendApproveRejectRequest)
	if err := c.BodyParser(req); err != nil {
		h.logger.Error("Failed to parse request body", err)
		return pkg.BadRequestResponse(c, "Failed to parse request body")
	}

	// Validasi input
	if req.DocCrID == 0 {
		return pkg.BadRequestResponse(c, "doc_cr_id is required")
	}

	// Map ke domain request
	domainReq := domainDocCr.SendApproveRejectRequest{
		DocCrID:       req.DocCrID,
		RequesterName: req.RequesterName,
		NikLog:        req.NikLog,
		NameLog:       req.NameLog,
		NoteLog:       req.NoteLog,
	}

	// Panggil service
	result, err := h.service.SendApproveReject(c.Context(), domainReq)
	if err != nil {
		h.logger.Error("Failed to reject approval", err)
		return pkg.BadRequestResponse(c, "Failed to reject approval: "+err.Error())
	}

	h.logger.Info("[%s] %s | Status: 200 | Approval rejected for DocCrID: %d", c.Method(), c.Path(), result.DocCrID)
	return pkg.SuccessResponse(c, fiber.Map{
		"doc_cr_id": result.DocCrID,
	})
}

// --- Request DTO untuk UpdateStatus ---
type docCrUpdateStatusRequest struct {
	DocCrID   uint   `json:"doc_cr_id"`
	DocStatus string `json:"doc_status"`
	NikLog    string `json:"nik_log"`
	NameLog   string `json:"name_log"`
	NoteLog   string `json:"note_log"`
}

// UpdateStatus menangani request untuk update status dokumen CR
func (h *DocumentCrHandler) UpdateStatus(c *fiber.Ctx) error {
	req := new(docCrUpdateStatusRequest)
	if err := c.BodyParser(req); err != nil {
		h.logger.Error("Failed to parse request body", err)
		return pkg.BadRequestResponse(c, "Failed to parse request body")
	}

	// Validasi input
	if req.DocCrID == 0 {
		return pkg.BadRequestResponse(c, "doc_cr_id is required")
	}
	if req.DocStatus == "" {
		return pkg.BadRequestResponse(c, "doc_status is required")
	}

	// Map ke domain request
	domainReq := domainDocCr.UpdateStatusRequest{
		DocCrID:   req.DocCrID,
		DocStatus: req.DocStatus,
		NikLog:    req.NikLog,
		NameLog:   req.NameLog,
		NoteLog:   req.NoteLog,
	}

	// Panggil service
	result, err := h.service.UpdateStatus(c.Context(), domainReq)
	if err != nil {
		h.logger.Error("Failed to update status", err)
		return pkg.BadRequestResponse(c, "Failed to update status: "+err.Error())
	}

	h.logger.Info("[%s] %s | Status: 200 | Document status updated to '%s' for DocCrID: %d", c.Method(), c.Path(), result.DocStatus, result.DocCrID)
	return pkg.SuccessResponse(c, result.DocCrID)
}

type docCrDeleteRequest struct {
	DocCrID uint `json:"application_id"`
	ApplicationName string `json:"application_name"`
}

func (h *DocumentCrHandler) Delete(c *fiber.Ctx) error {
	req := new(docCrDeleteRequest)
	if err := c.BodyParser(req); err != nil {
		h.logger.Error("Failed to parse request body", err)
		return pkg.BadRequestResponse(c, "Failed to parse request body")
	}

	if req.DocCrID == 0 {
		return pkg.BadRequestResponse(c, "doc_cr_id is required")
	}

	domainReq := domainDocCr.DeleteRequest{
		DocCrID: req.DocCrID,
		ApplicationName: req.ApplicationName,
	}

	result, err := h.service.Delete(c.Context(), domainReq)
	if err != nil {
		h.logger.Error("Failed to delete DocumentCr", err)
		return pkg.BadRequestResponse(c, "Failed to delete DocumentCr: "+err.Error())
	}

	h.logger.Info("[%s] %s | Status: 200 | DocumentCr deleted successfully for DocCrID: %d", c.Method(), c.Path(), req.DocCrID)
	return pkg.SuccessResponse(c, fiber.Map{
		"application_id":   result.DocCrID,
		"application_name": result.ApplicationName,
		"created_at":      result.CreatedAt,
		"updated_at":      result.UpdatedAt,
	})
}
