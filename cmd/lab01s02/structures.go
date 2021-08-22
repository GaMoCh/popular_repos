package main

import "time"

type GraphQLData struct {
	Search struct {
		PageInfo struct {
			HasNextPage bool   `json:"hasNextPage"`
			EndCursor   string `json:"endCursor"`
		} `json:"pageInfo"`
		Nodes []struct {
			NameWithOwner  string    `json:"nameWithOwner"`
			StargazerCount int       `json:"stargazerCount"`
			CreatedAt      time.Time `json:"createdAt"`
			UpdatedAt      time.Time `json:"updatedAt"`
			PullRequests   struct {
				TotalCount int `json:"totalCount"`
			} `json:"pullRequests"`
			Releases struct {
				TotalCount int `json:"totalCount"`
			} `json:"releases"`
			PrimaryLanguage *struct {
				Name string `json:"name"`
			} `json:"primaryLanguage"`
			OpenIssues struct {
				TotalCount int `json:"totalCount"`
			} `json:"openIssues"`
			ClosedIssues struct {
				TotalCount int `json:"totalCount"`
			} `json:"closedIssues"`
		} `json:"nodes"`
	} `json:"search"`
}
