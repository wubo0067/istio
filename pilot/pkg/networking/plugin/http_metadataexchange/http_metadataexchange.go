package http_metadataexchange

import (
	wasm_proto "github.com/envoyproxy/go-control-plane/envoy/extensions/wasm/v3"
	"istio.io/istio/pilot/pkg/model"
	"istio.io/istio/pilot/pkg/networking"
	"istio.io/istio/pilot/pkg/networking/plugin"
	"istio.io/istio/pilot/pkg/telemetry/http_metadataexchange/factory"
	"istio.io/pkg/log"
)

var (
	http_metadataexchange_log = log.RegisterScope("http_metadataexchange", "http_metadataexhcange debugging", 0)
)

// Plugin implements Istio http metadata exchange telemetry filter
type Plugin struct{}

// NewPlugin returns an instance of the http metadata exchange plugin
func NewPlugin() plugin.Plugin {
	return Plugin{}
}

// OnOutboundListener is called whenever a new outbound listener is added to the LDS output for a given service.
// Can be used to add additional filters on the outbound path.
func (p Plugin) OnOutboundListener(in *plugin.InputParams, mutable *networking.MutableObjects) error {
	if in.Node.Type != model.Router {
		// Only care about router.
		return nil
	}

	return buildFilter(in, mutable, false)
}

// OnInboundListener is called whenever a new inbound listener is added to the LDS output for a given service.
// Can be used to add additional filters on the inbound path.
func (p Plugin) OnInboundListener(in *plugin.InputParams, mutable *networking.MutableObjects) error {
	if in.Node.Type != model.SidecarProxy {
		// Only care about sidecar.
		return nil
	}
	return buildFilter(in, mutable, false)
}

// OnInboundFilterChains is called whenever a plugin needs to setup the filter chains, including relevant filter chain
// configuration, like FilterChainMatch and TLSContext.
func (p Plugin) OnInboundFilterChains(in *plugin.InputParams) []networking.FilterChain {
	return nil
}

// OnInboundPassthrough is called whenever a new passthrough filter chain is added to the LDS output.
// Can be used to add additional filters.
func (p Plugin) OnInboundPassthrough(in *plugin.InputParams, mutable *networking.MutableObjects) error {
	return nil
}

// OnInboundPassthroughFilterChains is called whenever a plugin needs to setup custom pass through filter chain.
func (p Plugin) OnInboundPassthroughFilterChains(in *plugin.InputParams) []networking.FilterChain {
	return nil
}

func buildFilter(in *plugin.InputParams, mutable *networking.MutableObjects, isPassthrough bool) error {
	config := &wasm_proto.PluginConfig{}
	applier := factory.NewTelemetryFilterApplier(config)
	//endpointPort := uint32(0)
	var _ uint32
	if in.ServiceInstance != nil {
		_ = in.ServiceInstance.Endpoint.EndpointPort
	}

	for i := range mutable.FilterChains {
		if isPassthrough {
			// Get the real port from the filter chain match if this is generated for pass through filter chain.
			_ = mutable.FilterChains[i].FilterChainMatch.GetDestinationPort().GetValue()
		}
		if in.ListenerProtocol == networking.ListenerProtocolHTTP || mutable.FilterChains[i].ListenerProtocol == networking.ListenerProtocolHTTP {
			// Adding HTTP Metadataexchange filter
			if filter := applier.HTTPMedataexchangeFilter(); filter != nil {
				mutable.FilterChains[i].HTTP = append(mutable.FilterChains[i].HTTP, filter)
			}
		}
	}

	return nil
}
