package http

import (
	"context"
	"encoding/json"
	"net/http"

	. "github.com/xiwujie/article/pkg/entity"
)

func decodeHTTPArticleInfoMgetRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req ArticleInfoMGetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return req, err
	}
	return req, nil
}
