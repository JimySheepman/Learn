package authenticator

import (
	"time"
)

type options struct {
	jwtDuration          time.Duration
	allowUnverifiedEmail bool
	allowUnverifiedPhone bool
}

type Option func(*options)

func WithCustomJwtDuration(duration time.Duration) Option {
	return func(options *options) {
		options.jwtDuration = duration
	}
}

func UnverifiedEmailAllowed() Option {
	return func(options *options) {
		options.allowUnverifiedEmail = true
	}
}

func UnverifiedPhoneAllowed() Option {
	return func(options *options) {
		options.allowUnverifiedEmail = true
	}
}

type Authenticator struct {
	name    string
	options options
}

func New(name string, opts ...Option) *Authenticator {
	options := options{
		jwtDuration: time.Hour,
	}
	for _, opt := range opts {
		opt(&options)
	}
	return &Authenticator{name: name, options: options}
}

func (a Authenticator) IsValidJWT(jwt string) bool {
	return false
}
