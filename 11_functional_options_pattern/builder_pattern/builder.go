package builder

import "errors"

type config struct {
	Port int
}
type ConfigBuilder struct {
	port *int
}

func (b *ConfigBuilder) Port(
	port int) *ConfigBuilder {
	b.port = &port
	return b
}

var defaultHTTPPort = 8080

func randomPort() int {
	return 8081
}

func (b *ConfigBuilder) Build() (config, error) {
	cfg := config{}
	if b.port == nil {
		cfg.Port = defaultHTTPPort
	} else {
		if *b.port == 0 {
			cfg.Port = randomPort()
		} else if *b.port < 0 {
			return config{}, errors.New("port should be positive")
		} else {
			cfg.Port = *b.port
		}
	}
	return cfg, nil
}

func NewServer(addr string, cfg config) {
	println("builder new server")
}
