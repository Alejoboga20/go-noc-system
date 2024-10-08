package infrastructure

import (
	"errors"
	"net/http"

	"github.com/Alejoboga20/go-noc-system/domain/services"
)

type HTTPClientImplementation struct{}

func NewHTTPClientImplementation() services.HTTPClient {
	return &HTTPClientImplementation{}
}

func (client *HTTPClientImplementation) GET(url string) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errors.New(("Error: " + resp.Status + " - " + url))
	}

	return nil
}
