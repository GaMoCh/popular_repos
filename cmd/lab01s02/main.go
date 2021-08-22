package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/gamoch/popular_repos/internal/pkg/configuration"
	"github.com/gamoch/popular_repos/pkg/graphql"
	"github.com/gamoch/popular_repos/pkg/graphql/providers/github"
	"github.com/gamoch/popular_repos/pkg/logs"
)

const query = `query PopularRepos($cursor: String) {
  search(query: "stars:>1", type: REPOSITORY, first: 100, after: $cursor) {
    pageInfo {
      hasNextPage
      endCursor
    }
    nodes {
      ... on Repository {
        nameWithOwner
        stargazerCount
        createdAt
        updatedAt
        pullRequests(states: MERGED) {
          totalCount
        }
        releases {
          totalCount
        }
        primaryLanguage {
          name
        }
        openIssues: issues(states: OPEN) {
          totalCount
        }
        closedIssues: issues(states: CLOSED) {
          totalCount
        }
      }
    }
  }
}`

func main() {
	config := configuration.Get()
	ctx := context.Background()

	githubClient := github.NewClient(config.Token)
	req := graphql.NewRequest(query)

	graphQLData := new(GraphQLData)
	if err := githubClient.Run(ctx, req, graphQLData); err != nil {
		logs.Error.Fatal(err)
	}

	graphqlJSON, err := json.MarshalIndent(graphQLData, "", "  ")
	if err != nil {
		logs.Error.Fatal(err)
	}

	graphqlJSON = append(graphqlJSON, '\n')
	os.Stdout.Write(graphqlJSON)
}
