package handler

import (
	"fmt"
	"time"

	docCrApp "itapp_ddd/application/doc_cr"
	notificationApp "itapp_ddd/application/notification"
	domainDocCr "itapp_ddd/domain/doc_cr"
	"itapp_ddd/pkg"

	"github.com/gofiber/fiber/v2"
)

// DocumentCrHandler adalah handler untuk DocumentCr
type DocumentCrHandler struct {
	service      *docCrApp.Service
	logger       docCrApp.Logger
	notification *notificationApp.Service
}

// NewDocumentCrHandler constructor
func NewDocumentCrHandler(service *docCrApp.Service, logger docCrApp.Logger, notification *notificationApp.Service) *DocumentCrHandler {
	return &DocumentCrHandler{service: service, logger: logger, notification: notification}
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
	ReviewDate   *time.Time `json:"review_date"`
	ReviewStatus *string `json:"review_status"`
	ReviewNote   *string `json:"review_note"`
	DocCrID      *uint   `json:"doc_cr_id"`
}

// Response DTO untuk Approval
type approvalResponse struct {
	ApprovalID         *uint   `json:"approval_id"`
	ApprovalNik        *string `json:"approval_nik"`
	ApprovalName       *string `json:"approval_name"`
	ApprovalDate       *time.Time `json:"approval_date"`
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
	DocCrID             uint                `json:"doc_cr_id" gorm:"primarykey"`
	DocCrName           string              `json:"doc_cr_name"`
	FormId              string              `json:"form_id"`
	DocReference        string              `json:"doc_reference"`
	RequesterNIK        string              `json:"requester_nik"`
	RequesterName       string              `json:"requester_name"`
	RequesterDepartment string              `json:"requester_department"`
	PICDevelopers       []picResponse       `json:"pic_developer" gorm:"many2many:document_cr_pic_developers;"`
	PICFunctionals      []picResponse       `json:"pic_functional" gorm:"many2many:document_cr_pic_functionals;"`
	PICViewer           []picResponse       `json:"pic_viewer" gorm:"many2many:document_cr_pic_viewers;"`
	PICQAs              []picResponse       `json:"pic_qa" gorm:"many2many:document_cr_pic_qas;"`
	ReasonDetail        string              `json:"reason_detail"`
	ProblemDetail       string              `json:"problem_detail"`
	ServerType          string              `json:"server_type"`
	CrType              string              `json:"cr_type"`
	CrPriority          string              `json:"cr_priority"`
	GitMergeRequest     string              `json:"git_merge_request"`
	GitRepository       string              `json:"git_repository"`
	DocNotes            string              `json:"doc_notes"`
	DocAttachment       string              `json:"doc_attachment"`
	DocAttachmentName   string              `json:"doc_attachment_name"`
	DocStatus           string              `json:"doc_status"`
	ApplicationName     string              `json:"application_name"`
	DeployedDate        *time.Time          `json:"deployed_date"`
	CreatedAt           time.Time           `json:"created_at"`
	Approval            []approvalResponse  `json:"approval" gorm:"foreignKey:DocCrID;references:DocCrID"`
	Review              reviewResponse      `json:"review" gorm:"foreignKey:DocCrID;references:DocCrID"`
	Milestones          []milestoneResponse `json:"milestones" gorm:"-"`
}

// ShowAll DocumentCr
// @Summary      Show all DocumentCr
// @Description  Show all DocumentCr
// @Tags         DocumentCr
// @Accept       json
// @Produce      json
// @Param        request body  docCrGetListRequest true "Request body containing the filter parameters"
// @Success      200  {object}  docCrListResponse
// @Failure        400  {object}  domainDocCr.ErrorResponse "Failed to retrieve DocumentCr data"
// @Router       /doc-cr [post]
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

// Show Detail DocumentCr
// @Summary      Show Detail DocumentCr
// @Description  Show Detail DocumentCr
// @Tags         DocumentCr
// @Accept       json
// @Produce      json
// @Param        request body  docCrGetDetailRequest true "Request body containing the filter parameters"
// @Success      200  {object}  []docCrDetailResponse
// @Failure        400  {object}  domainDocCr.ErrorResponse "Failed to retrieve DocumentCr data"
// @Router       /doc-cr-detail [post]
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
		DocAttachmentName:   item.DocCr.DocAttachmentName,
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
		doc.Review = reviewResponse{
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
			PICDepartment: "QA",
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
	DocAttachmentName      []string   `json:"doc_attachment_name"`
	DocStatus              string     `json:"doc_status"`
	ApplicationName        string     `json:"application_name"`
	DeployedDate           *time.Time `json:"deployed_date"`

	DocCrID        uint              `json:"doc_cr_id"`      // Deteksi upsert
	PICDevelopers  []picRequestDTO   `json:"pic_developer"`  // PIC Developer
	PICFunctionals []picRequestDTO   `json:"pic_functional"` // PIC Functional
	PICViewers     []picRequestDTO   `json:"pic_viewer"`     // PIC Viewer
	PICQAs         []picRequestDTO   `json:"pic_qa"`         // PIC QA
	Review         *reviewRequestDTO `json:"review"`         // Review (optional)
	NikLog         string            `json:"nik_log"`        // Log: NIK
	NameLog        string            `json:"name_log"`       // Log: Nama
	NoteLog        string            `json:"note_log"`       // Log: Catatan
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

// Show Add DocumentCr
// @Summary      ADD DocumentCr
// @Description  Add DocumentCr
// @Tags         DocumentCr
// @Accept       json
// @Produce      json
// @Param        request body  docCrCreateRequest true "Request body containing the filter parameters"
// @Success      200  {object}  domainDocCr.ApiResponse
// @Failure        400  {object}  domainDocCr.ErrorResponse "Failed to create DocumentCr"
// @Router       /doc-cr-add [post]
func (h *DocumentCrHandler) Create(c *fiber.Ctx) error {
	req := new(docCrCreateRequest)

	if err := c.BodyParser(req); err != nil {
		h.logger.Error("Failed to parse request body", err)
		return pkg.BadRequestResponse(c, "Failed to parse request body")
	}

	// Map PIC ke domain
	picsDev := make([]domainDocCr.PIC, 0)
	for _, p := range req.PICDevelopers {
		picsDev = append(picsDev, domainDocCr.PIC{PICNik: p.PICNik, PICName: p.PICName, PICDepartment: p.PICDepartment})
	}

	picsFunc := make([]domainDocCr.PIC, 0)
	for _, p := range req.PICFunctionals {
		picsFunc = append(picsFunc, domainDocCr.PIC{PICNik: p.PICNik, PICName: p.PICName, PICDepartment: p.PICDepartment})
	}

	picsView := make([]domainDocCr.PIC, 0)
	for _, p := range req.PICViewers {
		picsView = append(picsView, domainDocCr.PIC{PICNik: p.PICNik, PICName: p.PICName, PICDepartment: p.PICDepartment})
	}

	picsQA := make([]domainDocCr.PIC, 0)
	for _, p := range req.PICQAs {
		picsQA = append(picsQA, domainDocCr.PIC{PICNik: p.PICNik, PICName: p.PICName, PICDepartment: "QA"})
	}

	// Map Review ke domain (pointer, optional)
	var review *domainDocCr.ReviewCreateRequest
	if req.Review != nil {
		review = &domainDocCr.ReviewCreateRequest{
			ReviewNik:  req.Review.ReviewNik,
			ReviewName: req.Review.ReviewName,
		}
	}

	// Map request ke domain request
	domainReq := domainDocCr.CreateDocCrRequest{
		DocCrID:                req.DocCrID,
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
		DocAttachment:          req.DocAttachment,
		DocAttachmentExtension: req.DocAttachmentExtension,
		DocAttachmentName:      req.DocAttachmentName,
		DocStatus:              req.DocStatus,
		ApplicationName:        req.ApplicationName,
		DeployedDate:           req.DeployedDate,
		PICDevelopers:          picsDev,
		PICFunctionals:         picsFunc,
		PICViewers:             picsView,
		PICQAs:                 picsQA,
		Review:                 review,
		NikLog:                 req.NikLog,
		NameLog:                req.NameLog,
		NoteLog:                req.NoteLog,
	}

	doc, err := h.service.Create(c.Context(), domainReq)
	if err != nil {
		h.logger.Error("Documentcr not found", err)
		return pkg.NotFoundResponse(c, "Documentcr not found: "+err.Error())
	}
	if req.Review != nil {
		_ = h.notification.Send(req.Review.ReviewNik, fmt.Sprint("New CR ", doc.FormId, " - ", doc.DocCrName), "New CR needs your review, Please refresh the CR list")
	}
	for _, p := range req.PICFunctionals {
		_ = h.notification.Send(p.PICNik, fmt.Sprint("New CR ", doc.FormId, " - ", doc.DocCrName), "New CR needs your review, Please refresh the CR list")
	}
	for _, p := range req.PICDevelopers {
		_ = h.notification.Send(p.PICNik, fmt.Sprint("New CR ", doc.FormId, " - ", doc.DocCrName), "New CR needs your review, Please refresh the CR list")
	}
	for _, p := range req.PICViewers {
		_ = h.notification.Send(p.PICNik, fmt.Sprint("New CR ", doc.FormId, " - ", doc.DocCrName), "New CR needs your review, Please refresh the CR list")
	}
	for _, p := range req.PICQAs {
		_ = h.notification.Send(p.PICNik, fmt.Sprint("New CR ", doc.FormId, " - ", doc.DocCrName), "New CR needs your review, Please refresh the CR list")
	}

	h.logger.Info("[%s] %s | Status: 200 | DocumentCr created successfully", c.Method(), c.Path())
	return pkg.SuccessResponse(c, doc.DocCrID)
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
	DocAttachmentExtension []string            `json:"doc_attachment_extension"`
	DocAttachmentName      []string             `json:"doc_attachment_name"`
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

// Show Send Review DocumentCr
// @Summary      Send Review DocumentCr
// @Description  Send Review DocumentCr
// @Tags         DocumentCr
// @Accept       json
// @Produce      json
// @Param        request body  docCrSendReviewRequest true "Request body containing the filter parameters"
// @Success      200  {object}  docCrCreateResponse "Review sent successfully"
// @Failure       400  {object}  domainDocCr.ErrorResponse "Failed to send review"
// @Router       /doc-cr-review [post]
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
		picsQA = append(picsQA, domainDocCr.PIC{PICNik: p.PICNik, PICName: p.PICName, PICDepartment: "QA"})
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
		DocAttachmentExtension: req.DocAttachmentExtension,
		DocAttachmentName:      req.DocAttachmentName,
		DocStatus:           "Waiting Approve",
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

	for _, p := range req.PICFunctionals {
		_ = h.notification.Send(p.PICNik, fmt.Sprint("Your CR ", req.FormId, "-", req.DocCrName), "has been reviewed")
	}
	for _, p := range req.PICDevelopers {
		_ = h.notification.Send(p.PICNik, fmt.Sprint("Your CR ", req.FormId, "-", req.DocCrName), "has been reviewed")
	}
	for _, p := range req.PICViewers {
		_ = h.notification.Send(p.PICNik, fmt.Sprint("Your CR ", req.FormId, "-", req.DocCrName), "has been reviewed")
	}
	for _, p := range req.PICQAs {
		_ = h.notification.Send(p.PICNik, fmt.Sprint("Your CR ", req.FormId, "-", req.DocCrName), "has been reviewed")
	}
	for _, a := range req.Approvals {
		_ = h.notification.Send(a.ApprovalNik, fmt.Sprint("New CR ", req.FormId, "-", req.DocCrName), "New CR needs your review, Please refresh the CR list")
	}

	h.logger.Info("[%s] %s | Status: 200 | Review submitted for DocCrID: %d", c.Method(), c.Path(), result.DocCrID)
	return pkg.SuccessResponse(c, nil)
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
	DocAttachment []string `json:"doc_attachment"`
	DocAttachmentExtension []string `json:"doc_attachment_extension"`
	DocAttachmentName []string `json:"doc_attachment_name"`
}

// Show Send Review Reject DocumentCr
// @Summary      Send Review Reject DocumentCr
// @Description  Send Review Reject DocumentCr
// @Tags         DocumentCr
// @Accept       json
// @Produce      json
// @Param        request body  docCrSendReviewRejectRequest true "Request body containing the filter parameters"
// @Success      200  {object}  docCrCreateResponse "Review reject sent successfully"
// @Failure        400  {object}  domainDocCr.ErrorResponse "Failed to reject review"
// @Router       /doc-cr-reject [post]
func (h *DocumentCrHandler) SendReviewReject(c *fiber.Ctx) error {
	req := new(docCrSendReviewRequest)
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
		DocAttachment: req.DocAttachment,
		DocAttachmentExtension: req.DocAttachmentExtension,
		DocAttachmentName:      req.DocAttachmentName,
	}

	// Panggil service
	result, err := h.service.SendReviewReject(c.Context(), domainReq)
	if err != nil {
		h.logger.Error("Failed to reject review", err)
		return pkg.BadRequestResponse(c, "Failed to reject review: "+err.Error())
	}
	_ = h.notification.Send(req.RequesterNIK, fmt.Sprint("New CR ", req.FormId, " - ", req.DocCrName), "New CR needs your revision, Please refresh the CR list")

	h.logger.Info("[%s] %s | Status: 200 | Review rejected for DocCrID: %d", c.Method(), c.Path(), result.DocCrID)
	return pkg.SuccessResponse(c, nil)
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
	DocAttachment       []string           `json:"doc_attachment"`
	DocAttachmentExtension []string        `json:"doc_attachment_extension"`
	DocAttachmentName      []string        `json:"doc_attachment_name"`
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

// Show revision DocumentCr
// @Summary      revision DocumentCr
// @Description  revision DocumentCr
// @Tags         DocumentCr
// @Accept       json
// @Produce      json
// @Param        request body  docCrSendRevisionRequest true "Request body containing the filter parameters"
// @Success      200  {object}  docCrCreateResponse "Revision submitted"
// @Failure        400  {object}  domainDocCr.ErrorResponse "Failed to submit revision"
// @Router       /doc-cr-revision [post]
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
		DocAttachmentExtension: req.DocAttachmentExtension,
		DocAttachmentName:      req.DocAttachmentName,
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

	_ = h.notification.Send(req.Review.ReviewNik, fmt.Sprint("New CR ", req.FormId, "-",req.DocCrName), "New CR needs your review, Please refresh the CR list")

	h.logger.Info("[%s] %s | Status: 200 | Revision submitted for DocCrID: %d", c.Method(), c.Path(), result.DocCrID)
	return pkg.SuccessResponse(c, nil)
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
	ApprovalID         *uint      `json:"approval_id"`
	ApprovalNik        *string    `json:"approval_nik"`
	ApprovalName       *string    `json:"approval_name"`
	ApprovalDate       *time.Time `json:"approval_date"`
	ApprovalStatus     *string    `json:"approval_status"`
	ApprovalNote       *string    `json:"approval_note"`
	DocCrID            *uint      `json:"doc_cr_id"`
	ApprovalDepartment *string    `json:"approval_department"`
	CreatedAt          *time.Time `json:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at"`
}

// Show Approve DocumentCr
// @Summary      Approve DocumentCr
// @Description  Approve DocumentCr
// @Tags         DocumentCr
// @Accept       json
// @Produce      json
// @Param        request body  docCrSendApproveRequest true "Request body containing the filter parameters"
// @Success      200  {object}  approvalApproveResponse "Revision submitted"
// @Failure       400  {object}  domainDocCr.ErrorResponse "Failed to approve document"
// @Router    	/doc-cr-approve [post]
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
			CreatedAt:          a.CreatedAt,
			UpdatedAt:          a.UpdatedAt,
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
	ReviewNote    string `json:"review_note"`
}

// Show approve Reject DocumentCr
// @Summary      approve Reject DocumentCr
// @Description  approve Reject DocumentCr
// @Tags         DocumentCr
// @Accept       json
// @Produce      json
// @Param        request body  docCrSendApproveRejectRequest true "Request body containing the filter parameters"
// @Success      200  {object}  docCrCreateResponse "Approval rejected submitted"
// @Failure        400  {object}  domainDocCr.ErrorResponse "Failed to reject approval"
// @Router       /doc-cr-approve-reject [post]
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
		ReviewNote:	   req.ReviewNote,
	}

	// Panggil service
	result, err := h.service.SendApproveReject(c.Context(), domainReq)
	if err != nil {
		h.logger.Error("Failed to reject approval", err)
		return pkg.BadRequestResponse(c, "Failed to reject approval: "+err.Error())
	}

	h.logger.Info("[%s] %s | Status: 200 | Approval rejected for DocCrID: %d", c.Method(), c.Path(), result.DocCrID)
	return pkg.SuccessResponse(c, nil)
}

// --- Request DTO untuk UpdateStatus ---
type docCrUpdateStatusRequest struct {
	DocCrID   uint   `json:"doc_cr_id"`
	DocStatus string `json:"doc_status"`
	NikLog    string `json:"nik_log"`
	NameLog   string `json:"name_log"`
	NoteLog   string `json:"note_log"`
}

// UpdateStatus mengupdate status dokumen CR
// @Summary      Update Document CR Status
// @Description  Update the status of a Document CR
// @Tags         DocumentCr
// @Accept       json
// @Produce      json
// @Param        request body  docCrUpdateStatusRequest true "Request body containing the filter parameters"
// @Success      200  {object}  docCrCreateResponse "update status"
// @Failure        400  {object}  domainDocCr.ErrorResponse "Failed to update Document CR status"
// @Router       /doc-cr-update-status [post]
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
	DocCrID         uint   `json:"doc_cr_id"`
	NikLog         string `json:"nik_log"`
	NameLog        string `json:"name_log"`
	NoteLog        string `json:"note_log"`
}

type DeleteResponse struct {
	DocCrID   uint      `json:"doc_cr_id"`
}

// Show delete DocumentCr
// @Summary      Delete DocumentCr
// @Description  Delete DocumentCr
// @Tags         DocumentCr
// @Accept       json
// @Produce      json
// @Param        request body  domainDocCr.DeleteRequest true "Request body containing the filter parameters"
// @Success      200  {object}  domainDocCr.ApiResponse  "The data has been deleted"
// @Failure      400  {object}  domainDocCr.ErrorResponse "Failed to delete DocumentCr"
// @Router       /doc-cr-delete [post]
func (h *DocumentCrHandler) Delete(c *fiber.Ctx) error {
	req := new(docCrDeleteRequest)
	if err := c.BodyParser(req); err != nil {
		h.logger.Error("Failed to parse request body", err)
		return pkg.BadRequestResponse(c, "Failed to parse request body")
	}

	domainReq := domainDocCr.DeleteRequest{
		DocCrID:         req.DocCrID,
		NikLog:          req.NikLog,
		NameLog:         req.NameLog,
		NoteLog:         req.NoteLog,
	}

	result, err := h.service.Delete(c.Context(), domainReq)
	if err != nil {
		h.logger.Error("Failed to delete DocumentCr", err)
		return pkg.BadRequestResponse(c, "Failed to delete DocumentCr: "+err.Error())
	}

	h.logger.Info("[%s] %s | Status: 200 | DocumentCr deleted successfully for DocCrID: %d", c.Method(), c.Path(),result.DocCrID)
	return pkg.SuccessResponse(c, "The data has been deleted")
}
