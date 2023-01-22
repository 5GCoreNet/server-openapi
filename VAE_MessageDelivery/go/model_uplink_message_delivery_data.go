/*
 * VAE_MessageDelivery
 *
 * API for VAE Message Delivery Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package VAE_MessageDelivery

// UplinkMessageDeliveryData - Represents the uplink V2X message delivery data.
type UplinkMessageDeliveryData struct {

	// String providing an URI formatted according to RFC 3986.
	ResourceUri string `json:"resourceUri"`

	// Represents the identifier of the V2X UE.
	UeId string `json:"ueId"`

	// Represents a geographical area identifier.
	GeoId string `json:"geoId,omitempty"`

	// string with format 'bytes' as defined in OpenAPI
	Payload string `json:"payload"`
}
