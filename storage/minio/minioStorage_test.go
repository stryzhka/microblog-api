package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestMinioStorage_Int(t *testing.T) {
	ctx := context.Background()
	endpoint := "localhost:9000"
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4("minioadmin", "minioadmin", ""),
		Secure: false,
	})
	if err != nil {
		t.Fatalf("failed to create minio client: %v", err)
	}
	storage := NewMinioStorage(client, "test-bucket")
	t.Run("Upload and Get URL", func(t *testing.T) {
		fileContent := strings.NewReader("test")
		filename := "testFile"
		err = storage.UploadFile(ctx, fileContent, "text/plain", filename, 4)
		assert.NoError(t, err)
		url, err := storage.GetFileURL(ctx, filename)
		assert.NoError(t, err)
		assert.Contains(t, url, filename)
	})
}

func TestMinioStorage_Handler(t *testing.T) {

}
