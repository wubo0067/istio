package model

const (

	// EnvoyJwtFilterName is the name of the Envoy JWT filter. This should be the same as the name defined
	// in https://github.com/envoyproxy/envoy/blob/v1.9.1/source/extensions/filters/http/well_known_names.h#L48
	//EnvoyJwtFilterName = "envoy.filters.http.jwt_authn"

	//todo: could not find a name for it????
	// Envoy
	//EnvoyHTTPMetadataExchangeFilterName = "envoy.filters.network.metadata_exchange"
	//TO BE CLEAR: this part is istio's metadata exchange not envoy
	HTTPMetadataExchangeFilterName = "istio.metdata_exchange"
)

