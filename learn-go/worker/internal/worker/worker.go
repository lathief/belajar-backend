package worker

import (
	"context"
	"errors"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

const (
	queueSize       = 8
	requestTimeout  = 5 * time.Second
	MaxResponseBody = 1 << 20 // 1 MiB
)

type FetchFunc func(ctx context.Context, url string) (status int, size int64, err error)

type Queue struct {
	jobs   chan uint64
	nextID atomic.Uint64
}

type JobResponse struct {
	ID     uint64 `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
	Error  string `json:"error,omitempty"`
}

// Submit returns immediately even when the worker is busy.
func (q *Queue) Submit() (int, JobResponse) {
	id := q.nextID.Add(1)
	select {
	case q.jobs <- id:
		return http.StatusAccepted, JobResponse{ID: id, Status: "queued"}
	default:
		return http.StatusServiceUnavailable, JobResponse{Error: "job queue is full"}
	}
}

func (q *Queue) run(ctx context.Context, url string, fetch FetchFunc) {
	for ctx.Err() == nil {
		select {
		case <-ctx.Done():
			return
		case id := <-q.jobs:
			requestCtx, cancel := context.WithTimeout(ctx, requestTimeout)
			status, size, err := fetch(requestCtx, url)
			cancel()

			switch {
			case err != nil:
				log.Printf("job=%d failed error=%v", id, err)
			case status < 200 || status >= 300:
				log.Printf("job=%d failed http_status=%d", id, status)
			default:
				log.Printf("job=%d completed http_status=%d bytes=%d", id, status, size)
			}
		}
	}
}

var ErrResponseTooLarge = errors.New("upstream response exceeds 1 MiB")
