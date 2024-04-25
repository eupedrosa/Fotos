package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

    . "github.com/eupedrosa/Fotos/handlers"
)

const (
    major = 0
    minor = 0
    revision = 1
    codename = "copy-cat"
)

var version = fmt.Sprintf("%d.%d.%d %s", major, minor, revision, codename)

func main() {
    fmt.Printf("Fotos WebApp ... v%s\n", version)

    // :: Create an HTTP server with graceful shutdown ::

    server := &http.Server{ Addr: ":4646", Handler: Routes() }
    ctx, stop := context.WithCancel(context.Background())

    // interrupt/quit on these signals
    sig := make(chan os.Signal, 1)
    signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTERM)
    go func(){
        <-sig // wait for the signal

        slog.Info("Graceful shutdown requested!")
        // Allow a 30 seconds shutdown grace period
        shutCtx, cancel := context.WithTimeout(ctx, 30 * time.Second)
        defer cancel()
        go func(){
            <-shutCtx.Done() // wait until the server shutdowns

            if shutCtx.Err() == context.DeadlineExceeded {
                slog.Error("graceful shutdown timeout... forcing exit")
            }
        }()

        // trigger the shutdown
        err := server.Shutdown(shutCtx)
        if err != nil {
            slog.Error("HTTP server failed to shutdown", "error", err)
        }
        stop()
    }()

    slog.Info("listening at :4646")
    err := server.ListenAndServe()
    if err != nil && err != http.ErrServerClosed {
        slog.Error("HTTP server failed to start", "error", err)
        return
    }

    // wait for the server context to finish
    <-ctx.Done()
    slog.Info("Bye bye ...")
}
