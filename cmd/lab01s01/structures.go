package main

import "time"

type GraphQLData struct {
	Search struct {
		Nodes []struct {
			NameWithOwner  string    `json:"nameWithOwner"`
			StargazerCount int       `json:"stargazerCount"`
			CreatedAt      time.Time `json:"createdAt"`
		} `json:"nodes"`
	} `json:"search"`
}
