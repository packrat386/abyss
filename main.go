package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	srv := http.Server{
		Addr:    os.Getenv("ADDR"),
		Handler: http.HandlerFunc(handler),
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigterm := make(chan os.Signal, 1)
		signal.Notify(sigterm, os.Interrupt)
		<-sigterm

		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}

		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		bucketHandler(w, r)
	} else {
		objectHandler(w, r)
	}
}

func bucketHandler(w http.ResponseWriter, r *http.Request) {
	renderError(
		w,
		http.StatusForbidden,
		"AccessDenied",
		"Access Denied",
		"/",
		"badfood",
	)
}

func objectHandler(w http.ResponseWriter, r *http.Request) {
	renderError(
		w,
		http.StatusNotFound,
		"NoSuchKey",
		"The specified key does not exist.",
		r.URL.Path,
		"badfood",
	)
}
