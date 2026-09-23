package container

import (
	attachmentApp "itapp_ddd/application/attachment"
	docCrApp "itapp_ddd/application/doc_cr"
	aplicationApp "itapp_ddd/application/application"
	infraLogger "itapp_ddd/infrastructure/logger"
	infraSSO "itapp_ddd/infrastructure/sso"
	notificationApp "itapp_ddd/application/notification"
	infraNotification "itapp_ddd/infrastructure/notification"
	"itapp_ddd/infrastructure/persistence"
	"itapp_ddd/infrastructure/storage"
	"itapp_ddd/presentation/handler"
	"itapp_ddd/migrations"
)

type Container struct {
	AttachmentRepo    *persistence.AttachmentRepositoryDB
	DocCrRepo         *persistence.DocumentCrRepositoryDB
	FileStorage       *storage.FileStorage
	AttachmentService *attachmentApp.Service
	DocCrService      *docCrApp.Service
	AttachmentHandler *handler.AttachmentHandler
	DocCrHandler      *handler.DocumentCrHandler
	NotificationService *notificationApp.Service
	ApplicationService *aplicationApp.Service
	ApplicationHandler *handler.ApplicationHandler
}

func NewContainer() *Container {
	// 1. Init Database
	persistence.DatabaseInit()
	migration.RunMigration()

	// 2. Init Infrastructure
	attachmentRepo := persistence.NewAttachmentRepositoryDB(persistence.DB)
	docCrRepo := persistence.NewDocumentCrRepositoryDB(persistence.DB)
	fileStorage := storage.NewFileStorage()

	// 3. Init Application Service
	attachmentService := attachmentApp.NewService(attachmentRepo, fileStorage)
	docCrService := docCrApp.NewService(docCrRepo, &infraLogger.LoggerImpl{}, fileStorage, attachmentService)
	notificationService := notificationApp.NewService(infraNotification.NewHTTPNotifier(), &infraLogger.LoggerImpl{})

	// 4. Init Handler
	attachmentHandler := handler.NewAttachmentHandler(attachmentService)
	docCrHandler := handler.NewDocumentCrHandler(docCrService, &infraLogger.LoggerImpl{}, notificationService)

	ssoClient := infraSSO.NewClient()
	ApplicationService := aplicationApp.NewService(ssoClient, &infraLogger.LoggerImpl{})
	ApplicationHandler := handler.NewApplicationHandler(ApplicationService, &infraLogger.LoggerImpl{})

	return &Container{
		AttachmentRepo:    attachmentRepo,
		DocCrRepo:         docCrRepo,
		FileStorage:       fileStorage,
		AttachmentService: attachmentService,
		DocCrService:      docCrService,
		AttachmentHandler: attachmentHandler,
		DocCrHandler:      docCrHandler,
		NotificationService: notificationService,
		ApplicationService: ApplicationService,
		ApplicationHandler: ApplicationHandler,
	}
}
