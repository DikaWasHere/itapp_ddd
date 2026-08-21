package main

import (
	"fmt"
	"log"
	"os"

	"itapp_ddd/infrastructure/persistence"
	"itapp_ddd/infrastructure/storage"
	"itapp_ddd/presentation/handler"
	"itapp_ddd/presentation/route"

	attachmentApp "itapp_ddd/application/attachment"
	docCrApp "itapp_ddd/application/doc_cr"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// 1. Init Database
	persistence.DatabaseInit()

	// 2. Init Infrastructure
	attachmentRepo := persistence.NewAttachmentRepositoryDB(persistence.DB)
	docCrRepo := persistence.NewDocumentCrRepositoryDB(persistence.DB)
	fileStorage := storage.NewFileStorage()

	// 3. Init Application Service
	attachmentService := attachmentApp.NewService(attachmentRepo, fileStorage)
	docCrService := docCrApp.NewService(docCrRepo)

	// 4. Init Handler
	attachmentHandler := handler.NewAttachmentHandler(attachmentService)
	docCrHandler := handler.NewDocumentCrHandler(docCrService)

	// 5. Init Fiber
	app := fiber.New()
	app.Use(cors.New())

	// 6. Init Routes
	route.RouteInit(app, attachmentHandler, docCrHandler)

	// 7. Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8280"
	}

	fmt.Printf("Server running on port %s\n", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
