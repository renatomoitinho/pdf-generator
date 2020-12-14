package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "time"

    "pdf-generation/internal/pkg/router"
)

func main() {
    basicRouter, err := router.Router()
    if err != nil {
        log.Fatal(err.Error())
    }

    readHeaderTimeout, _ := time.ParseDuration("1000ms")
    shutdownTimeout, _ := time.ParseDuration("10000ms")

    server := http.Server{
        Addr:              "0.0.0.0:8080",
        Handler:           basicRouter,
        ReadHeaderTimeout: readHeaderTimeout,
    }

    stop := make(chan os.Signal)
    signal.Notify(stop, os.Interrupt)

    go func() {
        log.Printf("starting the http server at %s", server.Addr)
        if err := server.ListenAndServe(); err != nil {
            if err != http.ErrServerClosed {
                log.Printf("http ErrServerClosed: %s", err.Error())
                log.Fatal(err)
            }
        }
    }()

    <-stop

    log.Printf("shutting down the http server")
    ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
    defer cancel()
    if err := server.Shutdown(ctx); err != nil {
        log.Fatalf("unable to shutdown the http server: %s", err.Error())
    }
    log.Printf("http server successfully shut down")
}
