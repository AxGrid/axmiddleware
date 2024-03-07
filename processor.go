package axmiddleware

import (
	"context"
	"github.com/rs/zerolog"
)

type MiddlewareProcessor[R, S any] interface {
	Process(ctx context.Context, request R, response S)
}

type middlewareProcessor[R, S any] struct {
	logger        zerolog.Logger
	ctx           context.Context
	handlersChain HandlersChain[R, S]
}

func (b *middlewareProcessor[R, S]) Process(ctx context.Context, request R, response S) {
	if ctx == nil {
		ctx = b.ctx
	}
	rctx := newContext(ctx, request, response)
	rctx.logger = &b.logger
	rctx.handlers = b.handlersChain
	rctx.Next()
}
