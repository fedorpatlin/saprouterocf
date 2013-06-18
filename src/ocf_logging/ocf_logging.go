package ocf_logging

import (
	"fmt"
)

type Logger interface {
	Log(string, string)
}

type Generic_logger struct{}

func (g *Generic_logger) Log(severity string, message string) {
	fmt.Printf("%s - %s\n", severity, message)
}

func Ocf_log_backend(logger Logger, severity string, message string) {
	logger.Log(severity, message)
}
