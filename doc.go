/*
Package utilfx contains focused helpers for Uber Fx dependency wiring: the few
patterns that come up in every application built on it, each of which is
otherwise written as an adapter function or a hand-quoted struct tag.

# Exposing a service through an interface

Bind hands the container a concrete value under an interface, replacing the
adapter every such wiring otherwise needs:

	fx.Provide(
		makeService,
		utilfx.Bind[Handler, *Service],
	)

The concrete service stays in the graph, so anything that wants *Service still
gets it. An implementation that does not satisfy the contract is reported as an
Fx constructor error naming both types, rather than failing to compile in an
adapter nobody wrote.

# Value groups

Grouped provides several constructors into one Fx value group:

	utilfx.Grouped("http_handlers", makeHealthHandler, makeMetricsHandler)

GroupedAs does the same for constructors whose results are to be consumed
through an interface; each result must implement T. GroupedFor composes the
group name from a group and an instance name, for an application that wires the
same group per named instance:

	utilfx.GroupedFor("database_options", "analytics", makeOption)

All three take the group first, and all three ignore a nil constructor, so a
list assembled conditionally does not have to be compacted first.

# Tags

An Fx tag is a quoted struct tag, and quoting one by hand is where a typo turns
into a dependency nobody satisfies. NameTag, OptionalNameTag and GroupTag build
them:

	fx.Annotate(
		makeClient,
		fx.ParamTags(utilfx.OptionalNameTag("analytics")),
		fx.ResultTags(utilfx.NameTag("analytics")),
	)

GroupName builds the group a named instance feeds - the group itself when the
name is empty, and "group:name" otherwise. It is the same rule GroupedFor
applies, exposed for a module that needs the name without providing anything.

# Default and named modules

ModuleFor picks between them by the instance name, so a library exposes one
entry point for both:

	func moduleFor(name string) fx.Option {
		return utilfx.ModuleFor(name, Module, makeNamedModule)
	}

An empty name selects the default module - the single instance an application
usually has - and any other name builds the named one.
*/
package utilfx
