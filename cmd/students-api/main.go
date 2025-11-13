package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall" 
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

}