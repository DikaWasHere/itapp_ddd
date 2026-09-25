package attachment

import (
	"context"
	"encoding/base64"
	"fmt"

	domainAttachment "itapp_ddd/domain/attachment"
)

// FileStorer adalah interface untuk menyimpan file ke disk
type FileStorer interface {
	Save(base64Data string, filePath string) error
}

// Service adalah application service untuk attachment
type Service struct {
	repo      domainAttachment.AttachmentRepository
	fileStore FileStorer
}

func NewService(repo domainAttachment.AttachmentRepository, fileStore FileStorer) *Service {
	return &Service{repo: repo, fileStore: fileStore}
}

func (s *Service) ProcessAttachments(ctx context.Context, formID string, base64DataList []string, extensions []string,startNumber int) (string, error) {
	if len(base64DataList) == 0 {
		return "", nil
	}

	var fileNames []string

	for i, b64 := range base64DataList {
		// Validasi base64
		if _, err := base64.StdEncoding.DecodeString(b64); err != nil {
			return "", fmt.Errorf("invalid base64 data for attachment %d", i+1)
		}

		ext := ""
		if i < len(extensions) {
			ext = extensions[i]
		}
		fileName := fmt.Sprintf("%s%d%s", formID, startNumber+i+1, ext)
		filePath := fmt.Sprintf("public/attachment/%s", fileName)

		// Simpan file ke disk
		if err := s.fileStore.Save(b64, filePath); err != nil {
			return "", fmt.Errorf("failed to save attachment %d: %w", i+1, err)
		}

		// Simpan ke database
		att := domainAttachment.NewAttachment(
			domainAttachment.FileName(fileName),
			domainAttachment.FilePath(filePath),
			ext,
		)
		if err := s.repo.Save(ctx, att); err != nil {
			return "", fmt.Errorf("failed to save attachment record: %w", err)
		}

		fileNames = append(fileNames, fileName)
	}

	result := ""
	for i, name := range fileNames {
		if i > 0 {
			result += "|"
		}
		result += name
	}

	return result, nil
}
