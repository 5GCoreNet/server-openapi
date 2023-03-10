/*
 * 3GPP 5GC NRM
 *
 * OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_5GcNrm

// RouteToLocation - At least one of the \"routeInfo\" attribute and the \"routeProfId\" attribute shall be included in the \"RouteToLocation\" data type. 
type RouteToLocation struct {

	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	Dnai string `json:"dnai"`

	RouteInfo *RouteInformation `json:"routeInfo,omitempty"`

	// Identifies the routing profile Id.
	RouteProfId *string `json:"routeProfId,omitempty"`
}
