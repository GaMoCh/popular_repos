package github

import (
	"context"
	"errors"
	"net/http"

	"github.com/gamoch/popular_repos/pkg/graphql"
	"github.com/gamoch/popular_repos/pkg/logs"
)

const endpoint = "https://api.github.com/graphql"

type client struct {
	token  string
	client *graphql.Client
}

func NewClient(token string, options ...graphql.ClientOption) *client {
	if token == "" {
		logs.Error.Fatalln("GITHUB_TOKEN is required")
	}

	return &client{
		token:  token,
		client: graphql.NewClient(endpoint, options...),
	}
}

func (c *client) Run(ctx context.Context, req *graphql.Request, res interface{}) error {
	if req != nil {
		req.Header.Add("Authorization", "bearer "+c.token)
	}

	err := c.client.Run(ctx, req, res)
	if cause := errors.Unwrap(err); cause != nil {
		if errStatus, ok := cause.(*graphql.ErrStatus); ok {
			if errStatus.StatusCode == http.StatusUnauthorized {
				return &Err{err: &ErrInvalidToken{err: err}}
			}
		}
	}

	if err != nil {
		return &Err{err: err}
	}

	return nil
}
