package main

import (
	"context"
  "fmt"
  "math"
	"os"
	"path/filepath"
	"time"

	"github.com/gamoch/popular_repos/internal/pkg/configuration"
	"github.com/gamoch/popular_repos/pkg/graphql"
	"github.com/gamoch/popular_repos/pkg/graphql/providers/github"
	"github.com/gamoch/popular_repos/pkg/logs"

  "github.com/cheggaaa/pb/v3"
	"github.com/gocarina/gocsv"
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
        mergedPullRequests: pullRequests(states: MERGED) {
          totalCount
        }
        releases {
          totalCount
        }
        primaryLanguage {
          name
        }
        issues {
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

	var nodes []GraphQLNode
	var cursor *string

	barTemplate := `Getting repositories... {{counters . }}`
	bar := pb.ProgressBarTemplate(barTemplate).Start(10)

	for i := 0; i < 10; i++ {
    req.SetVariable("cursor", cursor)
    bar.Increment()

		graphQLData := new(GraphQLData)
		if err := githubClient.Run(ctx, req, graphQLData); err != nil {
			logs.Error.Fatal(err)
		}

		for _, node := range graphQLData.Search.Nodes {
			nodes = append(nodes, node)
		}

		if !graphQLData.Search.PageInfo.HasNextPage {
			break
		}

		cursor = &graphQLData.Search.PageInfo.EndCursor
	}
  bar.Finish()

  var repositories []Repository
  for _, node := range nodes {
    repositories = append(repositories, graphQLNodeToRepository(node))
  }

	if err := writeRepositoriesToCSV(config, repositories); err != nil {
		logs.Error.Fatal(err)
	}
}

func graphQLNodeToRepository(node GraphQLNode) Repository {
	percentageClosedIssues := 1.0
	if node.Issues.TotalCount != 0 {
		percentageClosedIssues = float64(node.ClosedIssues.TotalCount) / float64(node.Issues.TotalCount)
		percentageClosedIssues = math.Round(percentageClosedIssues*100) / 100
	}

	return Repository{
		NameWithOwner:             node.NameWithOwner,
		StargazerCount:            node.StargazerCount,
		Age:                       RepositoryTime{time.Since(node.CreatedAt).Truncate(time.Second)},
		TimeUntilLastUpdate:       RepositoryTime{time.Since(node.UpdatedAt).Truncate(time.Second)},
		TotalAcceptedPullRequests: node.MergedPullRequests.TotalCount,
		TotalReleases:             node.Releases.TotalCount,
		PrimaryLanguage:           node.PrimaryLanguage.Name,
		PercentageClosedIssues:    percentageClosedIssues,
	}
}

func writeRepositoriesToCSV(config *configuration.Config, repositories []Repository) error {
	filePath := "repositories.csv"
	if config.File != "" {
		filePath = config.File

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	if err = gocsv.MarshalFile(&repositories, file); err != nil {
		return err
	}

	fmt.Println("Repositories saved at " + file.Name())

	return nil
}
