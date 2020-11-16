package v1beta1

import (
	"istio.io/istio/pilot/pkg/networking/util"
	"istio.io/istio/pilot/pkg/telemetry/http_metadataexchange"
	"istio.io/istio/pilot/pkg/telemetry/model"

	wasm "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/wasm/v3"
	//"istio.io/istio/pilot/pkg/telemetry/model"
	wasm_proto "github.com/envoyproxy/go-control-plane/envoy/extensions/wasm/v3"
	//envoy_jwt "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/jwt_authn/v3"
	http_conn "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	"istio.io/pkg/log"
)

var (
	http_metadataexchange_log = log.RegisterScope("http_metadataexchange", "http_metadataexhcange debugging", 0)
)

// Implemenation of http_metadataexchange.telemetryfilter_applier with v1beta1 API.
type v1beta1TelemetryFilterApplier struct {
	http_metadataexchange_wasm_config *wasm_proto.PluginConfig
}

// NewPolicyApplier returns new applier for v1beta1 authentication policies.
func NewTelemetryFilterApplier(wasm *wasm_proto.PluginConfig) http_metadataexchange.TelemetryFilterApplier {
	http_metadataexchange_wasm_config_current := wasm

	//wasm_vm_config := wasm_proto.VmConfig{
	//	VmId: "fake_id",
	//	Runtime: "envoy.wasm.runtime.null",
	//	Code: nil, // v3.AsyncDataSource
	//	Configuration: nil, //
	//	AllowPrecompiled: true, //???
	//	NackOnCodeCacheMiss: true, //???
	//}
	return &v1beta1TelemetryFilterApplier{
		http_metadataexchange_wasm_config: http_metadataexchange_wasm_config_current,
	}
}


func (a *v1beta1TelemetryFilterApplier) HTTPMedataexchangeFilter() *http_conn.HttpFilter {
	if a.http_metadataexchange_wasm_config == nil {
		return nil
	}

	filterConfigProto := convertToEnvoyWasmConfig(a.http_metadataexchange_wasm_config)

	if filterConfigProto == nil {
		return nil
	}
	return &http_conn.HttpFilter{
		Name:      model.HTTPMetadataExchangeFilterName,
		ConfigType: &http_conn.HttpFilter_TypedConfig{TypedConfig: util.MessageToAny(filterConfigProto)},
	}
}


func convertToEnvoyWasmConfig(http_metadataexchange_wasm_config *wasm_proto.PluginConfig) *wasm.Wasm {
	if http_metadataexchange_wasm_config == nil {
		return nil
	}
	//wasm_wasms := map[string]*wasm.Wasm{}

	//for i, wasm_config := range http_metadataexchange_wasm_config {
	//	wasm_wasm := &wasm.Wasm{
	//		// General Plugin configuration.
	//		Config: wasm_config,
	//	}
	//
	//	name := fmt.Sprintf("origins-%d", i)
	//	wasm_wasms[name] = wasm_wasm
	//}
	return &wasm.Wasm{
		Config: http_metadataexchange_wasm_config,
	}
}
