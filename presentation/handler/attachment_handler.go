package handler

import (
	//"context"

	attachmentApp "itapp_ddd/application/attachment"
	//"itapp_ddd/pkg"
	//"github.com/gofiber/fiber/v2"
)

type AttachmentHandler struct {
	service *attachmentApp.Service
}

func NewAttachmentHandler(service *attachmentApp.Service) *AttachmentHandler {
	return &AttachmentHandler{service: service}
}
