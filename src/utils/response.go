package utils

import (
	"encoding/json"
	"net/http"

	"github.com/giftxtrade/api/src/types"
)

func JsonResponse(w http.ResponseWriter, data interface{}) {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		json.NewEncoder(w).Encode(
			types.Response{
				Message: "Could not parse response",
			},
		)
	}
}

// Writes a types.Errors json response to the http.ResponseWriter,
// with a default Http 400 status
func FailResponse(w http.ResponseWriter, errors interface{}) {
	w.WriteHeader(400)
	JsonResponse(w, types.Errors{
		Errors: errors,
	})
}

// Writes a types.Data json response to the http.ResponseWriter,
// with a default Http 200 status
func DataResponse(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(200)
	JsonResponse(w, types.Result{
		Data: data,
	})
}