// Copyright (c) 2021 Damien Stuart. All rights reserved.
//
// Use of this source code is governed by the MIT License that can be found
// in the LICENSE file.
//
package main

import (
	"github.com/creasty/defaults"
	g "github.com/gosnmp/gosnmp"
	pluginLoader "github.com/keruzu/trapmux/api"
)

type trapListenerConfig struct {
	Hostname string `json:"hostname"`

	ListenAddr string `default:"0.0.0.0" json:"listen_address"`
	ListenPort string `default:"162" json:"listen_port"`

	GoSnmpDebug        bool   `default:"false" json:"gosnmp_debug"`
	GoSnmpDebugLogName string `default:"" json:"gosnmp_debug_logfile_name"`

	IgnoreVersions_str []string        `default:"[]" json:"ignore_versions"`
	IgnoreVersions     []g.SnmpVersion `default:"[]"`

	Community string `default:"" json:"snmp_community"`

	// SNMP v3 settings
	MsgFlags_str     string               `default:"NoAuthNoPriv" json:"msg_flags"`
	MsgFlags         g.SnmpV3MsgFlags     `default:"g.NoAuthNoPriv"`
	Username         string               `default:"XXv3Username" json:"username"`
	AuthProto_str    string               `default:"NoAuth" json:"auth_protocol"`
	AuthProto        g.SnmpV3AuthProtocol `default:"g.NoAuth"`
	AuthPassword     string               `default:"XXv3authPass" json:"auth_password"`
	PrivacyProto_str string               `default:"NoPriv" json:"privacy_protocol"`
	PrivacyProto     g.SnmpV3PrivProtocol `default:"g.NoPriv"`
	PrivacyPassword  string               `default:"XXv3Pass" json:"privacy_password"`
}

type IpSet map[string]bool

// filterObj represents one of the filterable items in a filter line from
// the config file (i.e. Src IP, AgentAddress, GenericType, SpecificType,
// and Enterprise OID).
//
type filterObj struct {
	filterItem  int
	filterType  int
	filterValue interface{} // string, *regex.Regexp, *network, int
}

// Get in a set of action arg pairs, convert to a map to pass into plugins

// trapmuxFilter holds the filter data and action for a specfic
// filter line from the config file.
type trapmuxFilter struct {
	// SnmpVersions - an empty array will indicate ALL versions
	SnmpVersions []string `default:"[]" json:"snmp_versions"`
	SourceIp     string   `default:"" json:"source_ip"`
	AgentAddress string   `default:"" json:"agent_address"`

	// GenericType can have values from 0 - 6: -1 indicates all types
	GenericType int `default:"-1" json:"snmp_generic_type"`
	// SpecificType can have values from 0 - n: -1 indicates all types
	SpecificType int `default:"-1" json:"snmp_specific_type"`

	EnterpriseOid string            `default:"" json:"enterprise_oid"`
	ActionName    string            `default:"" json:"action"`
	ActionArg     string            `default:"" json:"action_arg"`
	BreakAfter    bool              `default:"false" json:"break_after"`
	ActionArgs    map[string]string `default:"{}" json:"plugin_args"`

	// Compiled definition of above
	matchAll   bool
	matchers   []filterObj
	actionType int
	plugin     pluginLoader.ActionPlugin
}

type MetricConfig struct {
	PluginName string            `default:"" json:"plugin"`
	Args       map[string]string `default:"{}" json:"args"`
	plugin     pluginLoader.MetricPlugin
}

type trapmuxConfig struct {
	teConfigured bool

	General struct {
		PluginPath string `default:"txPlugins" json:"plugin_path"`
	}

	Reporting []MetricConfig `default:"[]" json:"metric_reporting"`

	Logging struct {
		Level         string `default:"debug" json:"level"`
		LogMaxSize    int    `default:"1024" json:"log_size_max"`
		LogMaxBackups int    `default:"7" json:"log_backups_max"`
		LogMaxAge     int    `json:"log_age_max"`
		LogCompress   bool   `default:"false" json:"compress_rotated_logs"`
	}

	TrapReceiverSettings trapListenerConfig `json:"listener"`

	IpSets_str []map[string][]string `default:"{}" json:"ip_sets"`
	IpSets     map[string]IpSet      `default:"{}"`

	Filters []trapmuxFilter `default:"[]" json:"filters"`

	// Bad things happen to good plugins. How do you want to handle exceptions?
	PluginErrorActions []trapmuxFilter `default:"[]" json:"plugin_error_actions"`
}
