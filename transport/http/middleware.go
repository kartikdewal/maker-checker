package http

import (
	"context"
	"maker-checker/logger"
	"time"
)

type Middleware func(ApiHandler) ApiHandler

func LoggingMiddleware(log logger.ContextLogger) Middleware {
	return func(next ApiHandler) ApiHandler {
		return &loggingMiddleware{
			next:   next,
			logger: log,
		}
	}
}

type loggingMiddleware struct {
	next   ApiHandler
	logger logger.ContextLogger
}

func (mw loggingMiddleware) GetProfile(ctx context.Context, id string) (p Profile, err error) {
	defer func(begin time.Time) {
		mw.logger.Infow(ctx, "", "method", "GetProfile", "id", id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.GetProfile(ctx, id)
}

func (mw loggingMiddleware) PostProfile(ctx context.Context, p Profile) (id string, err error) {
	defer func(begin time.Time) {
		mw.logger.Infow(ctx, "", "method", "PostProfile", "id", id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.PostProfile(ctx, p)
}

func (mw loggingMiddleware) GetDocument(ctx context.Context, id string) (d Document, err error) {
	defer func(begin time.Time) {
		mw.logger.Infow(ctx, "", "method", "GetDocument", "id", id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.GetDocument(ctx, id)
}

func (mw loggingMiddleware) PostDocument(ctx context.Context, d Document) (id string, err error) {
	defer func(begin time.Time) {
		mw.logger.Infow(ctx, "", "method", "PostDocument", "id", id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.PostDocument(ctx, d)
}

func (mw loggingMiddleware) PostDocumentRequest(ctx context.Context, r DocumentRequest) (id string, err error) {
	defer func(begin time.Time) {
		mw.logger.Infow(ctx, "", "method", "PostDocumentRequest", "id", id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.PostDocumentRequest(ctx, r)
}

func (mw loggingMiddleware) PutDocumentRequest(ctx context.Context, reqID string, r DocumentRequest) (id string, err error) {
	defer func(begin time.Time) {
		mw.logger.Infow(ctx, "", "method", "PutDocumentRequest", "id", id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.PutDocumentRequest(ctx, reqID, r)
}

func (mw loggingMiddleware) GetDocumentRequest(ctx context.Context, id string) (d DocumentRequest, err error) {
	defer func(begin time.Time) {
		mw.logger.Infow(ctx, "", "method", "GetDocumentRequest", "id", d.ID, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.GetDocumentRequest(ctx, id)
}
