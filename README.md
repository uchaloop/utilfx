# utilfx

[![CI](https://github.com/uchaloop/utilfx/actions/workflows/ci.yml/badge.svg)](https://github.com/uchaloop/utilfx/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/uchaloop/utilfx.svg)](https://pkg.go.dev/github.com/uchaloop/utilfx)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

Small helpers for recurring [Uber Fx](https://pkg.go.dev/go.uber.org/fx)
dependency wiring.

## Install

```bash
go get github.com/uchaloop/utilfx@latest
```

## Bind

`Bind` exposes an existing concrete service under an interface while preserving
the concrete dependency in the graph.

Without `utilfx`:

```go
func makeHandler(service *Service) Handler {
	return service
}

fx.Provide(
	makeService,
	makeHandler,
)
```

With `utilfx`:

```go
fx.Provide(
	makeService,
	utilfx.Bind[Handler, *Service],
)
```

An incompatible binding is returned as an Fx constructor error.

## Grouped

`Grouped` registers one or more constructors in the same value group. Nil
constructors are ignored.

Without `utilfx`:

```go
fx.Provide(
	fx.Annotate(makeHealthHandler, fx.ResultTags(`group:"http_handlers"`)),
	fx.Annotate(makeMetricsHandler, fx.ResultTags(`group:"http_handlers"`)),
)
```

With `utilfx`:

```go
utilfx.Grouped("http_handlers", makeHealthHandler, makeMetricsHandler)
```

## GroupedAs

`GroupedAs` provides multiple implementations of one interface in a value
group.

Without `utilfx`:

```go
fx.Provide(
	fx.Annotate(
		makeHealthHandler,
		fx.As(new(Handler)),
		fx.ResultTags(`group:"http_handlers"`),
	),
	fx.Annotate(
		makeMetricsHandler,
		fx.As(new(Handler)),
		fx.ResultTags(`group:"http_handlers"`),
	),
)
```

With `utilfx`:

```go
utilfx.GroupedAs[Handler](
	"http_handlers",
	makeHealthHandler,
	makeMetricsHandler,
)
```

## GroupedFor

`GroupName` builds a group from a group name and an optional Fx name.
`GroupedFor` uses that group to register constructors.

Without `utilfx`:

```go
group := "database_options"
if len(name) > 0 {
	group += ":" + name
}

option := fx.Provide(
	fx.Annotate(makeOption, fx.ResultTags(`group:"`+group+`"`)),
)
```

With `utilfx`:

```go
group := utilfx.GroupName("database_options", name)
option := utilfx.GroupedFor(name, "database_options", makeOption)
```

## Tags

`NameTag`, `OptionalNameTag`, and `GroupTag` keep manual Fx annotations
readable.

Without `utilfx`:

```go
fx.Annotate(
	makeClient,
	fx.ParamTags(`name:"analytics" optional:"true"`),
	fx.ResultTags(`name:"analytics"`),
)
```

With `utilfx`:

```go
fx.Annotate(
	makeClient,
	fx.ParamTags(utilfx.OptionalNameTag("analytics")),
	fx.ResultTags(utilfx.NameTag("analytics")),
)
```

For a value group, without `utilfx`:

```go
fx.Annotate(
	makeHandler,
	fx.ResultTags(`group:"http_handlers"`),
)
```

With `GroupTag`:

```go
fx.Annotate(
	makeHandler,
	fx.ResultTags(utilfx.GroupTag("http_handlers")),
)
```

## ModuleFor

`ModuleFor` selects a default module for an empty name and otherwise builds a
named variant.

Without `utilfx`:

```go
func moduleFor(name string) fx.Option {
	if len(name) == 0 {
		return Module
	}

	return makeNamedModule(name)
}
```

With `utilfx`:

```go
func moduleFor(name string) fx.Option {
	return utilfx.ModuleFor(name, Module, makeNamedModule)
}
```

## Acknowledgements

Thanks to the creators and maintainers of [Uber Fx](https://github.com/uber-go/fx)
for building the dependency injection framework that makes this package useful.
