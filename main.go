package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DeepanshuChaid/NET-HTTP.git/internal/config"
	"github.com/DeepanshuChaid/NET-HTTP.git/internal/http/handlers/todo"
)

func main () {
  config := config.MustLoad()


  router := http.NewServeMux()

  router.HandleFunc("/", todo.New())

  server := http.Server{
    Addr: config.HttpServer.Address,
    Handler: router,
  }
  

  done := make(chan os.Signal, 1)

  signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
  
  go func() {
    err := server.ListenAndServe()
    if err != nil {
      log.Fatalf("cannot start server: %s", err)
    }
  }()
  
  <- done

  slog.Info("Shuttting down server")

  ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
  defer cancel()

  err := server.Shutdown(ctx)
  if err != nil {
    log.Fatalf("cannot shutdown server: %s", err)
  }

  slog.Info("Server stopped")
}