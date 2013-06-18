// saprouterocf project main.go
//
//   Resource Agent for managing simple stateless saprouter resources.
//
//   License:      GNU General Public License (GPL)
//   (c) 2013 Fyodor Patlin
//

package main

import (
	"ocf_const"
	"ocf_functions"
	"os"
)

//Operations start, stop, meta-data and monitor for minimal OCF implementation
func saprouter_start() int {
	return start_service()
}

func saprouter_stop() int {
	verify_all()
	return ocf_const.OCF_SUCCESS
}

func saprouter_metadata() string {
	verify_all()
	return "metadata xml"
}

func saprouter_monitor() int {
	verify_all()
	return check_service()
}

func start_service() int {
	verify_all()
	return ocf_const.OCF_SUCCESS
}

func check_service() int {
	return ocf_const.OCF_SUCCESS
}

func dispatch() int {
	cmd := os.Getenv("__OCF_ACTION")
	if cmd == "" {
		ocf_functions.Ocf_log(ocf_functions.OCF_ERR, "Error: no actions")
		os.Exit(ocf_const.OCF_ERR_ARGS)
	}
	switch cmd {
	case "meta-data":
		saprouter_metadata()
	case "monitor":
		saprouter_monitor()
	case "start":
		saprouter_start()
	case "stop":
		saprouter_stop()
	}
	return ocf_const.OCF_ERR_UNIMPLEMENTED
}

func verify_all() {
	ocf_functions.Check_binary("SAPROUTER")
}

func main() {
	ocf_functions.Ocf_log(ocf_functions.OCF_DEBUG, "resource agent is running")
	dispatch()
}
