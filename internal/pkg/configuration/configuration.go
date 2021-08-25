package configuration

import (
	"flag"
	"os"
)

type Config struct {
	Token string
	File  string
}

func Get() *Config {
	token := flag.String("token", os.Getenv("GITHUB_TOKEN"), "GitHub Token (env: GITHUB_TOKEN)")
	file := flag.String("file", os.Getenv("CSV_FILE"), "CSV File Location (env: CSV_FILE)")
	flag.Parse()

	return &Config{
		Token: *token,
		File:  *file,
	}
}
