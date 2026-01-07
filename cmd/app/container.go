package main

import (
	"codex-documents/pkg/container"
	httpServer "codex-documents/servers/http"

	"go.uber.org/dig"
)

func BuildContainer() (*dig.Container, error) {
	c, err := container.BuildAppContainer()
	if err != nil {
		return nil, err
	}

	if err := c.Provide(httpServer.NewServer); err != nil {
		return nil, err
	}

	return c, nil
}
