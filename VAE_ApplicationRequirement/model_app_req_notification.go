/*
 * VAE_ApplicationRequirement
 *
 * API for VAE Application Requirement Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_VAE_ApplicationRequirement

// AppReqNotification - Represents a notificaton of the result of the network resource adaptation corresponding to the V2X application requirement. 
type AppReqNotification struct {

	// String providing an URI formatted according to RFC 3986.
	ResourceUri string `json:"resourceUri"`

	Result ReservationResult `json:"result"`
}
