package rscpgx

type opts struct {
	name                      string
	disableConnectionOnStatup bool
}

// Option is a function that can be used to configure a resource.
type Option func(r *opts)

func defaultOpts() opts {
	return opts{
		name: "PGX Connection",
	}
}

// WithName sets the name of the PGX resource.
func WithName(name string) Option {
	return func(r *opts) {
		r.name = name
	}
}

func WithDisableTestConnectionOnStartUp() Option {
	return func(r *opts) {
		r.disableConnectionOnStatup = true
	}
}
