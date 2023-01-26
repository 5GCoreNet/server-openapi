/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_PolicyAuthorization

// SpatialValidity - Describes explicitly the route to an Application location.
type SpatialValidity struct {

	// Defines the presence information provisioned by the AF. The praId attribute within the PresenceInfo data type is the key of the map.
	PresenceInfoList map[string]PresenceInfo `json:"presenceInfoList"`
}