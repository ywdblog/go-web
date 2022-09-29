package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/WiFeng/go-sky/log"
	kitendpoint "github.com/go-kit/kit/endpoint"
)

func genericErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	log.Error(ctx, err)
	// TODO:
	resp := ""
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(resp)
}

// encodeHTTPGenericResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(kitendpoint.Failer); ok && f.Failed() != nil {
		genericErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
