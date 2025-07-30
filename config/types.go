package config

type Config struct {
	Mode     string
	Host     string
	Port     int
	Provider string
}

const (
	ModeStdio = "stdio"
	ModeSSE   = "sse"
)
