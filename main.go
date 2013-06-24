// saprouterocf project main.go
//
//   Resource Agent for managing simple stateless saprouter resources.
//
//   License:      GNU General Public License (GPL)
//   (c) 2013 Fyodor Patlin
//

package main

import (
	"fmt"
	"ocf_const"
	"ocf_functions"
	"os"
)

const saprouter_binary string = "/etc/init.d/saprouter"

//Operations start, stop, meta-data and monitor for minimal OCF implementation
func saprouter_start() int {
	//verify_all()
	return ocf_functions.Ocf_run("", false, saprouter_binary, "start")
}

//returns OCF_SUCCESS if successfully stop process or process is not running
func saprouter_stop() int {
	//verify_all()
	return ocf_functions.Ocf_run("", false, saprouter_binary, "stop")
}

func saprouter_metadata() int {
	//	verify_all()
	fmt.Printf("%s\n", metadata_xml)
	return ocf_const.OCF_SUCCESS
}

//returns OCF_SUCCESS or OCF_NOT_RUNNING
func saprouter_monitor() int {
	//verify_all()
	rc := ocf_functions.Ocf_run("", false, saprouter_binary, "status")
	ocf_functions.Ocf_log(ocf_functions.OCF_ERR, string(rc))
	if rc == 0 {
		return ocf_const.OCF_SUCCESS
	} else {
		return ocf_const.OCF_NOT_RUNNING
	}
}

func dispatch() {
	cmd := os.Getenv("__OCF_ACTION")
	if cmd == "" {
		saprouter_metadata()
		os.Exit(ocf_const.OCF_ERR_ARGS)
	}
	switch cmd {
	case "meta-data":
		os.Exit(saprouter_metadata())
	case "monitor":
		ocf_functions.Ocf_log(ocf_functions.OCF_ERR, "monitor action is running")
		os.Exit(saprouter_monitor())
	case "start":
		os.Exit(saprouter_start())
	case "stop":
		os.Exit(saprouter_stop())
	case "validate-all":
		os.Exit(0)
	}
	os.Exit(ocf_const.OCF_ERR_UNIMPLEMENTED)
}

func verify_all() {
	ocf_functions.Check_binary(saprouter_binary)
}

func main() {
	//ocf_functions.Ocf_log(ocf_functions.OCF_DEBUG, "resource agent is running")
	dispatch()
}

var metadata_xml = "<?xml version=\"1.0\"?><!DOCTYPE resource-agent SYSTEM \"ra-api-1.dtd\">"
