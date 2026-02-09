package todo 

import (
 "net/http"
  "github.com/DeepanshuChaid/NET-HTTP.git/internal/types"
  "encoding/json"
  "errors"
  "io"
)

func New() http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    var todo types.Todo

    err := json.NewDecoder(r.Body).Decode(&todo)
    if errors.Is(err, io.EOF) {
      w.WriteHeader(http.StatusBadRequest)
    }
    
    w.Write([]byte("Hello, World!"))
  }
}