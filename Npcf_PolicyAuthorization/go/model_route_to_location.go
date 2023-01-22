/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Npcf_PolicyAuthorization

// RouteToLocation - At least one of the \"routeInfo\" attribute and the \"routeProfId\" attribute shall be included in the \"RouteToLocation\" data type. 
type RouteToLocation struct {

	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	Dnai string `json:"dnai"`

	RouteInfo *RouteInformation `json:"routeInfo,omitempty"`

	// Identifies the routing profile Id.
	RouteProfId *string `json:"routeProfId,omitempty"`
}
