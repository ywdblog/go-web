package http

import (
	"net/http"

	skyhttp "github.com/WiFeng/go-sky/http"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/xiwujie/article/pkg/endpoint"
	. "github.com/xiwujie/article/pkg/entity"
)

// NewHandler returns an HTTP handler that makes a set of endpoints
// available on predefined paths.
func NewHandler(endpoints endpoint.Endpoints) http.Handler {
	r := skyhttp.NewRouter()

	genericOptions := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(genericErrorEncoder),
	}

	r.Methods(http.MethodPost).Path(ArticleInfoMgetURI).Handler(skyhttp.NewServer(
		endpoints.Article.MGet,
		decodeHTTPArticleInfoMgetRequest,
		encodeHTTPGenericResponse,
		genericOptions...,
	))

	return r
}
