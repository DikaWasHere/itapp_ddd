package persistence

import (
	"context"

	domainAttachment "itapp_ddd/domain/attachment"

	"gorm.io/gorm"
)

// AttachmentModel adalah model GORM untuk tabel attachment
type AttachmentModel struct {
	ID        uint   `gorm:"primarykey"`
	FileName  string `gorm:"column:file_name"`
	FilePath  string `gorm:"column:file_path"`
	Extension string `gorm:"column:extension"`
}

func (AttachmentModel) TableName() string {
	return "mst.mst_attachment"
}

// toEntity konversi model ke domain entity
func (m *AttachmentModel) toEntity() *domainAttachment.Attachment {
	return &domainAttachment.Attachment{
		FileName:  domainAttachment.FileName(m.FileName),
		FilePath:  domainAttachment.FilePath(m.FilePath),
		Extension: m.Extension,
	}
}

// fromEntity konversi domain entity ke model
func fromEntity(a *domainAttachment.Attachment) *AttachmentModel {
	return &AttachmentModel{
		FileName:  string(a.FileName),
		FilePath:  string(a.FilePath),
		Extension: a.Extension,
	}
}

// AttachmentRepositoryDB implementasi konkret dari AttachmentRepository
type AttachmentRepositoryDB struct {
	db *gorm.DB
}

func NewAttachmentRepositoryDB(db *gorm.DB) *AttachmentRepositoryDB {
	return &AttachmentRepositoryDB{db: db}
}

func (r *AttachmentRepositoryDB) Save(ctx context.Context, att *domainAttachment.Attachment) error {
	model := fromEntity(att)
	return r.db.WithContext(ctx).Create(model).Error
}

func (r *AttachmentRepositoryDB) FindByName(ctx context.Context, fileName string) (*domainAttachment.Attachment, error) {
	var model AttachmentModel
	if err := r.db.WithContext(ctx).Where("file_name = ?", fileName).First(&model).Error; err != nil {
		return nil, err
	}
	return model.toEntity(), nil
}

func (r *AttachmentRepositoryDB) DeleteByName(ctx context.Context, fileName string) error {
	return r.db.WithContext(ctx).Where("file_name = ?", fileName).Delete(&AttachmentModel{}).Error
}
