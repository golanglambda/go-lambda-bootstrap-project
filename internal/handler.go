package internal

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	cfg *Config
}

func NewHandler(cfg *Config) *Handler {
	return &Handler{cfg: cfg}
}

func (h *Handler) Execute(ctx context.Context, evt events.CloudWatchEvent) {
	fmt.Println(h.cfg)
	fmt.Println(evt)
	log.Info("Inside Handler - Info level message")
	log.Debug("Inside Handler - Debug level message")
}
