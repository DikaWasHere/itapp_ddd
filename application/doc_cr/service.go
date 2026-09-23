package doc_cr

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"strings"
	"encoding/base64"
	domainDocCr "itapp_ddd/domain/doc_cr"
	
)

type Service struct {
	repo      domainDocCr.DocumentCrRepository
	logger    Logger
	fileStore FileStorer
	attachment AttachmentProcessor
}

func NewService(repo domainDocCr.DocumentCrRepository, logger Logger, fileStore FileStorer,attachment AttachmentProcessor) *Service {
	return &Service{repo: repo, logger: logger, fileStore: fileStore,attachment: attachment}
}

type FileStorer interface {
	Save(base64Data string, filePath string) error
}

type AttachmentProcessor interface {
	ProcessAttachments(ctx context.Context, formID string, base64DataList []string, extensions []string,startNumber int) (string, error)
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


func (s *Service) Create(ctx context.Context, req domainDocCr.CreateDocCrRequest) (*domainDocCr.DocumentCr, error) {
	formId := req.FormId
	existing := make([]string, 0) // filename/link lama yang dipertahankan
 
	if req.DocCrID == 0 {
		// ===== MODE CREATE (kode existing Anda, tidak berubah) =====
		nowYear := time.Now().Format("2006")
		nowYearId := time.Now().Format("06")
 
		count, err := s.repo.CountByYear(ctx, nowYear)
		if err != nil {
			return nil, err
		}
 
		formId = "IT.CR." + nowYearId + "." + padLeft(strconv.Itoa(int(count+1)), 4)
		if req.FormId != "" {
			formId = req.FormId
		}
 
		req.FormId = formId
		req.DocStatus = "Waiting Submit"
	} else {
		for _, att := range req.DocAttachment {
			if !isBase64(att) {
				existing = append(existing, att)
			}
		}
	}
 
	if len(req.DocAttachment) > 0 {
		newB64 := make([]string, 0)
		newExt := make([]string, 0)
		for i, att := range req.DocAttachment {
			if isBase64(att) {
				ext := ""
				if i < len(req.DocAttachmentExtension) {
					ext = req.DocAttachmentExtension[i]
				}
				newB64 = append(newB64, att)
				newExt = append(newExt, ext)
			}
		}
 
		if len(newB64) > 0 {
			joined, err := s.attachment.ProcessAttachments(ctx, formId, newB64, newExt, len(existing))
			if err != nil {
				s.logger.Error("Failed to process attachments", err)
				return nil, err
			}
			existing = append(existing, strings.Split(joined, "|")...)
		}
 
		req.DocAttachment = existing
	}
 
	return s.repo.Upsert(ctx, req)
}
 
// padLeft menambahkan leading zeros
func padLeft(str string, length int) string {
	for len(str) < length {
		str = "0" + str
	}
	return str
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

// SendApprove mengirim approve CR
func (s *Service) SendApprove(ctx context.Context, req domainDocCr.SendApproveRequest) (*domainDocCr.SendApproveResult, error) {
	if req.ApprovalID == 0 {
		s.logger.Warn("SendApprove failed: ApprovalID is required")
		return nil, fmt.Errorf("ApprovalID is required")
	}
 
	if req.DocCrID == 0 {
		s.logger.Warn("SendApprove failed: DocCrID is required")
		return nil, fmt.Errorf("DocCrID is required")
	}
 
	result, err := s.repo.SendApprove(ctx, req)
	if err != nil {
		s.logger.Error("SendApprove failed", err)
		return nil, err
	}
 
	s.logger.Info("Approval submitted for DocCrID: %d", req.DocCrID)
	return result, nil
}

// SendApproveReject menolak hasil approval CR
func (s *Service) SendApproveReject(ctx context.Context, req domainDocCr.SendApproveRejectRequest) (*domainDocCr.SendApproveRejectResult, error) {
	if req.DocCrID == 0 {
		s.logger.Warn("SendApproveReject failed: DocCrID is required")
		return nil, fmt.Errorf("DocCrID is required")
	}
 
	result, err := s.repo.SendApproveReject(ctx, req)
	if err != nil {
		s.logger.Error("SendApproveReject failed", err)
		return nil, err
	}
 
	s.logger.Info("Approval rejected for DocCrID: %d", result.DocCrID)
	return result, nil
}

func (s *Service) UpdateStatus(ctx context.Context, req domainDocCr.UpdateStatusRequest) (*domainDocCr.UpdateStatusResult, error) {
	if req.DocCrID == 0 {
		s.logger.Warn("UpdateStatus failed: DocCrID is required")
		return nil, fmt.Errorf("DocCrID is required")
	}
 
	if req.DocStatus == "" {
		s.logger.Warn("UpdateStatus failed: DocStatus is required")
		return nil, fmt.Errorf("DocStatus is required")
	}
 
	result, err := s.repo.UpdateStatus(ctx, req)
	if err != nil {
		s.logger.Error("UpdateStatus failed", err)
		return nil, err
	}
 
	s.logger.Info("Document status updated to '%s' for DocCrID: %d", result.DocStatus, result.DocCrID)
	return result, nil
}

// Delete menghapus dokumen CR
func (s *Service) Delete(ctx context.Context, req domainDocCr.DeleteRequest) (*domainDocCr.DeleteResult, error) {
	if req.DocCrID == 0 {
		s.logger.Warn("Delete failed: DocCrId is required")
		return nil, fmt.Errorf("DocCrID is required")
	}

	result, err := s.repo.Delete(ctx, req)
	if err != nil {
		s.logger.Error("Delete failed", err)
		return nil, err
	}

	s.logger.Info("document telah berhasil dihapus dengan ID: %d", result.DocCrID)
	return result,nil
}

func isBase64(s string) bool {
	if len(s) < 100 { // filename pendek; base64 file asli selalu panjang
		return false
	}
	_, err := base64.StdEncoding.DecodeString(s)
	return err == nil
}