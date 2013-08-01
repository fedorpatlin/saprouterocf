package ocf_logging

import (
	"fmt"
	"os"
)

const LOGFILE = "/var/log/saprouterocf.log"

type Logger interface {
	Log(string, string)
}

type Generic_logger struct{}

func (g *Generic_logger) Log(severity string, message string) {
	logf, err := os.OpenFile(LOGFILE, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeType)
	if err != nil {
		fmt.Printf("can't open file: %s\n", err.Error())
		os.Exit(1)
	}
	logf.WriteString(fmt.Sprintf("%s - %s\n", severity, message))
	return
}

func Ocf_log_backend(logger Logger, severity string, message string) {
	logger.Log(severity, message)
}
