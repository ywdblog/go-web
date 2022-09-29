package client

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	skyhttp "github.com/WiFeng/go-sky/http"
	"github.com/WiFeng/go-sky/log"
	. "github.com/xiwujie/article/pkg/entity"
)

type Channel struct {
}

func NewChannel() *Channel {
	return &Channel{}
}

func decodeChannelCallbackResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp ChannelCallbackResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

func (r *Channel) Callback(ctx context.Context, req ChannelCallbackRequest) (ChannelCallbackResponse, error) {
	var resp ChannelCallbackResponse

	cli, err := skyhttp.NewClient(ctx, sevice, http.MethodPost, ChannelCallbackURI,
		encodeHTTPGenericRequest, decodeChannelCallbackResponse)
	if err != nil {
		return resp, err
	}

	result, err := cli.Endpoint()(ctx, req)
	if err != nil {
		log.Errorw(ctx, "cli.Endpoint error", "req", req, "err", err)
		return resp, err
	}

	resp = result.(ChannelCallbackResponse)
	return resp, nil
}
