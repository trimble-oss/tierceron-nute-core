module github.com/trimble-oss/tierceron-nute-core

go 1.24.0

require (
	google.golang.org/grpc v1.79.3
	google.golang.org/protobuf v1.36.10
)

require (
	golang.org/x/net v0.49.0 // indirect
	golang.org/x/sys v0.41.0 // indirect
	golang.org/x/text v0.34.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251202230838-ff82c1b0f217 // indirect
)

replace (
	go.opentelemetry.io/otel => go.opentelemetry.io/otel v1.40.0
	go.opentelemetry.io/otel/metric => go.opentelemetry.io/otel/metric v1.40.0
	go.opentelemetry.io/otel/sdk => go.opentelemetry.io/otel/sdk v1.40.0
	go.opentelemetry.io/otel/sdk/metric => go.opentelemetry.io/otel/sdk/metric v1.40.0
	go.opentelemetry.io/otel/trace => go.opentelemetry.io/otel/trace v1.40.0
)
