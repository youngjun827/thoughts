// Test Handler. Will be removed in production.
package hackgrp

import (
	"context"
	"encoding/json"
	"net/http"
)

func Hack(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
    status := struct {
        Status string
    }{
        Status: "OK",
    }

    err := json.NewEncoder(w).Encode(status)
    if err != nil {
        return err
    }

    return nil
}