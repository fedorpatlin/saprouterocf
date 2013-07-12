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
	"os"
	"time"
)

//Operations start, stop, meta-data and monitor for minimal OCF implementation
func saprouter_start() int {
	verify_all()
	Ocf_log(OCF_INFO, "starting saprouter process")
	rc := run_service("-r")
	if rc != OCF_SUCCESS {
		Ocf_log(OCF_ERR, "starting saprouter process failed, code "+string(rc))
		return OCF_ERR_GENERIC
	}
	status := saprouter_monitor()
	for status != OCF_SUCCESS {
		time.Sleep(1 * time.Second)
		status = saprouter_monitor()
	}
	Ocf_log(OCF_INFO, "saprouter process is working")
	return status
}

//returns OCF_SUCCESS if successfully stop process or process is not running
func saprouter_stop() int {
	verify_all()
	run_service("-s")
	return OCF_SUCCESS
}

func saprouter_reload() int {
	rc := run_service("-n")
	if rc == 0 {
		return OCF_SUCCESS
	} else {
		return OCF_NOT_RUNNING
	}
}

func saprouter_metadata() int {
	//	verify_all()
	fmt.Printf("%s\n", metadata_xml)
	return OCF_SUCCESS
}

func saprouter_monitor() int {
	//verify_all()
	rc1 := run_service("-L")
	rc2 := check_port(get_param("host"), get_param("port"))
	Ocf_log(OCF_ERR, string(rc1))
	if rc1 == 0 {
		return OCF_SUCCESS
	} else {
		return OCF_NOT_RUNNING
	}
}

func run_service(params string) int {
	return Ocf_run("info", false, get_param("binary"), params)
}

func dispatch() {
	cmd := os.Getenv("__OCF_ACTION")
	if cmd == "" {
		cmd = "meta-data"
	}
	switch cmd {
	case "meta-data":
		os.Exit(saprouter_metadata())
	case "monitor":
		//		Ocf_log(OCF_ERR, "monitor action is running")
		os.Exit(saprouter_monitor())
	case "start":
		os.Exit(saprouter_start())
	case "stop":
		os.Exit(saprouter_stop())
	case "reload":
		os.Exit(saprouter_reload())
	case "validate-all":
		os.Exit(0)
	}
	os.Exit(OCF_ERR_UNIMPLEMENTED)
}

func verify_all() {
	Check_binary(get_param("binary"))
}

func get_param(name string) string {
	value := os.Getenv(fmt.Sprintf("OCF_RESKEY_%s", name))
	if value == "" {
		value = os.Getenv(fmt.Sprintf("OCF_RESKEY_%s_default", name))
	}
	return value
}

func set_default_param(name string, value string) {
	vname := fmt.Sprintf("OCF_RESKEY_%s_default", name)
	os.Setenv(vname, value)
}

func myinit() {

}

func main() {
	//Ocf_log(OCF_DEBUG, "resource agent is running")
	myinit()
	dispatch()
}

var metadata_xml = `<?xml version="1.0"?>
<!DOCTYPE resource-agent SYSTEM "ra-api-1.dtd">
<resource-agent name="saprouterocf" version="0.01">
	<version>0.01</version>
	<longdesc lang="en">
		resource agent for saprouter installed from special rpm package
	</longdesc>
	<shortdesc>
		saprouter
	</shortdesc>
	<parameters>
		<parameter name="binary" required="1" unique="0">
			<longdesc lang="en">Path to program</longdesc>
			<shortdesc>Program location</shortdesc>
			<content type="string" default="saprouter"/>
		</parameter>
		<parameter name="config" required="1" unique="0">
			<longdesc lang="en">Configuration file location</longdesc>
			<shortdesc>Configfile</shortdesc>
			<content type="string" default="/etc/saprouter/saprouttab"/>
		</parameter>
		<parameter name="log" required="0" unique="0">
			<longdesc lang="en"> the full path to a log file</longdesc>
			<shortdesc>Log file</shortdesc>
			<content type="string" default="/var/log/saprouter/saprouter.log"/>
		</parameter>
		<parameter name="trace" required="0" unique="0">
			<longdesc lang="en"> the full path to a trace file</longdesc>
			<shortdesc>Trace file</shortdesc>
			<content type="string" default="/var/log/saprouter/saprouter.trc"/>
		</parameter>
		<parameter name="host" required="1" unique="0">
			<longdesc lang="en">host name or IP address</longdesc>
			<shortdesc>host</shortdesc>
			<content type="string" default="localhost"/>
		</parameter>
		<parameter name="port" required="1" unique="0">
			<longdesc lang="en">Port number</longdesc>
			<shortdesc>Port number</shortdesc>
			<content type="num" default="3299"/>
		</parameter>
	</parameters>
	<actions>
		<action name="start" timeout="20"/>	
		<action name="stop" timeout="20"/>
		<action name="monitor" timeout="5"/>
		<action name="meta-data" timeout="20"/>
		<action name="reload" timeout="20"/>
	</actions>
</resource-agent>`
