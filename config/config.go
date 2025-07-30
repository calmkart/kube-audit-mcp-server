package config

import (
	"flag"
)

var config Config

// Init initializes the configuration
func Init() {
	mode := flag.String("mode", "sse", "Transport type (stdio or sse)")
	host := flag.String("host", "127.0.0.1", "Host for SSE transport")
	port := flag.Int("port", 8080, "TCP port for SSE transport")
	provider := flag.String("provider", "aliyunsls", "Provider type (localfile or aliyunsls)")
	config = Config{
		Mode:     *mode,
		Host:     *host,
		Port:     *port,
		Provider: *provider,
	}
	flag.Parse()
}

// GetConfig returns the configuration
func GetConfig() Config {
	return config
}
