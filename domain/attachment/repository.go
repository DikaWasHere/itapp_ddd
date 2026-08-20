package attachment

import "context"

// AttachmentRepository adalah kontrak untuk akses data
type AttachmentRepository interface {
	Save(ctx context.Context, att *Attachment) error
	FindByName(ctx context.Context, fileName string) (*Attachment, error)
	DeleteByName(ctx context.Context, fileName string) error
}
