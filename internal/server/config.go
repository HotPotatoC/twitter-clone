package server

// Config is the configuration for the server
type Config struct {
	AppName string `json:"app_name"`
	Host    string `json:"host"`
	Port    string `json:"port"`
}

func (s *Config) init() {
	if s.Host == "" {
		s.Host = "127.0.0.1"
	}

	if s.Port == "" {
		s.Port = "5000"
	}
}
