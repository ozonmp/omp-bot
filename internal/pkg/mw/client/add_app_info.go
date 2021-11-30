package mwclient

import (
	"context"

	"github.com/ozonmp/omp-bot/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	appNameHeader    = "x-app-name"
	appVersionHeader = "x-app-version"
)

// AddAppInfoUnary добавляет в единичные запросы информацию о клиенте.
func AddAppInfoUnary(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

	cfg := config.GetConfigInstance()

	ctx = metadata.AppendToOutgoingContext(ctx, appNameHeader, cfg.Project.Name)
	ctx = metadata.AppendToOutgoingContext(ctx, appVersionHeader, cfg.Project.Version)
	return invoker(ctx, method, req, reply, cc, opts...)
}
