# Trap MUX - A Program for receiving SNMP traps and sneding to different destinations
Trap MUX is an SNMP Trap proxy/forwarder.  It can receive, filter, manipulate, 
log, and forward SNMP traps to zero or mulitple destinations.  It can receive 
and process __SNMPv1__, __SNMPv2c__, or __SNMPv3__ traps.  

Presently, __v2c__ and __v3__ traps are converted to __v1__ before they are
logged and/or forwarded.  Support for sending other versions may be added in
a future release.

# Overview
Trap MUX* has the following features:

* Receive SNMP v1/v2c/v3 traps
* Selectively ignore traps based on their SNMP version.
* Filter traps based on one or more criteria:
  * Source IP address of the trap packet
  * Trap AgentAddress
  * Generic trap type
  * Specific trap type
  * Trap Enterprise OID
* Perform one of the followint actions based on a filter match:
  * Forward the trap
  * Change the AgentAddress value (_nat_ function)
  * Log the trap to a specified file
  * Log the trap data in a CSV format (specifically for feeding to a Clickhouse database).
  * Drop the trap and discontinue processing further filters.

This initial version's functionality and configuration options are very
similar to the current *eHealth Trap MUXploder* program, but there are some
differences, so be sure to carefully read this documentation if you will
be working with this program.

# Notes
This implementation uses external GO modules:

* [gosnmp](https://github.com/gosnmp/gosnmp)
(for all SNMP-based operations)
*  [lumberjack](https://github.com/natefinch/lumberjack)
(for logfile management). 

## Trap MUX Configuration
Trap MUX gets its configuration and runtime options from the configuration file.
Some options can be overridden by command-line arguments to _trap_mux_

Though the out-of-the box configuration file will have resonable defaults
and will have a filter directive for logging any traps, you should edit it
and set any configuration options as needed, and add/edit any other filter
directives. 

See the [trap_mux.conf file section](#markdown-header-the-trap_mux-configuration-file) below for details on the
configuration file and its directives

#### Running trap_mux
For a usage message for running *trap_mux*, run: `./trap_mux -h`
and you will get the following usage message:

```
Usage: trap_mux [-h] [-c <config_file>] [-b <bind_ip>] [-p <listen_port>]
              [-d] [-v]
  -h  - Show this help message and exit.
  -c  - Override the location of the trap_mux configuration file.
  -b  - Override the bind IP address on which to listen for incoming traps.
  -p  - Override the UDP port on which to listen for incoming traps.
  -d  - Enable debug mode (note: produces very verbose runtime output).
  -v  - Print the version of trap_mux and exit.
```

*trap_mux* will stay in the foreground and print information
to STDOUT.  On startup, any *filter* directives that forward a trap will
be printed as they are loaded from the configuration file. Here is an example:

```
Loading configuration for trap_mux version 0.9.2 from: /opt/trap_mux/etc/trap_mux.conf.
 -Added trap destination: 216.203.5.174, port 162
 -Added trap destination: 10.131.125.61, port 162
 -Added trap destination: 10.131.125.62, port 162
 -Added trap destination: 69.164.113.76, port 162
 -Added log destination: /opt/trap_mux/log/trap_mux.log
```

Also, any actions triggered by a signal will cause output to be printed to STDOUT as well.

#### Signals
*Trap MUX* has handlers for the following signals:

* *SIGHUP*

  Sending a SIGHUP to the *trap_mux* process ID (PID) will cause it to re-read
  the configuration file.  This is useful if you made a change in the
  configuration file and want to start using it without having to stop/start
  the service.

* *SIGUSR1*

  Sending a SIGUSR1 signal will cause *trap_mux* to dump its current trap
  stats (counters and trap rates) to STDOUT (or the trap_mux-run.log file).
  Here is a sample of one of these stat dumps:  

```
  Got SIGUSR1
  Trap MUX stats as of: Fri, 10 Jul 2020 08:35:30 EDT
   - Uptime..............: 0d-15h-41m-0s
   - Traps Received......: 2744275
   - Traps Processed.....: 2744275
   - Traps Dropped.......: 879493
   - Translated from v2c.: 288232
   - Translated from v3..: 497
   - Trap Rates (based on all traps received):
      - Last Minute......: 48
      - Last 5 Minutes...: 50
      - Last 15 Minutes..: 50
      - Last Hour........: 49
      - Last 4 Hours.....: 49
      - Last 8 Hours.....: 49
      - Last 24 Hours....: 32
      - Since Start......: 49
```

The trap rates above are calculated based on all incoming traps before any
filtering or processing takes place.  If the *ignoreVersion* option was used,
there would be a counter for *Traps Ignored* in the list.  Also, if `v2c`
and/or `v3` versions are ignored, the *Translated from vX* line will not be
included in the list.
 
Note that the *Traps Dropped* counter indicates traps that were matched on a
filter that has the `break` action.  Depending on the position of the filter
entry it matched, the trap may or may not have been forwarded or logged. This
is just an indicator that the trap did not traverse the entire filter list. 

* *SIGUSR2*

  Sending a SIGUSR2 signal will cause *trap_mux* to force a rotation of any
  configured CSV logs.  Since the CSV logs are meant to be used for feeding
  trap data to a database, this mechanism allows for doing a rotation
  on-demand so data can be synced to the database on a schedule.  


# Author

Trapex was writen by Damien Stuart - (<dstuart@dstuart.org>)
Trap_mux was forked by Kells Kearney and subsequently maintained
