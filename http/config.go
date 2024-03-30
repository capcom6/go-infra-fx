package http

var ConfigDefault = Config{
	Listen:  ":3000",
	Proxies: []string{},
}

type Config struct {
	Listen  string
	Proxies []string
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

	return config
}
