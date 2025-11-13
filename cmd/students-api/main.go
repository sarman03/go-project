package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/sarman03/student-api/internal/config"
)

func main() {
	//load config krni h
	//db setup
	//router and server setup

	cfg := config.MustLoad()

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Welcome to students api!"))
	})

	server := http.Server{
		Addr: cfg.HTTPServer.Addr,
		Handler: router,
	}

	slog.Info("Server is starting", "address", cfg.HTTPServer.Addr)
	fmt.Printf("Server is running on port %s", cfg.HTTPServer.Addr)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
				err := server.ListenAndServe()
				if err != nil {
				log.Fatalf("failed to start server: %s", err)
			}
	} ()

	<-done

	slog.Info("Server is shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server shutdown failed", "error", err)
		os.Exit(1)
	}

	slog.Info("Server shutdown successfully")
	os.Exit(0)	
}