package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"io"
)

type MinioStorage struct {
	client *minio.Client
	bucket string
}

func NewMinioStorage(client *minio.Client, bucket string) *MinioStorage {
	return &MinioStorage{
		client: client,
		bucket: bucket,
	}
}

func (s *MinioStorage) UploadFile(ctx context.Context, file io.Reader, contentType string, filename string, filesize int64) error {
	_, err := s.client.PutObject(ctx, s.bucket, filename, file, filesize, minio.PutObjectOptions{
		ContentType: contentType,
	})
	return err
}

func (s *MinioStorage) GetFileURL(ctx context.Context, filename string) (string, error) {
	presignedURL := fmt.Sprintf("/files/app/%s", filename)
	return presignedURL, nil
}
