package ocf_logging

import (
	"fmt"
)

type Generic_logger struct {
}

func (Generic_logger) Log(severity string, message string) {
	fmt.Printf("%s - %s\n", severity, message)
}

func Ocf_log_backend(logger Generic_logger, severity string, message string) {
	logger.Log(severity, message)
}
