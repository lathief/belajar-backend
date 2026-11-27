package main

import (
	"context"
	"crypto/tls"
	"errors"
	"log"
	"net"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/lathief/test-go/worker/internal/worker"
)

func main() {
	if err := worker.Run("127.0.0.1:8081", fetch, func(queue *worker.Queue) (
		func(net.Listener) error, func(context.Context) error,
	) {
		app := fiber.New(fiber.Config{DisableStartupMessage: true, ReadTimeout: 5 * time.Second})
		app.Post("/jobs", func(c *fiber.Ctx) error {
			status, result := queue.Submit()
			return c.Status(status).JSON(result)
		})
		return app.Listener, app.ShutdownWithContext
	}); err != nil {
		log.Fatal(err)
	}
}

func fetch(ctx context.Context, url string) (int, int64, error) {
	if err := ctx.Err(); err != nil {
		return 0, 0, err
	}
	deadline, _ := ctx.Deadline() // Set by the worker for each job.
	timeout := time.Until(deadline)
	if timeout <= 0 {
		return 0, 0, context.DeadlineExceeded
	}

	agent := fiber.AcquireAgent()
	req := agent.Request()
	req.Header.SetMethod(fiber.MethodGet)
	req.SetRequestURI(url)
	if err := agent.Parse(); err != nil {
		fiber.ReleaseAgent(agent)
		return 0, 0, err
	}

	agent.HostClient.MaxResponseBodySize = worker.MaxResponseBody
	agent.TLSConfig(&tls.Config{MinVersion: tls.VersionTLS12})
	status, body, errs := agent.Set("Accept", "application/json").Timeout(timeout).Bytes()
	return status, int64(len(body)), errors.Join(errs...)
}
