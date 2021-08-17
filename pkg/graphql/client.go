package graphql

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type Client struct {
	endpoint   string
	httpClient *http.Client
}

type ClientOption func(*Client)

type ClientRunner interface {
	Run(ctx context.Context, req *Request, res interface{}) error
}

func NewClient(endpoint string, options ...ClientOption) *Client {
	client := &Client{
		endpoint:   endpoint,
		httpClient: http.DefaultClient,
	}

	for _, optionFunc := range options {
		optionFunc(client)
	}

	return client
}

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(client *Client) {
		client.httpClient = httpClient
	}
}

func (c *Client) Run(ctx context.Context, req *Request, res interface{}) error {
	if req == nil {
		return &Err{err: ErrNilRequest}
	}

	if ctx == nil {
		ctx = context.Background()
	} else if err := ctx.Err(); err != nil {
		return &Err{err: err}
	}

	reqBody := new(bytes.Buffer)
	if err := json.NewEncoder(reqBody).Encode(req.requestBody); err != nil {
		return &Err{err: err}
	}

	graphQLReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.endpoint, reqBody)
	if err != nil {
		return &Err{err: err}
	}
	graphQLReq.Header = req.Header

	resp, err := c.httpClient.Do(graphQLReq)
	if err != nil {
		return &Err{err: err}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &Err{err: &ErrStatus{
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
		}}
	}

	graphQLRes := &response{Data: res}
	if err = json.NewDecoder(resp.Body).Decode(graphQLRes); err != nil {
		return &Err{err: err}
	}

	if len(graphQLRes.Errors) > 0 {
		return &Err{err: joinResponseErrors(graphQLRes.Errors)}
	}

	return nil
}
