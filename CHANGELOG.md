# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2026-07-21

The first release of `actions-demo`, a small Go command-line tool that exists to
give a GitHub Actions pipeline something real to lint, test and build.

### Added

- An `actions-demo` CLI covering nine subcommands; running it with no arguments
  prints the usage text.
- Arithmetic commands `add`, `sub`, `mul` and `div` operate on two integers,
  whilst `sum` totals an arbitrary list of them. Division by zero is reported as
  an error rather than crashing.
- Text commands `reverse`, `palindrome`, `words` and `title` operate on the
  remainder of the command line joined together, so quoting is optional.
- A CI workflow that lints the code, runs the tests against both Go 1.25 and the
  current stable release with the race detector and coverage enabled, and then
  builds the binary.

[1.0.0]: https://github.com/dpage/actions-demo/releases/tag/v1.0.0
