package main

//defaults
const (
	SAPROUTER_BINARY = "saprouter"
	SAPROUTER_CONFIG = "/etc/saprouter/saprouttab"
	SAPROUTER_HOST   = "localhost"
	SAPROUTER_PORT   = "3299"
	SAPROUTER_LOG    = "/var/log/saprouter.log"
	SAPROUTER_TRACE  = "/var/log/saprouter.trc"
)

/*saprouter options
-R routtab   : name of route-permission-file  (default ./saprouttab)
-G logfile   : name of log file               (default no logging)
-T tracefile : name of trace file             (default dev_rout)
-V tracelev  : trace level to run with        (default 1)
-H hostname  : of running SAProuter           (default localhost)
-S service   : service-name / number          (default 3299)
-P infopass  : password for info requests
-C clients   : maximum no of clients          (default 800)
-Y servers   : maximum no of servers to start (default 1)
-K [myname]  : activate SNC; if given, use 'myname' as own sec-id
-A initstring: initialization options for third-party library
-D           : switch DNS reverse lookup off
-E           : append log- and trace-files to existing
-J filesize  : maximum log file size in byte  (default off)
-6           : IPv6 enabled
-Z           : hide connect error information for clients
*/
const (
	OPTION_ROUTTAB   = "-R"
	OPTION_LOG       = "-G"
	OPTION_TRACE     = "-T"
	OPTION_NODNS     = "-D"
	OPTION_NOUSERTRC = "-Z"
)

/*saprouter commands
*start router : saprouter -r
*stop router  : saprouter -s
soft shutdown: saprouter -p
*router info  : saprouter -l (-L)
*new routtab  : saprouter -n
toggle trace : saprouter -t
cancel route : saprouter -c id
dump buffers : saprouter -d
flush   "    : saprouter -f
hide errInfo : saprouter -z
*/
const (
	CMD_RUN    = "-r"
	CMD_STOP   = "-s"
	CMD_STATUS = "-l"
)

//API return codes
const (
	OCF_SUCCESS           = 0
	OCF_ERR_GENERIC       = 1
	OCF_ERR_ARGS          = 2
	OCF_ERR_UNIMPLEMENTED = 3
	OCF_ERR_PERM          = 4
	OCF_ERR_INSTALLED     = 5
	OCF_ERR_CONFIGURED    = 6
	OCF_NOT_RUNNING       = 7
	OCF_RUNNING_MASTER    = 8
	OCF_FAILED_MASTER     = 9
)

//log message categories
const (
	OCF_DEBUG = "debug"
	OCF_INFO  = "info"
	OCF_WARN  = "warn"
	OCF_ERR   = "err"
	OCF_CRIT  = "crit"
)

const METADATA_XML = `<?xml version="1.0"?>
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
			<content type="string" default="3299"/>
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
