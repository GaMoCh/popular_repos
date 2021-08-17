package graphql

import "net/http"

type Request struct {
	requestBody requestBody
	Header      http.Header
}

type requestBody struct {
	Query     string           `json:"query"`
	Variables RequestVariables `json:"variables,omitempty"`
}

type RequestVariables map[string]interface{}

type RequestOption func(*Request)

func NewRequest(query string, options ...RequestOption) *Request {
	request := &Request{
		requestBody: requestBody{
			Query:     query,
			Variables: make(RequestVariables),
		},
		Header: http.Header{
			"Content-Type": {"application/json; charset=utf-8"},
			"Accept":       {"application/json; charset=utf-8"},
		},
	}

	for _, optionFunc := range options {
		optionFunc(request)
	}

	return request
}

func WithHTTPHeader(httpHeader http.Header) RequestOption {
	return func(request *Request) {
		for key, values := range httpHeader {
			for _, value := range values {
				request.Header.Add(key, value)
			}
		}
	}
}

func WithVariables(variables RequestVariables) RequestOption {
	return func(request *Request) {
		request.requestBody.Variables = variables
	}
}

func (r *Request) GetQuery() string {
	return r.requestBody.Query
}

func (r *Request) GetVariables() RequestVariables {
	return r.requestBody.Variables
}

func (r *Request) SetVariable(key string, value interface{}) {
	r.requestBody.Variables[key] = value
}
