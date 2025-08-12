package graceful

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ShutdownServer(server *http.Server, timeout time.Duration) {
	stopChan := make(chan os.Signal, 1)

	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)
	<-stopChan

	log.Println("Shutdown signal received, shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Error during server shutdown: %v\n", err)
	} else {
		log.Println("Server stopped gracefully")
	}

}
