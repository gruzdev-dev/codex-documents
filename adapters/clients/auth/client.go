package auth

import (
	"context"
	"github.com/gruzdev-dev/codex-auth/proto"
	"github.com/gruzdev-dev/codex-documents/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type client struct {
	grpc   proto.TmpAccessClient
	secret string
}

func NewClient(conn grpc.ClientConnInterface, secret string) *client {
	return &client{
		grpc:   proto.NewTmpAccessClient(conn),
		secret: secret,
	}
}

func (c *client) GenerateTmpToken(ctx context.Context, data domain.GenerateTmpTokenRequest) (*domain.GenerateTmpTokenResponse, error) {
	md := metadata.New(map[string]string{
		"x-internal-token": c.secret,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	req := &proto.GenerateTmpTokenRequest{
		Payload:    data.Payload,
		TtlSeconds: data.TtlSeconds,
	}

	resp, err := c.grpc.GenerateTmpToken(ctx, req)
	if err != nil {
		return nil, err
	}

	return &domain.GenerateTmpTokenResponse{
		TmpToken: resp.GetToken(),
	}, nil
}
