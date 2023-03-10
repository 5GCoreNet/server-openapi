/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudr_DR

// UpuData2 - Contains UE parameters update data set (e.g., the updated Routing ID Data or the Default configured NSSAI).
type UpuData2 struct {

	// Contains a secure packet.
	SecPacket string `json:"secPacket,omitempty"`

	DefaultConfNssai []Snssai `json:"defaultConfNssai,omitempty"`

	// Represents a routing indicator.
	RoutingId string `json:"routingId,omitempty"`
}
