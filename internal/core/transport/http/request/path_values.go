package core_http_request

import (
	"fmt"
	"net/http"

	core_errors "github.com/Sklame132/rep/internal/core/errors"
)

func GetStringPathValue(r *http.Request, key string) (string, error) {
	pathValue := r.PathValue(key)
	if pathValue == "" {
		return "", fmt.Errorf(
			"no key='%s' in path values: %w",
			key,
			core_errors.ErrInvalidArgument,
		)
	}

	return pathValue, nil
}
