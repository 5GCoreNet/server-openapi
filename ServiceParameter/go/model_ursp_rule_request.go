/*
 * 3gpp-service-parameter
 *
 * API for AF service paramter   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ServiceParameter

// UrspRuleRequest - Contains parameters that can be used to guide the URSP.
type UrspRuleRequest struct {

	TrafficDesc TrafficDescriptorComponents `json:"trafficDesc,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	RelatPrecedence int32 `json:"relatPrecedence,omitempty"`

	// Sets of parameters that may be used to guide the Route Selection Descriptors of the URSP.
	RouteSelParamSets []RouteSelectionParameterSet `json:"routeSelParamSets,omitempty"`
}
