package http

import "time"

type Config struct {
	Listen  string
	Proxies []string

	WriteTimeout time.Duration
}

var ConfigDefault = Config{
	Listen:  ":3000",
	Proxies: []string{},

	WriteTimeout: 5 * time.Second,
}

// Helper function to set default values
func configDefault(config Config) Config {
	// Override default config
	if config.Listen == "" {
		config.Listen = ConfigDefault.Listen
	}

	if len(config.Proxies) == 0 {
		config.Proxies = ConfigDefault.Proxies
	}

	if config.WriteTimeout == 0 {
		config.WriteTimeout = ConfigDefault.WriteTimeout
	} else if config.WriteTimeout < 0 {
		config.WriteTimeout = 0
	}

	return config
}
