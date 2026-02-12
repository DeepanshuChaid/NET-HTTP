package todo

import (
	"encoding/json"
	"errors"
	"io"
  "strconv"
	"net/http"

	"github.com/DeepanshuChaid/NET-HTTP.git/internal/response"
	"github.com/DeepanshuChaid/NET-HTTP.git/internal/storage"
	"github.com/DeepanshuChaid/NET-HTTP.git/internal/types"
	"github.com/go-playground/validator/v10"
)

func New(storage storage.Storage) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    var todo types.Todo

    err := json.NewDecoder(r.Body).Decode(&todo)
    if errors.Is(err, io.EOF) {
      response.WriteJson(w, http.StatusBadRequest, response.GeneralError(errors.New("request body is empty")))
      return
    }

    if err != nil {
      response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
      return
    }

    // validate kar behan ke lode
    if err := validator.New().Struct(todo); err != nil {
      response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
      return
    }

    data, err := storage.Create(
      todo.Title,
      todo.Description,
      todo.Completed,
    )
    if err != nil {
      response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
    }
    
    response.WriteJson(w, http.StatusOK, map[string]any{"success": "OK", "data": *data}) 
  }
}

func GetById(storage storage.Storage) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")

    // Convert string to int
    num, err := strconv.Atoi(id)
    if err != nil {
        // handle error
      response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
    }

    data, err := storage.Delete(num)
    if err != nil {
      response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
      return
    }

    response.WriteJson(w, http.StatusOK, map[string]any{"success": "OK", "data": *data})
    
  }
} 