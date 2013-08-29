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
	pth, cb := Check_binary(exefile)
	if cb != OCF_SUCCESS {
		return cb
	}
	finfo, err := os.Stat(pth)
	if err != nil {
		return OCF_ERR_INSTALLED
	}
	if (finfo.Mode() | 0111) != 0 {
		return OCF_ERR_INSTALLED
	}
	return OCF_SUCCESS
}
func Check_binary(exefile string) (string, int) {
	path, err := exec.LookPath(exefile)
	if err != nil {
		return "", OCF_ERR_INSTALLED
	}
	return path, OCF_SUCCESS
}

func ocf_run(binary string, params ...string) error {
	cmd := exec.Command(binary, params...)
	return cmd.Run()
}

func ocf_daemon(binary string, params ...string) error {
	syscall.Umask(27)

	nulldev, _ := os.Open("/dev/null")
	cmd := exec.Command(binary, params...)

	err := cmd.Start()

	syscall.Dup2(int(nulldev.Fd()), int(os.Stdout.Fd()))
	syscall.Dup2(int(nulldev.Fd()), int(os.Stderr.Fd()))
	syscall.Dup2(int(nulldev.Fd()), int(os.Stdin.Fd()))

	os.Stdout.Close()
	os.Stdin.Close()
	os.Stderr.Close()

	if _, err1 := syscall.Setsid(); err1 != nil {
		Ocf_log(OCF_INFO, "Oops! "+err1.Error())
		os.Exit(255)
	}
	return err
}

func Ocf_is_true() {}

func check_port(host, port string) int {
	return OCF_SUCCESS
}
