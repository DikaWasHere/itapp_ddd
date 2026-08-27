package route

import (
	"itapp_ddd/presentation/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App, attachmentHandler *handler.AttachmentHandler, docCrHandler *handler.DocumentCrHandler) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Static file serving (sama seperti proyek asli)
	app.Static("/attachment", "./public/attachment")

	// Document CR
	v1.Post("/doc-cr", docCrHandler.GetAll)
}
