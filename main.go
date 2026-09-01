package main
 
import (
	"fmt"
	"log"
	"os"
 
	"itapp_ddd/container"
	"itapp_ddd/infrastructure/logger"
	"itapp_ddd/presentation/route"
 
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)
 
func main() {
	logger.SetupLogger()
	logger.Info("Logger initialized")

	// 1. Init Container
	c := container.NewContainer()
 
	// 2. Init Fiber
	app := fiber.New()
	app.Use(cors.New())
 
	// 3. Init Routes
	route.RouteInit(app, c.AttachmentHandler, c.DocCrHandler)
 
	// 4. Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8280"
	}
 
	fmt.Printf("Server running on port %s\n", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

