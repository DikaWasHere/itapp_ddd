package handler
 
import (
	applicationApp "itapp_ddd/application/application"
	domainApp "itapp_ddd/domain/application"
	"itapp_ddd/pkg"
 
	"github.com/gofiber/fiber/v2"
	"time"
)
 
type ApplicationHandler struct {
	service *applicationApp.Service
	logger  applicationApp.Logger
}
 
func NewApplicationHandler(service *applicationApp.Service, logger applicationApp.Logger) *ApplicationHandler {
	return &ApplicationHandler{service: service, logger: logger}
}
 
// Request DTO untuk GetAll
type applicationGetListRequest struct {
	Username     string `json:"username"`
	SessionToken string `json:"session_token"`
}
 
// Response item untuk daftar aplikasi
type applicationItem struct {
	ApplicationID   uint  			`json:"application_id"`
	ApplicationName string 			`json:"application_name"`
	CreatedAt       time.Time      `json:"created_at"`
    UpdatedAt       time.Time      `json:"updated_at"`
    DeletedAt       *time.Time     `json:"-" gorm:"index"`
}
 
// GetAll mengambil daftar aplikasi dari API SSO eksternal
func (h *ApplicationHandler) GetAll(c *fiber.Ctx) error {
	req := new(applicationGetListRequest)
	if err := c.BodyParser(req); err != nil {
		h.logger.Error("Failed to parse request body", err)
		return pkg.BadRequestResponse(c, "Failed to parse request body")
	}
 
	domainReq := domainApp.GetListApplicationRequest{
		Username:     req.Username,
		SessionToken: req.SessionToken,
	}
 
	apps, err := h.service.GetAll(c.Context(), domainReq)
	if err != nil {
		h.logger.Error("Failed to get applications", err)
		return pkg.BadRequestResponse(c, "Failed to get applications: "+err.Error())
	}
 
	var items []applicationItem
	for _, a := range apps {
		items = append(items, applicationItem{
			ApplicationID:   a.ApplicationID,
			ApplicationName: a.ApplicationName,
			CreatedAt:       a.CreatedAt,
			UpdatedAt:       a.UpdatedAt,
			DeletedAt:       a.DeletedAt,
		})
	}
 
	h.logger.Info("[%s] %s | Status: 200 | Fetched %d applications", c.Method(), c.Path(), len(items))
	return pkg.SuccessResponse(c, items)
}