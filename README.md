# utilfx

[![CI](https://github.com/uchaloop/utilfx/actions/workflows/ci.yml/badge.svg)](https://github.com/uchaloop/utilfx/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/uchaloop/utilfx.svg)](https://pkg.go.dev/github.com/uchaloop/utilfx)
[![License: MIT](https://img.shields.io/github/license/uchaloop/utilfx)](LICENSE)

Focused helpers for the [Uber Fx](https://pkg.go.dev/go.uber.org/fx) wiring
patterns that come up in every application built on it.

- **An interface without an adapter** - `Bind` hands the container a concrete
  value under a contract, and says which types did not match when they do not.
- **Value groups without repetition** - several constructors into one group, as
  themselves or as an interface, nil entries ignored.
- **Tags that are not hand-quoted** - where a typo turns into a dependency
  nobody satisfies.

```bash
go get github.com/uchaloop/utilfx
```

## Quick start

```go
fx.Provide(
	makeService,
	utilfx.Bind[Handler, *Service],          // *Service is also still available
)

utilfx.GroupedAs[Handler](                   // both into one group, as Handler
	"http_handlers",
	makeHealthHandler,
	makeMetricsHandler,
)

fx.Annotate(
	makeClient,
	fx.ParamTags(utilfx.OptionalNameTag("analytics")),
	fx.ResultTags(utilfx.NameTag("analytics")),
)
```

A library exposing one entry point for its default and named instances:

```go
func moduleFor(name string) fx.Option {
	return utilfx.ModuleFor(name, Module, makeNamedModule)
}
```

## Documentation

What each helper does and the reasons behind it are in the package
documentation:
**[pkg.go.dev/github.com/uchaloop/utilfx](https://pkg.go.dev/github.com/uchaloop/utilfx)**.

## Acknowledgements

I am grateful to the authors of [Uber Fx](https://github.com/uber-go/fx). Their
work made this library possible.

## License

[MIT](LICENSE)
