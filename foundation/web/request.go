package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type validator interface {
	Validate() error
}

func Decode(r *http.Request, val any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(val)
	if err != nil {
		return fmt.Errorf("unable to decode payload: %w", err)
	}

	if v, ok := val.(validator); ok {
		err := v.Validate()
		if err != nil {
			return fmt.Errorf("unable to validate payload: %w", err)
		}
	}

	return nil
}
