// Copyright (c) 2022 Kells Kearney. All rights reserved.
//
// Use of this source code is governed by the MIT License that can be found
// in the LICENSE file.
//
package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"path/filepath"

	pluginMeta "github.com/keruzu/trapmux/txPlugins"
	pluginLoader "github.com/keruzu/trapmux/api"

	"github.com/rs/zerolog"
)

var replayLog = zerolog.New(os.Stdout).With().Timestamp().Logger()

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	flag.Usage = func() {
		fmt.Printf("Usage:\n")
		fmt.Printf("   %s\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}

	processCommandLine()
	if err := getConfig(); err != nil {
		replayLog.Fatal().Err(err).Msg("Unable to load configuration")
		os.Exit(1)
	}
	var count int

	if teCmdLine.isFile {
		replayTrap(teCmdLine.filenames)
	} else {
		files, err := ioutil.ReadDir(teCmdLine.filenames)
		if err != nil {
			replayLog.Fatal().Err(err).Str("dir", teCmdLine.filenames).Msg("Unable to process capture file directory")
		}

		for _, fd := range files {
			count++
			filename := fd.Name()
			if strings.HasSuffix(filename, ".gob") {
				replayTrap(filename)
			}

		}
	}

	replayLog.Info().Int("replayed_traps", count).Msg("Replayed traps")
	/*
	   	startTime := time.Now()
	   	endTime := time.Now()
	           duration := endTime - startTime
	   		replayLog.Info().Int("replay_duration", duration).Msg("Replayed trap in %v seconds")

	*/
}

// replayTrap reads a file from disk and processes the trap accordingly.
//
func replayTrap(filename string) {
	trap, err := loadCaptureGob(filename)
	if err != nil {
		replayLog.Fatal().Err(err).Str("format", "gob").Str("capture_file", filename).Msg("Unable to load capture file")
		os.Exit(1)
	}

	for _, destination := range teConfig.Destinations {
		err = destination.plugin.(pluginLoader.ActionPlugin).ProcessTrap(&trap)
		replayLog.Warn().Err(err).Str("capture_file", filename).Msg("Unable to replay file")
	}
}

func loadCaptureGob(filename string) (pluginMeta.Trap, error) {
	var trap pluginMeta.Trap
	fd, err := os.Open(filepath.Clean(filename))
	if err != nil {
		return trap, err
	}
defer func() {
    if err := fd.Close(); err != nil { 
		replayLog.Error().Err(err).Str("capture_file", filename).Msg("Unable to load capture file")
    }
}()

	decoder := gob.NewDecoder(fd)
	err = decoder.Decode(&trap)
	return trap, err
}
