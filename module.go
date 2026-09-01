package utilfx

import "go.uber.org/fx"

// ModuleFor returns the default module when name is empty. Otherwise it builds
// the named module.
func ModuleFor(name string, def fx.Option, mk func(string) fx.Option) fx.Option {
	if len(name) == 0 {
		return def
	}

	return mk(name)
}
