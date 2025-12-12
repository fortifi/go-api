package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-openapi/runtime"
)

type errorResponse struct {
	Payload struct {
		Data []interface{} `json:"data"`
		Meta struct {
			Code      int    `json:"code"`
			Message   string `json:"message"`
			RequestId string `json:"requestId"`
			Type      string `json:"type"`
		} `json:"meta"`
	} `json:"Payload"`
}

type Transport struct {
	original runtime.ClientTransport
	clear    func(transport runtime.ClientTransport) error
}

func NewTransport(transport runtime.ClientTransport, clear func(transport runtime.ClientTransport) error) *Transport {
	return &Transport{original: transport, clear: clear}
}

func (t *Transport) Submit(operation *runtime.ClientOperation) (interface{}, error) {
	resp, err := t.original.Submit(operation)

	if err != nil {
		if rawR, ok := err.(runtime.ClientResponseStatus); ok && rawR.IsCode(http.StatusForbidden) {
			env := &errorResponse{}
			if bytes, mErr := json.Marshal(err); mErr == nil {
				err = json.Unmarshal(bytes, env)
			}

			if env.Payload.Meta.Type == "InvalidTokenException" || env.Payload.Meta.Message == "A valid access token must be supplied" {
				clearErr := t.clear(t.original)
				if clearErr == nil {
					return t.original.Submit(operation)
				}
			}
		}
	}

	return resp, err
}
