## License
[LICENSE](LICENSE)

# Tierceron-nute-core

[![GitHub release](https://img.shields.io/github/release/trimble-oss/tierceron-nute-core.svg?style=flat-square)](https://github.com/trimble-oss/tierceron-nute-core/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/trimble-oss/tierceron-nute-core)](https://goreportcard.com/report/github.com/trimble-oss/tierceron-nute-core)
[![PkgGoDev](https://img.shields.io/badge/go.dev-docs-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/trimble-oss/tierceron-nute-core)

## What is it?
Tierceron-nute-core publishes the shared `mashupsdk` package without the heavier GUI and example dependencies from tierceron-nute. It is the lightweight import path for protobuf, gRPC, display-state, certificate, and helper types used by mashup clients and servers.

## What is in this repo?
- `mashupsdk`: generated protobuf and gRPC types, mashup interfaces, display state and context types, certificate helpers, and GUI-neutral shared helpers.

## Key Features
- A focused module for consumers that only need the shared mashup SDK.
- Generated protobuf and gRPC assets kept separate from the larger UI and example repositories.
- Shared display and interaction types used by higher-level mashup front ends.

## Getting started
To work with the module locally:

- Run `go mod download`.
- Install `protoc` if you need to regenerate gRPC assets: https://grpc.io/docs/protoc-installation/
- Run `make mashupsdk` to regenerate the protobuf and gRPC files.
- Run `go test ./...` to validate the module.

## Trusted Committers
- [Joel Rieke](mailto:joel_rieke@trimble.com)
- [David Mkrtychyan](mailto:david_mkrtychyan@trimble.com)
- [Karnveer Gill](mailto:karnveer_gill@trimble.com)
- [Meghan Bailey](mailto:meghan_bailey@trimble.com)

## Contributing
Contributions are always welcome. Before contributing, please read the [code of conduct](CODE_OF_CONDUCT.MD).

See [Contributing](CONTRIBUTING.MD).

## Security
Please review [SECURITY.md](SECURITY.md) for vulnerability reporting guidance.

