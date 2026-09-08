package container

import (
	attachmentApp "itapp_ddd/application/attachment"
	docCrApp "itapp_ddd/application/doc_cr"
	infraLogger "itapp_ddd/infrastructure/logger"
	"itapp_ddd/infrastructure/persistence"
	"itapp_ddd/infrastructure/storage"
	"itapp_ddd/presentation/handler"
)

type Container struct {
	AttachmentRepo    *persistence.AttachmentRepositoryDB
	DocCrRepo         *persistence.DocumentCrRepositoryDB
	FileStorage       *storage.FileStorage
	AttachmentService *attachmentApp.Service
	DocCrService      *docCrApp.Service
	AttachmentHandler *handler.AttachmentHandler
	DocCrHandler      *handler.DocumentCrHandler
}

func NewContainer() *Container {
	// 1. Init Database
	persistence.DatabaseInit()

	// 2. Init Infrastructure
	attachmentRepo := persistence.NewAttachmentRepositoryDB(persistence.DB)
	docCrRepo := persistence.NewDocumentCrRepositoryDB(persistence.DB)
	fileStorage := storage.NewFileStorage()

	// 3. Init Application Service
	attachmentService := attachmentApp.NewService(attachmentRepo, fileStorage)
	docCrService := docCrApp.NewService(docCrRepo, &infraLogger.LoggerImpl{}, fileStorage)

	// 4. Init Handler
	attachmentHandler := handler.NewAttachmentHandler(attachmentService)
	docCrHandler := handler.NewDocumentCrHandler(docCrService, &infraLogger.LoggerImpl{})

	return &Container{
		AttachmentRepo:    attachmentRepo,
		DocCrRepo:         docCrRepo,
		FileStorage:       fileStorage,
		AttachmentService: attachmentService,
		DocCrService:      docCrService,
		AttachmentHandler: attachmentHandler,
		DocCrHandler:      docCrHandler,
	}
}
