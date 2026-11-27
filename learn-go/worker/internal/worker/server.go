package worker

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const serverTimeout = 5 * time.Second

func RunHTTP(addr string, handler func(*Queue) http.Handler) error {
	client := &http.Client{
		CheckRedirect: func(*http.Request, []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	defer client.CloseIdleConnections()

	return Run(addr, HTTPFetch(client), func(queue *Queue) (
		func(net.Listener) error, func(context.Context) error,
	) {
		server := &http.Server{Handler: handler(queue), ReadHeaderTimeout: serverTimeout}
		return server.Serve, server.Shutdown
	})
}

// Run starts the worker and its HTTP server, and waits for both on shutdown.
func Run(addr string, fetch FetchFunc, server func(*Queue) (
	func(net.Listener) error, func(context.Context) error,
)) error {
	externalURL := os.Getenv("EXTERNAL_URL")
	if externalURL == "" {
		externalURL = "https://jsonplaceholder.typicode.com/todos/1"
	}
	parsed, err := url.Parse(externalURL)
	if err != nil || parsed.Hostname() == "" || (parsed.Scheme != "http" && parsed.Scheme != "https") {
		return errors.New("EXTERNAL_URL must be an absolute HTTP or HTTPS URL")
	}
	if configuredAddr := os.Getenv("ADDR"); configuredAddr != "" {
		addr = configuredAddr
	}
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer listener.Close()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	queue := &Queue{jobs: make(chan uint64, queueSize)}
	serve, shutdown := server(queue)
	workerDone := make(chan struct{})
	go func() {
		defer close(workerDone)
		queue.run(ctx, parsed.String(), fetch)
	}()
	defer func() {
		stop()
		<-workerDone
	}()

	serverDone := make(chan error, 1)
	go func() { serverDone <- serve(listener) }()
	log.Printf("listening on http://%s (POST /jobs)", listener.Addr())

	select {
	case err := <-serverDone:
		return err
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), serverTimeout)
		defer cancel()
		err := shutdown(shutdownCtx)
		// Also close the listener if the signal arrived before Serve started.
		_ = listener.Close()
		<-serverDone
		return err
	}
}
