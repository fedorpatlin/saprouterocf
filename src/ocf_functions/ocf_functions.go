// ocf_functions.go
package ocf_functions

import (
	"ocf_const"
	"ocf_logging"
	"os"
	"os/exec"
)

const (
	OCF_DEBUG = "debug"
	OCF_INFO  = "info"
	OCF_WARN  = "warn"
	OCF_ERR   = "err"
	OCF_CRIT  = "crit"
)

var Ocf_logger ocf_logging.Logger

func Ocf_log(severity string, message string) {
	if Ocf_logger == nil {
		Ocf_logger = new(ocf_logging.Generic_logger)
	}
	ocf_logging.Ocf_log_backend(Ocf_logger, severity, message)
}

func Have_binary(exefile string) int {
	_, err := exec.LookPath(exefile)
	if err != nil {
		return ocf_const.OCF_ERR_INSTALLED
	}
	return ocf_const.OCF_SUCCESS
}
func Check_binary(exefile string) {
	_, err := exec.LookPath(exefile)
	if err != nil {
		os.Exit(ocf_const.OCF_ERR_INSTALLED)
	}
}