package testgrp

import (
	"context"
	"github.com/jcsix694/service3-video/foundation/web"
	"go.uber.org/zap"
	"net/http"
)

// Handlers manages the set of check endpoints
type Handlers struct {
	Log *zap.SugaredLogger
}

// Test handle for development
func (h Handlers) Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{
		Status: "OK",
	}

	statusCode := http.StatusOK

	h.Log.Infow("readiness", "statusCode", statusCode, "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)

	return web.Respond(ctx, w, status, http.StatusOK)
}
