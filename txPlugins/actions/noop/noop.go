// Copyright (c) 2022 Kells Kearney. All rights reserved.
//
// Use of this source code is governed by the MIT License that can be found
// in the LICENSE file.
//
package main

/*
This plugin performs no useful action, but can be used for unit testing or performance
testing purposes.
*/

import (
	pluginMeta "github.com/keruzu/trapex/txPlugins"
	"github.com/rs/zerolog"
)

type noopFilter struct {
	pluginLog *zerolog.Logger
}

const plugin_name = "noop"

func (p *noopFilter) Configure(pluginLog *zerolog.Logger, actionArgs map[string]string) error {
	pluginLog.Info().Str("plugin", plugin_name).Msg("Initialization of plugin")
	p.pluginLog = pluginLog
	return nil
}

func (p noopFilter) ProcessTrap(trap *pluginMeta.Trap) error {
	p.pluginLog.Info().Str("plugin", plugin_name).Msg("Noop processing trap")
	return nil
}

func (p noopFilter) Close() error {
	return nil
}

func (p noopFilter) SigUsr1() error {
	return nil
}

func (p noopFilter) SigUsr2() error {
	return nil
}

// Exported symbol which supports filter.go's FilterAction type
var ActionPlugin noopFilter
