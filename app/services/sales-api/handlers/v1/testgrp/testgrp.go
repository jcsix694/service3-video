package testgrp

import (
	"context"
	// "errors"
	// "github.com/jcsix694/service3-video/business/sys/validate"
	"github.com/jcsix694/service3-video/foundation/web"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
)

// Handlers manages the set of check endpoints
type Handlers struct {
	Log *zap.SugaredLogger
}

// Test handle for development
func (h Handlers) Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
		// return errors.New("untrusted error")
		// return validate.NewRequestError(errors.New("trusted error"), http.StatusBadRequest)
		return web.NewShutdownError("Restart Service")
	}

	status := struct {
		Status string
	}{
		Status: "OK",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
