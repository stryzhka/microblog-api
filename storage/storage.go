package storage

import (
	"context"
	"io"
)

type FileData struct {
	File        io.Reader
	Size        int64
	ContentType string
}

type FileStorage interface {
	UploadFile(ctx context.Context, file io.Reader, contentType, filename string, filesize int64) error
	GetFileURL(ctx context.Context, filename string) (string, error)
}
