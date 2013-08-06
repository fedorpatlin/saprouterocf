// ocf_functions.go
package main

import (
	//	"net"
	"ocf_logging"
	"os"
	"os/exec"
	"syscall"
)

var Ocf_logger ocf_logging.Logger

func Ocf_log(severity string, message string) {
	if Ocf_logger == nil {
		Ocf_logger = new(ocf_logging.Generic_logger)
	}
	ocf_logging.Ocf_log_backend(Ocf_logger, severity, message)
}

func Have_binary(exefile string) int {
	if cb := Check_binary(exefile); cb != OCF_SUCCESS {
		return cb
	}
	finfo, _ := os.Stat(exefile)

	if (finfo.Mode() | 0111) != 0 {
		return OCF_ERR_INSTALLED
	}
	return OCF_SUCCESS
}
func Check_binary(exefile string) int {
	_, err := exec.LookPath(exefile)
	if err != nil {
		return OCF_ERR_INSTALLED
	}
	return OCF_SUCCESS
}

func ocf_run(severity string, quiet bool, binary string, command string) int {
	cmd := exec.Command(binary, command)
	if err := cmd.Run(); err != nil {
		Ocf_log(OCF_CRIT, err.Error())
		return OCF_ERR_INSTALLED
	}

	syscall.Setsid()
	return OCF_SUCCESS
}

func ocf_daemon(daemon func() int) int {
	syscall.Umask(27)
	if err := daemon(); err > OCF_SUCCESS {
		return OCF_ERR_GENERIC
	}
	os.Stdin.Close()
	os.Stdout.Close()
	os.Stderr.Close()
	syscall.Setsid()
	return (OCF_SUCCESS)
}

func Ocf_is_true() {}

func check_port(host, port string) int {
	return OCF_SUCCESS
}
