package files

import (
	"context"
	"fmt"
)

type FilesClient struct {
}

func NewFilesClient() *FilesClient {
	return &FilesClient{}
}

func (c *FilesClient) GetUploadURL(ctx context.Context, fileName string, contentType string) (string, error) {
	uploadURL := fmt.Sprintf("http://localhost:8080/api/v1/files/upload/%s", fileName)
	return uploadURL, nil
}
