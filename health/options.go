package health

import "connectrpc.com/connect"

type option struct {
	BaseURL       string
	ClientOptions []connect.ClientOption
}

// Option is a function that configures an option.
type Option func(*option)

// WithBaseURL sets the base URL for the health check.
func WithBaseURL(baseURL string) Option {
	return func(o *option) {
		o.BaseURL = baseURL
	}
}

// WithClientOptions sets the client options for the health check.
func WithClientOptions(clientOptions ...connect.ClientOption) Option {
	return func(o *option) {
		o.ClientOptions = clientOptions
	}
}
