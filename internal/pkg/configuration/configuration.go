package configuration

import (
	"flag"
	"os"
)

type config struct {
	Token string
}

func Get() *config {
	token := flag.String("token", os.Getenv("GITHUB_TOKEN"), "GitHub Token (env: GITHUB_TOKEN)")
	flag.Parse()

	return &config{
		Token: *token,
	}
}
