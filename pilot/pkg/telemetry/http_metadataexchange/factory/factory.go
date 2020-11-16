package factory

import (
	"istio.io/istio/pilot/pkg/telemetry/http_metadataexchange"
	"istio.io/istio/pilot/pkg/telemetry/http_metadataexchange/v1beta1"
	wasm_proto "github.com/envoyproxy/go-control-plane/envoy/extensions/wasm/v3"
)

// NewPolicyApplier returns the appropriate (policy) applier, depends on the versions of the policy exists
// for the given service instance.
func NewTelemetryFilterApplier(config *wasm_proto.PluginConfig) http_metadataexchange.TelemetryFilterApplier {
	return v1beta1.NewTelemetryFilterApplier(config)
}
