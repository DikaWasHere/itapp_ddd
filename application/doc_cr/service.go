package doc_cr

import (
	"context"
	"fmt"
	"strconv"
	"time"
	domainDocCr "itapp_ddd/domain/doc_cr"
	"itapp_ddd/pkg"
)

// Service adalah application service untuk DocumentCr
type Service struct {
	repo      domainDocCr.DocumentCrRepository
	logger    Logger
	fileStore FileStorer
}

func NewService(repo domainDocCr.DocumentCrRepository, logger Logger, fileStore FileStorer) *Service {
	return &Service{repo: repo, logger: logger, fileStore: fileStore}
}

// FileStorer adalah interface untuk menyimpan file ke disk
type FileStorer interface {
	Save(base64Data string, filePath string) error
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
            // Decode base64 dan simpan lewat FileStorage (DRY — pakai utilitas yang sama dengan attachment service)
            ext := ""
            if index < len(req.DocAttachmentExtension) {
                ext = req.DocAttachmentExtension[index]
            }
            fileName := fmt.Sprintf("%s_%d%s", FormId, index+1, ext)
            filePath := fmt.Sprintf("public/attachment/%s", fileName)

            if err := s.fileStore.Save(item, filePath); err != nil {
                s.logger.Error("Failed to save attachment", err)
                return nil, fmt.Errorf("failed to save attachment: %w", err)
            }

            imageFileNames = append(imageFileNames, fileName)
        }
    }
 
    // Update request dengan FormId dan nama file attachment
    req.FormId = FormId
    req.DocAttachment = imageFileNames
 
    result, err := s.repo.Create(ctx, req)
    if err != nil {
        s.logger.Error("Create failed", err)
        return nil, err
    }
 
    s.logger.Info("DocumentCr created successfully with ID: %d", result.DocCrID)
    return result, nil
}

func (s *Service) SendReview(ctx context.Context, req domainDocCr.SendReviewRequest) (*domainDocCr.SendReviewResult, error) {
	if req.DocCrID == 0 {
		s.logger.Warn("SendReview failed: DocCrID is required")
		return nil, fmt.Errorf("DocCrID is required")
	}
 
	result, err := s.repo.SendReview(ctx, req)
	if err != nil {
		s.logger.Error("SendReview failed", err)
		return nil, err
	}
 
	s.logger.Info("Review submitted for DocCrID: %d", result.DocCrID)
	return result, nil
}

// SendReviewReject menolak CR dari tahap review
func (s *Service) SendReviewReject(ctx context.Context, req domainDocCr.SendReviewRejectRequest) (*domainDocCr.SendReviewRejectResult, error) {
	if req.DocCrID == 0 {
		s.logger.Warn("SendReviewReject failed: DocCrID is required")
		return nil, fmt.Errorf("DocCrID is required")
	}
 
	result, err := s.repo.SendReviewReject(ctx, req)
	if err != nil {
		s.logger.Error("SendReviewReject failed", err)
		return nil, err
	}
 
	s.logger.Info("Review rejected for DocCrID: %d", result.DocCrID)
	return result, nil
}

// SendRevision mengirim revisi CR (mengubah status ke "Waiting Review")
func (s *Service) SendRevision(ctx context.Context, req domainDocCr.SendRevisionRequest) (*domainDocCr.SendRevisionResult, error) {
	if req.DocCrID == 0 {
		s.logger.Warn("SendRevision failed: DocCrID is required")
		return nil, fmt.Errorf("DocCrID is required")
	}
 
	// Handle base64 attachment jika ada
	if req.DocAttachment != "" {
		imageFileName := fmt.Sprintf("%s.png", req.FormId)
		imageFilePath := fmt.Sprintf("public/attachment/%s", imageFileName)
 
		if err := s.fileStore.Save(req.DocAttachment, imageFilePath); err != nil {
			s.logger.Error("Failed to save attachment", err)
			return nil, fmt.Errorf("failed to save attachment: %w", err)
		}
		req.DocAttachment = imageFileName
	}
 
	result, err := s.repo.SendRevision(ctx, req)
	if err != nil {
		s.logger.Error("SendRevision failed", err)
		return nil, err
	}
 
	s.logger.Info("Revision submitted for DocCrID: %d", result.DocCrID)
	return result, nil
}