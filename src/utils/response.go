package utils

import (
	"encoding/json"
	"net/http"

	"github.com/giftxtrade/api/src/types"
)

func write_json(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		json.NewEncoder(w).Encode(
			types.Response{
				Message: "Could not parse response",
			},
		)
	}
}

func JsonResponse(w http.ResponseWriter, data interface{}) {
	write_json(w, 200, data)
}

// Writes a types.Errors json response to the http.ResponseWriter,
// with a default Http 400 status
func FailResponse(w http.ResponseWriter, errors interface{}) {
	write_json(w, 400, types.Errors{
		Errors: errors,
	})
}

func FailResponseUnauthorized(w http.ResponseWriter, errors interface{}) {
	write_json(w, 401, types.Errors{
		Errors: errors,
	})
}

// Writes a types.Data json response to the http.ResponseWriter,
// with a default Http 200 status
func DataResponse(w http.ResponseWriter, data interface{}) {
	write_json(w, 200, types.Result{
		Data: data,
	})
}