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
	"github.com/DeepanshuChaid/NET-HTTP.git/internal/storage/sqlite"
)

func main () {
  config := config.MustLoad()

  db, err := sqlite.New(config)
  if err != nil {
    log.Fatal("Error white initiating database")
  }


  router := http.NewServeMux()

  router.HandleFunc("POST /create", todo.New(db))

  router.HandleFunc("GET /get/{id}", todo.GetById(db))

  router.HandleFunc("DELETE /delete/{id}", todo.Delete(db))

  router.HandleFunc("PUT /update/{id}", todo.Update(db))

  router.HandleFunc("GET /all", todo.GetAll(db)

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

  err = server.Shutdown(ctx)
  if err != nil {
    log.Fatalf("cannot shutdown server: %s", err)
  }

  slog.Info("Server stopped")
}