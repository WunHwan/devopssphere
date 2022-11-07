package options

type ServerRunOptions struct {
	// server bind address
	BindAddress string

	InsecurePort int
}

func NewServerRunOptions() *ServerRunOptions {
	return &ServerRunOptions{
		BindAddress:  "0.0.0.0",
		InsecurePort: 8080,
	}
}
