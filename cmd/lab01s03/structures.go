package main

import (
	"time"
)

type GraphQLData struct {
	Search struct {
		PageInfo struct {
			HasNextPage bool   `json:"hasNextPage"`
			EndCursor   string `json:"endCursor"`
		} `json:"pageInfo"`
		Nodes []GraphQLNode `json:"nodes"`
	} `json:"search"`
}

type GraphQLNode struct {
	NameWithOwner      string    `json:"nameWithOwner"`
	StargazerCount     int       `json:"stargazerCount"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
	MergedPullRequests struct {
		TotalCount int `json:"totalCount"`
	} `json:"mergedPullRequests"`
	Releases struct {
		TotalCount int `json:"totalCount"`
	} `json:"releases"`
	PrimaryLanguage struct {
		Name string `json:"name"`
	} `json:"primaryLanguage"`
	Issues struct {
		TotalCount int `json:"totalCount"`
	} `json:"issues"`
	ClosedIssues struct {
		TotalCount int `json:"totalCount"`
	} `json:"closedIssues"`
}

type Repository struct {
	NameWithOwner             string  `csv:"nameWithOwner"`
	StargazerCount            int     `csv:"stargazerCount"`
	Age                       int     `csv:"age"`
	TimeUntilLastUpdate       int     `csv:"timeUntilLastUpdate"`
	TotalAcceptedPullRequests int     `csv:"totalAcceptedPullRequests"`
	TotalReleases             int     `csv:"totalReleases"`
	PrimaryLanguage           string  `csv:"primaryLanguage"`
	PercentageClosedIssues    float64 `csv:"percentageClosedIssues"`
}
