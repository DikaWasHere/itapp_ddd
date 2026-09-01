package handler

import (
	attachmentApp "itapp_ddd/application/attachment"
)

type AttachmentHandler struct {
	service *attachmentApp.Service
}

func NewAttachmentHandler(service *attachmentApp.Service) *AttachmentHandler {
	return &AttachmentHandler{service: service}
}
