package response

import (
	"encoding/json"
	"net/http"
  "fmt"
	"github.com/go-playground/validator/v10"
  "strings"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOk    = "OK"
	StatusError = "Error"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)

	return err
}

func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}


func ValidatorError (errs validator.ValidationErrors) Response {
  var errMsg []string

  for _, err := range errs {
    switch err.ActualTag() {
      case "required":
        errMsg = append(errMsg, fmt.Sprintf("field %s is a required field", err.Field()))
      default:
        errMsg = append(errMsg, fmt.Sprintf("field %s is not valid", err.Field()))
    }
  }
  
  return Response{
    Status: StatusError,
    Error: strings.Join(errMsg, ", "),
  }
}