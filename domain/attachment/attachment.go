package attachment

import "errors"

// Value Object: FileName
type FileName string

func NewFileName(name string) (FileName, error) {
	if name == "" {
		return "", errors.New("file name cannot be empty")
	}
	return FileName(name), nil
}

// Value Object: FilePath
type FilePath string

func NewFilePath(path string) (FilePath, error) {
	if path == "" {
		return "", errors.New("file path cannot be empty")
	}
	return FilePath(path), nil
}

// Entity: Attachment
type Attachment struct {
	FileName  FileName
	FilePath  FilePath
	Extension string
}

// NewAttachment constructor
func NewAttachment(fileName FileName, filePath FilePath, extension string) *Attachment {
	return &Attachment{
		FileName:  fileName,
		FilePath:  filePath,
		Extension: extension,
	}
}
