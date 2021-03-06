package main

import (
	"time"

	"github.com/baidu/openedge/protocol/mqtt"
	"github.com/baidu/openedge/utils"
)

// Config function service config
type Config struct {
	Rules     []RuleInfo     `yaml:"rules" json:"rules" default:"[]"`
	Functions []FunctionInfo `yaml:"functions" json:"functions" default:"[]"`
}

// RuleInfo function rule config
type RuleInfo struct {
	ClientID  string         `yaml:"clientid" json:"clientid"`
	Subscribe mqtt.TopicInfo `yaml:"subscribe" json:"subscribe"`
	Function  struct {
		Name string `yaml:"name" json:"name" validate:"nonzero"`
	} `yaml:"function" json:"function"`
	Publish mqtt.TopicInfo `yaml:"publish" json:"publish"`
}

// FunctionInfo function config
type FunctionInfo struct {
	Name     string       `yaml:"name" json:"name" validate:"nonzero"`
	Service  string       `yaml:"service" json:"service" validate:"nonzero"`
	Instance InstanceInfo `yaml:"instance" json:"instance"`
	Message  struct {
		Length utils.Length `yaml:"length" json:"length" default:"{\"max\":4194304}"`
	} `yaml:"message" json:"message"`
	Backoff struct {
		Max time.Duration `yaml:"max" json:"max" default:"1m"`
	} `yaml:"backoff" json:"backoff"`
	Timeout time.Duration `yaml:"timeout" json:"timeout" default:"30s"`
}

// InstanceInfo function instance info
type InstanceInfo struct {
	Min       int           `yaml:"min" json:"min" default:"0"`
	Max       int           `yaml:"max" json:"max" default:"1"`
	IdleTime  time.Duration `yaml:"idletime" json:"idletime" default:"10m"`
	EvictTime time.Duration `yaml:"evicttime" json:"evicttime" default:"1m"`
}
