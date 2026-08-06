# utilfx

[![CI](https://github.com/uchaloop/utilfx/actions/workflows/ci.yml/badge.svg)](https://github.com/uchaloop/utilfx/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/uchaloop/utilfx.svg)](https://pkg.go.dev/github.com/uchaloop/utilfx)
[![License: MIT](https://img.shields.io/badge/github/license/uchaloop/utilfx)](LICENSE)

Small helpers for common [Uber Fx](https://pkg.go.dev/go.uber.org/fx)
dependency-wiring patterns.

## Installation

```bash
go get github.com/uchaloop/utilfx
```

## Bind

Expose a concrete service through an interface:

```go
fx.Provide(
	makeService,
	utilfx.Bind[Handler, *Service],
)
```

The concrete service remains available in the graph. An incompatible binding
is reported as an Fx constructor error.

## Value groups

Register constructors in one group:

```go
utilfx.Grouped(
	"http_handlers",
	makeHealthHandler,
	makeMetricsHandler,
)
```

Register implementations of an interface:

```go
utilfx.GroupedAs[Handler](
	"http_handlers",
	makeHealthHandler,
	makeMetricsHandler,
)
```

Build a group name for a named instance:

```go
group := utilfx.GroupName("database_options", "analytics")

utilfx.GroupedFor(
	"analytics",
	"database_options",
	makeOption,
)
```

`Grouped`, `GroupedAs`, and `GroupedFor` ignore nil constructors.

## Fx tags

Build readable annotation tags:

```go
fx.Annotate(
	makeClient,
	fx.ParamTags(utilfx.OptionalNameTag("analytics")),
	fx.ResultTags(utilfx.NameTag("analytics")),
)
```

For value groups:

```go
fx.Annotate(
	makeHandler,
	fx.ResultTags(utilfx.GroupTag("http_handlers")),
)
```

Available helpers:

- `NameTag`
- `OptionalNameTag`
- `GroupTag`
- `GroupName`

## ModuleFor

Select the default module for an empty name and a named module otherwise:

```go
func moduleFor(name string) fx.Option {
	return utilfx.ModuleFor(name, Module, makeNamedModule)
}
```

## Acknowledgements

I am grateful to the authors of [Uber Fx](https://github.com/uber-go/fx). Their
work made this library possible.

## License

[MIT](LICENSE)
