package files

import (
	"context"
	"github.com/gruzdev-dev/codex-documents/core/domain"
	"github.com/gruzdev-dev/codex-documents/core/ports"
	"github.com/gruzdev-dev/codex-files/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type client struct {
	grpc   proto.FilesServiceClient
	secret string
}

func NewClient(conn grpc.ClientConnInterface, secret string) ports.FileProvider {
	return &client{
		grpc:   proto.NewFilesServiceClient(conn),
		secret: secret,
	}
}

func (c *client) GetPresignedUrls(ctx context.Context, data domain.GetPresignedUrlsRequest) (*domain.PresignedUrlsResponse, error) {
	md := metadata.New(map[string]string{
		"x-internal-token": c.secret,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	req := &proto.GeneratePresignedUrlsRequest{
		UserId:      data.UserId,
		ContentType: data.ContentType,
		Size:        data.Size,
	}

	resp, err := c.grpc.GeneratePresignedUrls(ctx, req)
	if err != nil {
		return nil, err
	}

	return &domain.PresignedUrlsResponse{
		FileId:      resp.FileId,
		UploadUrl:   resp.UploadUrl,
		DownloadUrl: resp.DownloadUrl,
	}, nil
}
