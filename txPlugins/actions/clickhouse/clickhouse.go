// Copyright (c) 2021 Damien Stuart. All rights reserved.
// Copyright (c) 2022 Kells Kearney. All rights reserved.
//
// Use of this source code is governed by the MIT License that can be found
// in the LICENSE file.
//
package main

/*
 * Dump data out in Clickhouse CSV data format, for adding to a Clickhouse database
 */

import (
	"fmt"
	"log"
	"os"
	"strings"

	pluginMeta "github.com/keruzu/trapmux/txPlugins"
	"github.com/rs/zerolog"

	"github.com/natefinch/lumberjack"
)

const pluginName = "Clickhouse"

type ClickhouseExport struct {
	logFile   string
	fd        *os.File
	logger    lumberjack.Logger
	logHandle *log.Logger
	isBroken  bool

	main_log *zerolog.Logger
}

// makeCsvLogger initializes and returns a lumberjack.Logger (logger with
// built-in log rotation management).
//
func makeCsvLogger(logfile string) *lumberjack.Logger {
	l := lumberjack.Logger{
		Filename: logfile,
	}
	return &l
}

func validateArguments(actionArgs map[string]string) error {
	validArgs := map[string]bool{"filename": true, "size_mb": true, "backups_max": true, "compress_after_rotate": true}

	for key, _ := range actionArgs {
		if _, ok := validArgs[key]; !ok {
			return fmt.Errorf("Unrecognized option to %s plugin: %s", pluginName, key)
		}
	}
	return nil
}

func (a *ClickhouseExport) Configure(pluginLog *zerolog.Logger, actionArgs map[string]string) error {
	if err := validateArguments(actionArgs); err != nil {
		return err
	}

	a.main_log = pluginLog

	a.logFile = actionArgs["filename"]
        // gosec complains if perms are less than 600
	fd, err := os.OpenFile(a.logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	a.fd = fd
	a.logHandle = log.New(fd, "", 0)
	a.logger = lumberjack.Logger{
		Filename: a.logFile,
	}
	a.logHandle.SetOutput(&a.logger)
	a.main_log.Info().Str("logfile", a.logFile).Msg("Added Clickhouse CSV log destination")

	return nil
}

func (a ClickhouseExport) ProcessTrap(trap *pluginMeta.Trap) error {
	logCsvTrap(trap, a.logHandle)
	return nil
}

func (a ClickhouseExport) SigUsr1() error {
	fmt.Println("SigUsr1")
	return nil
}

func (a ClickhouseExport) Close() error {
	return a.fd.Close()
}

func (a ClickhouseExport) SigUsr2() error {
	a.main_log.Info().Str("logfile", a.logFile).Msg("Rotating Clickhouse CSV file")
	return a.logger.Rotate()
}

// logCsvTrap takes care of logging the given trap to the given ClickhouseExport
// destination.
//
func logCsvTrap(trap *pluginMeta.Trap, l *log.Logger) {
	l.Printf(makeTrapLogCsvEntry(trap))
}

// makeTrapLogEntry creates a log entry string for the given trap data.
// Note that this particular implementation expects to be dealing with
// only v1 traps.
//
func makeTrapLogCsvEntry(trap *pluginMeta.Trap) string {
	var csv [11]string
	trapMap := trap.Trap2Map()

	csv[0] = trapMap["TrapDate"]
	csv[1] = trapMap["TrapTimestamp"]
	csv[2] = trapMap["TrapHost"]
	csv[3] = trapMap["TrapNumber"]
	csv[4] = trapMap["TrapSourceIP"]
	csv[5] = trapMap["TrapAgentAddress"]
	csv[6] = trapMap["TrapGenericType"]
	csv[7] = trapMap["TrapSpecificType"]
	csv[8] = trapMap["TrapEnterpriseOID"]

	// Varbinds are split to separate arrays - one for the ObjectIDs,
	// and the other for Values
	var vbObj []string
	var vbVal []string

	for key, value := range trapMap {
		if strings.HasPrefix(key, "Trap") {
			continue
		}
		vbObj = append(vbObj, key)
		vbVal = append(vbVal, value)
	}

	// Now we create the CS-escaped string representation of our varbind arrays
	// and add them to the CSV array.
	csv[9] = fmt.Sprintf("\"['%v']\"", strings.Join(vbObj, "','"))
	csv[10] = fmt.Sprintf("\"['%v']\"", strings.Join(vbVal, "','"))

	return strings.Join(csv[:], ",")
}

var ActionPlugin ClickhouseExport
