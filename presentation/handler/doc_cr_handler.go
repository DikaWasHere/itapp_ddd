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
}

// NewDocumentCrHandler constructor
func NewDocumentCrHandler(service *docCrApp.Service) *DocumentCrHandler {
	return &DocumentCrHandler{service: service}
}

// Request DTO untuk GetAll
type docCrGetListRequest struct {
	StartDate    string   `json:"start_date"`
	EndDate      string   `json:"end_date"`
	RequesterNik string   `json:"requester_nik"`
	ApprovalNik  string   `json:"approval_nik"`
	DocStatus    []string `json:"doc_status"`
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

	return pkg.SuccessResponse(c, data)
}

func (h *DocumentCrHandler) GetAndika(c *fiber.Ctx) error {
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
	items, err := h.service.GetAndika(c.Context(), domainReq)
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

	return pkg.SuccessResponse(c, data)
}

func (h *DocumentCrHandler) GetOneMine(c *fiber.Ctx) error {
	req := new(docCrGetListRequest)

	if err := c.BodyParser(req); err != nil {
		return pkg.BadRequestResponse(c, "Failed to parse request body")
	}

	// Validasi NIK wajib diisi
	if req.RequesterNik == "" {
		return pkg.BadRequestResponse(c, "RequesterNik is required")
	}

	// Map request ke domain request
	domainReq := domainDocCr.GetListRequest{
		StartDate:    req.StartDate,
		EndDate:      req.EndDate,
		RequesterNik: req.RequesterNik,
		DocStatus:    req.DocStatus,
	}

	// Panggil service
	items, err := h.service.GetOneMine(c.Context(), domainReq)
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

	return pkg.SuccessResponse(c, data)
}
