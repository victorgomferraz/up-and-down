package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	port, delay := getPort(), getDelay()

	log.Printf("Starting HTTP server on port %s with delay %d", port, delay)
	httpServerExitDone := &sync.WaitGroup{}
	httpServerExitDone.Add(1)
	srv := startHttpServer(httpServerExitDone, port)

	log.Printf("Server started, wait for %d seconds", delay)
	time.Sleep(time.Duration(delay) * time.Second)
	log.Printf("Stopping HTTP server")

	if err := srv.Shutdown(context.TODO()); err != nil {
		panic(err)
	}

	httpServerExitDone.Wait()
	log.Printf("Server stopped, exiting")
}

func startHttpServer(wg *sync.WaitGroup, port string) *http.Server {
	srv := &http.Server{Addr: ":" + port}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "I'm ready.\n")
	})

	go func() {
		defer wg.Done()
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	return srv
}

func getPort() string {
	port := os.Getenv("UP_AND_DOWN_PORT")
	if _, err := strconv.Atoi(port); err != nil {
		panic("Port is not a number:" + port)
	}
	return port
}

func getDelay() int {
	delayString := os.Getenv("UP_AND_DOWN_DELAY")
	if delayString == "" {
		delayString = "3"
	}
	if delay, err := strconv.Atoi(delayString); err != nil {
		panic("Delay is not a number:" + delayString)
	} else {
		return delay
	}
}
