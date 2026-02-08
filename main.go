package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
  "context"

	"github.com/DeepanshuChaid/NET-HTTP.git/internal/config"
)

func main () {
  config := config.MustLoad()


  router := http.NewServeMux()

  router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello i will make a todo list api"))
  })

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