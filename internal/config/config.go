package config

import (
	cfg "github.com/335is/config"
)

// Config holds our configuration settings
//	OCTOREPORT_OCTOPUS_ADDRESS
//	OCTOREPORT_OCTOPUS_APIKEY
type Config struct {
	Octopus *Octopus `yaml:"octopus"`
}

// Octopus holds our octopus settings
type Octopus struct {
	Address string `yaml:"address" default:"https://demo.octopus.com"`
	APIKey  string `yaml:"apikey" default:"API-GUEST"`
}

// New starts with a default config that works with a public demo Octopus Deploy server,
// then overrides settings from a YAML config file and env vars if they exist.
func New() *Config {
	c := Config{
		Octopus: &Octopus{},
	}

	cfg.Load("", &c)

	return &c
}

// Dump returns the Octopus configuration in YAML string form
func (c *Config) Dump() string {
	s, _ := cfg.ToYaml(c)
	return s
}
