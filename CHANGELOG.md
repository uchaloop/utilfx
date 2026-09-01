# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.2.0] - 2026-09-01

### Changed

- `GroupedFor` takes the group before the name, the order `Grouped`, `GroupedAs`
  and `GroupName` already use. Both arguments are strings, so a call written from
  the habit of the others compiled and fed a group nobody consumes. Swap the
  arguments at any call site.
- The package documentation carries what each helper does and why; the README is
  a landing page. The package comment moved from `module.go` into `doc.go`.
- The module is built with Go 1.27. A module that depends on this one has to
  declare 1.27 as well.

## [0.1.1] - 2026-08-06

### Changed

- Reworked the README as concise, user-focused documentation.

## [0.1.0] - 2026-08-04

### Added
- Helpers for binding implementations, grouping constructors, selecting named
  modules, and building Fx tags.

[Unreleased]: https://github.com/uchaloop/utilfx/compare/v0.1.1...HEAD
[0.1.1]: https://github.com/uchaloop/utilfx/compare/v0.1.0...v0.1.1
[0.1.0]: https://github.com/uchaloop/utilfx/releases/tag/v0.1.0
