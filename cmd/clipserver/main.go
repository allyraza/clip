package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/allyraza/clip"
)

func main() {
	var (
		addr = flag.String("addr", "localhost:8080", "address to listen on.")
	)

	flag.Parse()

	s, err := clip.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting server on %s\n", *addr)

	server := http.Server{
		Addr:    *addr,
		Handler: s,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	signalc := make(chan os.Signal, 1)
	signal.Notify(signalc, syscall.SIGINT, syscall.SIGTERM)
	<-signalc

	log.Printf("Shutdown signal received, quiting...")
	server.Shutdown(context.Background())
}
