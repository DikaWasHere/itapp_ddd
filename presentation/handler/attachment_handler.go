package handler

import (
	"context"

	attachmentApp "itapp_ddd/application/attachment"
	"itapp_ddd/pkg"

	"github.com/gofiber/fiber/v2"
)

type AttachmentHandler struct {
	service *attachmentApp.Service
}

func NewAttachmentHandler(service *attachmentApp.Service) *AttachmentHandler {
	return &AttachmentHandler{service: service}
}

// UploadRequest request body untuk upload attachment
type UploadRequest struct {
	FormID      string   `json:"form_id"`
	Attachments []string `json:"doc_attachment"`
	Extensions  []string `json:"doc_attachment_extension"`
}

// Upload menangani POST /api/v1/attachment/upload
func (h *AttachmentHandler) Upload(c *fiber.Ctx) error {
	req := new(UploadRequest)
	if err := c.BodyParser(req); err != nil {
		return pkg.BadRequestResponse(c, "Failed to parse request body")
	}

	if req.FormID == "" {
		return pkg.BadRequestResponse(c, "form_id is required")
	}

	ctx := context.Background()
	result, err := h.service.ProcessAttachments(ctx, req.FormID, req.Attachments, req.Extensions)
	if err != nil {
		return pkg.BadRequestResponse(c, err.Error())
	}

	return pkg.SuccessResponse(c, fiber.Map{
		"doc_attachment": result,
	})
}
