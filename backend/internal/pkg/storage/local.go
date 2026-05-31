package storage

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

const uploadDir = "/app/uploads"

// EnsureUploadDir creates the upload directory if it doesn't exist
func EnsureUploadDir() error {
	dirs := []string{
		filepath.Join(uploadDir, "ijazah"),
		filepath.Join(uploadDir, "transcript"),
		filepath.Join(uploadDir, "stempel"),
	}
	for _, d := range dirs {
		if err := os.MkdirAll(d, 0755); err != nil {
			return fmt.Errorf("create dir %s: %v", d, err)
		}
	}
	return nil
}

// UploadDocument saves a file locally and returns the relative path
func UploadDocument(name string, file multipart.File) error {
	dst := filepath.Join(uploadDir, name)

	// Ensure parent directory exists
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return fmt.Errorf("create parent dir: %v", err)
	}

	out, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("create file: %v", err)
	}
	defer out.Close()

	if _, err := io.Copy(out, file); err != nil {
		return fmt.Errorf("write file: %v", err)
	}

	return nil
}

// GetDocumentPath returns the full filesystem path for a document
func GetDocumentPath(name string) string {
	return filepath.Join(uploadDir, name)
}

// GetDocumentURL returns the public URL for a document
func GetDocumentURL(name string) string {
	return "/uploads/" + name
}
