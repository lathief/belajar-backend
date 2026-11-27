package worker

import (
	"context"
	"io"
	"net/http"
)

// HTTPFetch builds a fetch function for the Gin and net/http examples.
func HTTPFetch(client *http.Client) FetchFunc {
	return func(ctx context.Context, url string) (int, int64, error) {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return 0, 0, err
		}
		req.Header.Set("Accept", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			return 0, 0, err
		}
		defer resp.Body.Close()

		size, err := io.Copy(io.Discard, io.LimitReader(resp.Body, MaxResponseBody+1))
		if size > MaxResponseBody {
			return resp.StatusCode, size, ErrResponseTooLarge
		}
		return resp.StatusCode, size, err
	}
}
