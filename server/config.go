package server

import "github.com/HotPotatoC/twitter-clone/internal/version"

// Config is the configuration for the server
type Config struct {
	AppName string `json:"app_name"`
	Version string `json:"version"`
	BuildID string `json:"build_id"`
	Host    string `json:"host"`
	Port    string `json:"port"`
}

func (s *Config) init() {
	if s.Version == "" {
		s.Version = version.Version
	}

	if s.BuildID == "" {
		s.BuildID = version.BuildID
	}

	if s.Host == "" {
		s.Host = "127.0.0.1"
	}

	if s.Port == "" {
		s.Port = "5000"
	}
}
