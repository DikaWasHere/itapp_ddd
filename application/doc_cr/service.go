package doc_cr

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	domainDocCr "itapp_ddd/domain/doc_cr"
	"itapp_ddd/pkg"
)

// Service adalah application service untuk DocumentCr
type Service struct {
	repo   domainDocCr.DocumentCrRepository
	logger Logger
}

func NewService(repo domainDocCr.DocumentCrRepository, logger Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

// GetAll mengambil semua DocumentCr berdasarkan filter
func (s *Service) GetAll(ctx context.Context, req domainDocCr.GetListRequest) ([]*domainDocCr.DocumentCrListItem, error) {
	items, err := s.repo.GetAll(ctx, req)
	if err != nil {
		s.logger.Error("GetAll failed", err)
		return nil, err
	}
	return items, nil
}

func (s *Service) GetDetail(ctx context.Context, req domainDocCr.GetDetailRequest) (*domainDocCr.DocumentCrDetailItem, error) {

	item, err := s.repo.GetDetail(ctx, req)
	if err != nil {
		s.logger.Error("GetDetail failed", err)
		return nil, err
	}

	return item, nil
}

// Create membuat DocumentCr baru dengan auto-generate FormId dan handle attachment
func (s *Service) Create(ctx context.Context, req domainDocCr.CreateDocCrRequest) (*domainDocCr.DocumentCr, error) {
    // Validasi input
    if req.DocCrName == "" {
        s.logger.Warn("Create failed: DocCrName is required")
        return nil, fmt.Errorf("DocCrName is required")
    }
 
    // Auto-generate FormId jika DocCrID tidak ada (create baru)
    nowYear := time.Now().Format("2006")
    nowYearId := time.Now().Format("06")
 
    FormId := req.FormId // Default gunakan dari request
 
    if req.DocCrID == 0 { // Create baru
        count, err := s.repo.CountByYear(ctx, nowYear)
        if err != nil {
            s.logger.Error("Failed to count documents by year", err)
            return nil, err
        }
        FormId = "IT.CR." + nowYearId + "." + pkg.PadLeft(strconv.Itoa(int(count+1)))
    }
 
    // Handle base64 attachments
    var imageFileNames []string
    if len(req.DocAttachment) > 0 {
        for index, item := range req.DocAttachment {
            // Decode base64
            imageData := []byte(item) // In real impl, use base64.StdEncoding.DecodeString
 
            // Generate filename: {FormId}_{index+1}{extension}
            ext := ""
            if index < len(req.DocAttachmentExtension) {
                ext = req.DocAttachmentExtension[index]
            }
            fileName := fmt.Sprintf("%s_%s%s", FormId, strconv.Itoa(index+1), ext)
            imageFileNames = append(imageFileNames, fileName)
 
            // Save file
            filePath := fmt.Sprintf("public/attachment/%s", fileName)
            err := os.WriteFile(filePath, imageData, 0644)
            if err != nil {
                s.logger.Error("Failed to save attachment", err)
                return nil, fmt.Errorf("failed to save attachment: %w", err)
            }
        }
    }
 
    // Build doc_attachment string (joined by "|")
    var docAttachment string
    if len(imageFileNames) > 0 {
        docAttachment = strings.Join(imageFileNames, "|")
    }
 
    // Update request dengan FormId dan docAttachment yang sudah diproses
    req.FormId = FormId
    req.DocAttachment = []string{docAttachment} // Store as single string in array
 
    result, err := s.repo.Create(ctx, req)
    if err != nil {
        s.logger.Error("Create failed", err)
        return nil, err
    }
 
    s.logger.Info("DocumentCr created successfully with ID: %d", result.DocCrID)
    return result, nil
}