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
	Hostname string `yaml:"hostname" json:"hostname"`

	ListenAddr string `default:"0.0.0.0" yaml:"listen_address" json:"listen_address"`
	ListenPort string `default:"162" yaml:"listen_port" json:"listen_port"`

	GoSnmpDebug        bool   `default:"false" yaml:"gosnmp_debug" json:"gosnmp_debug"`
	GoSnmpDebugLogName string `default:"" yaml:"gosnmp_debug_logfile_name" json:"gosnmp_debug_logfile_name"`

	IgnoreVersions_str []string        `default:"[]" yaml:"ignore_versions" json:"ignore_versions"`
	IgnoreVersions     []g.SnmpVersion `default:"[]"`

	Community string `default:"" yaml:"snmp_community"`

	// SNMP v3 settings
	MsgFlags_str     string               `default:"NoAuthNoPriv" yaml:"msg_flags" json:"msg_flags"`
	MsgFlags         g.SnmpV3MsgFlags     `default:"g.NoAuthNoPriv"`
	Username         string               `default:"XXv3Username" yaml:"username" json:"username"`
	AuthProto_str    string               `default:"NoAuth" yaml:"auth_protocol" json:"auth_protocol"`
	AuthProto        g.SnmpV3AuthProtocol `default:"g.NoAuth"`
	AuthPassword     string               `default:"XXv3authPass" yaml:"auth_password" json:"auth_password"`
	PrivacyProto_str string               `default:"NoPriv" yaml:"privacy_protocol" json:"privacy_protocol"`
	PrivacyProto     g.SnmpV3PrivProtocol `default:"g.NoPriv"`
	PrivacyPassword  string               `default:"XXv3Pass" yaml:"privacy_password" json:"privacy_password"`
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
	SnmpVersions []string `default:"[]" yaml:"snmp_versions" json:"snmp_versions"`
	SourceIp     string   `default:"" yaml:"source_ip" json:"source_ip"`
	AgentAddress string   `default:"" yaml:"agent_address" json:"agent_address"`

	// GenericType can have values from 0 - 6: -1 indicates all types
	GenericType int `default:"-1" yaml:"snmp_generic_type"`
	// SpecificType can have values from 0 - n: -1 indicates all types
	SpecificType int `default:"-1" yaml:"snmp_specific_type"`

	EnterpriseOid string            `default:"" yaml:"enterprise_oid" json:"enterprise_oid"`
	ActionName    string            `default:"" yaml:"action" json:"action"`
	ActionArg     string            `default:"" yaml:"action_arg" json:"action_arg"`
	BreakAfter    bool              `default:"false" yaml:"break_after" json:"break_after"`
	ActionArgs    map[string]string `default:"{}" yaml:"plugin_args" json:"plugin_args"`

	// Compiled definition of above
	matchAll   bool
	matchers   []filterObj
	actionType int
	plugin     pluginLoader.ActionPlugin
}

// UnmarshalYAML is what enables the setter to work for the trapmuxFilter
func (s *trapmuxFilter) UnmarshalYAML(unmarshal func(interface{}) error) error {
	configerr := defaults.Set(s)
	if configerr != nil {
		return configerr
	}

	type plain trapmuxFilter
	if err := unmarshal((*plain)(s)); err != nil {
		return err
	}
	return nil
}

type MetricConfig struct {
	PluginName string            `default:"" yaml:"plugin" json:"plugin"`
	Args       map[string]string `default:"{}" yaml:"args" json:"args"`
	plugin     pluginLoader.MetricPlugin
}

type trapmuxConfig struct {
	teConfigured bool

	General struct {
		PluginPath string `default:"txPlugins" yaml:"plugin_path" json:"plugin_path"`
	}

	Reporting []MetricConfig `default:"[]" yaml:"metric_reporting" json:"metric_reporting"`

	Logging struct {
		Level         string `default:"debug" yaml:"level" json:"level"`
		LogMaxSize    int    `default:"1024" yaml:"log_size_max" json:"log_size_max"`
		LogMaxBackups int    `default:"7" yaml:"log_backups_max" json:"log_backups_max"`
		LogMaxAge     int    `yaml:"log_age_max" json:"log_age_max"`
		LogCompress   bool   `default:"false" yaml:"compress_rotated_logs" json:"compress_rotated_logs"`
	}

	TrapReceiverSettings trapListenerConfig `yaml:"listener" json:"listener"`

	IpSets_str []map[string][]string `default:"{}" yaml:"ip_sets" json:"ip_sets"`
	IpSets     map[string]IpSet      `default:"{}"`

	Filters []trapmuxFilter `default:"[]" yaml:"filters" json:"filters"`

	// Bad things happen to good plugins. How do you want to handle exceptions?
	PluginErrorActions []trapmuxFilter `default:"[]" yaml:"plugin_error_actions" json:"plugin_error_actions"`
}
