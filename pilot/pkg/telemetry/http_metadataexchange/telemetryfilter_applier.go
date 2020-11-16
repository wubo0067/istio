package http_metadataexchange

import (
	//wasm "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/wasm/v3"
	http_conn "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"

)

// TelemetryFilterApplier is the interface provides essential functionalities to help config Envoy (xDS) to enforce
// telemetry filters. Each kind of Istio telemetry filter needs to implement this interface.
type TelemetryFilterApplier interface {
	// HTTPMedataexchangeFilter returns the HTTP Metadata Exchange filter to enforce the underlying Istio telemetry filter.
	HTTPMedataexchangeFilter()  *http_conn.HttpFilter
}
