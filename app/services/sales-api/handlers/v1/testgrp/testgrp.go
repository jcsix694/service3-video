package testgrp

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

// Handlers manages the set of check endpoints
type Handlers struct {
	Log *zap.SugaredLogger
}

// Test handle for development
func (h Handlers) Test(w http.ResponseWriter, r *http.Request) {
	status := struct {
		Status string
	}{
		Status: "OK",
	}
	json.NewEncoder(w).Encode(status)

	statusCode := http.StatusOK

	h.Log.Infow("readiness", "statusCode", statusCode, "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)

}
