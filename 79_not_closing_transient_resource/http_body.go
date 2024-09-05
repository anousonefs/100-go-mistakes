package main

import (
	"io"
	"log"
	"net/http"
)

type handler struct {
	client http.Client
	url    string
}

/* If we close the body without a read, the default HTTP transport may close the connection. */
/* If we close the body following a read, the default HTTP transport won’t close the connection; hence, it may be reused. */

func (h handler) getStatusCode(body io.Reader) (int, error) {
	resp, err := h.client.Post(h.url, "application/json", body)
	if err != nil {
		return 0, err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Printf("failed to close response: %v\n", err)
		}
	}()
	return resp.StatusCode, nil
}

// if getStatusCode is called repeatedly and we want to use keep-alive connections, we should read the body even though we aren’t interested in it
func (h handler) getStatusCode2(body io.Reader) (int, error) {
	resp, err := h.client.Post(h.url, "application/json", body)
	if err != nil {
		return 0, err
	}
	// Close response body
	_, _ = io.Copy(io.Discard, resp.Body)
	return resp.StatusCode, nil
}
