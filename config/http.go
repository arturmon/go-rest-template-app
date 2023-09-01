package config

import "time"

// HTTP is the configuration of http server.
type HTTP struct {
	ListenAddr        string
	ReadHeaderTimeout time.Duration
}
