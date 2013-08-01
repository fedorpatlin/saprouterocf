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

//must start as daemon, asynchronously
func saprouter_start() int {
	//check already running
	Ocf_log(OCF_DEBUG, "trying to start service")
	if st := saprouter_monitor(); st == OCF_SUCCESS {
		Ocf_log(OCF_INFO, "service already running")
		return OCF_SUCCESS
	}
	cmd := func() int { return run_service("start") }

	if err := ocf_daemon(cmd); err == OCF_SUCCESS {
		Ocf_log(OCF_DEBUG, "service executed")
		for {
			if err := saprouter_monitor(); err != OCF_SUCCESS {
				time.Sleep(1 * time.Second)
			} else {
				Ocf_log(OCF_DEBUG, "service up and running")
				return OCF_SUCCESS
			}
		}
	}
	return OCF_ERR_GENERIC
}

//returns OCF_SUCCESS if successfully stop process or process is not running
func saprouter_stop() int {
	if st := saprouter_monitor(); st != OCF_SUCCESS {
		return OCF_SUCCESS
	}
	if st := run_service("stop"); st != OCF_SUCCESS {
		Ocf_log(OCF_ERR, "Can't stop service")
		return OCF_ERR_GENERIC
	}
	Ocf_log(OCF_INFO, "Service stopped")
	return OCF_SUCCESS
}

func saprouter_reload() int {
	if st := saprouter_monitor(); st != OCF_SUCCESS {
		return OCF_NOT_RUNNING
	}
	if st := run_service("restart"); st == OCF_SUCCESS {
		Ocf_log(OCF_INFO, "service reloaded")
		return OCF_SUCCESS
	}
	return OCF_ERR_GENERIC
}

func saprouter_metadata() int {
	//	verify_all()
	fmt.Printf("%s\n", METADATA_XML)
	return OCF_SUCCESS
}

func saprouter_monitor() int {
	verify_all()
	Ocf_log(OCF_DEBUG, "geting service status")
	if st := run_service("status"); st == OCF_SUCCESS {
		Ocf_log(OCF_DEBUG, "rc is "+string(st))
		return OCF_SUCCESS
	} else {
		return OCF_NOT_RUNNING
	}
	return OCF_ERR_GENERIC
}

func dispatch() {
	cmd := os.Getenv("__OCF_ACTION")
	if cmd == "" {
		if os.Getenv("OCF_RESKEY_ocf-action") != "" {
			cmd = os.Getenv("OCF_RESKEY_ocf-action")
		} else {
			if len(os.Args) < 2 {
				Ocf_log(OCF_DEBUG, "saprouter action not set "+cmd)
				for _, v := range os.Environ() {
					Ocf_log(OCF_DEBUG, v)
				}
				saprouter_metadata()
				os.Exit(OCF_ERR_UNIMPLEMENTED)
			} else {
				cmd = os.Args[1]
			}
		}
	}
	Ocf_log(OCF_DEBUG, "saprouter operation: "+cmd)
	switch cmd {
	case "meta-data":
		os.Exit(saprouter_metadata())
	case "monitor":
		//		Ocf_log(OCF_ERR, "monitor action is running")
		os.Exit(saprouter_monitor())
	case "start":
		os.Stdout.Sync()
		os.Exit(saprouter_start())
	case "stop":
		os.Exit(saprouter_stop())
	case "reload":
		os.Exit(saprouter_reload())
	case "validate-all":
		os.Exit(verify_all())
	}
	os.Exit(OCF_ERR_UNIMPLEMENTED)
}

func verify_all() int {
	return Check_binary(get_param("binary"))
}

/** returns value of corresponding environment variable, or empty string
if parameter is required but not set function returns default value
*/
func get_param(name string) string {
	value := os.Getenv(fmt.Sprintf("OCF_RESKEY_%s", name))
	if value == "" {
		value = os.Getenv(fmt.Sprintf("OCF_RESKEY_%s_default", name))
	}
	return value
}

func set_param(name string, value string) {
	vname := fmt.Sprintf("OCF_RESKEY_%s", name)
	os.Setenv(vname, value)
}

func set_param_default(name string, value string) {
	defaultname := name + "_default"
	set_param(defaultname, value)
}

func init_me() {
	//setup defaults
	set_param_default("binary", SAPROUTER_BINARY)
	set_param_default("config", SAPROUTER_CONFIG)
	set_param_default("host", SAPROUTER_HOST)
	set_param_default("port", SAPROUTER_PORT)
	set_param_default("log", SAPROUTER_LOG)
	set_param_default("trace", SAPROUTER_TRACE)
}

func run_service(command string) int {
	Ocf_log(OCF_INFO, "current command is "+command)
	err := ocf_run(OCF_INFO, false, get_param("binary"), command)
	return err
}

func main() {
	//Ocf_log(OCF_DEBUG, "resource agent is running")
	init_me()
	dispatch()
}
