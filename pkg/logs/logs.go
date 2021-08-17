package logs

import (
	"log"
	"os"
)

var (
	Error = log.New(os.Stderr, "[ERROR] ", 0)
)
