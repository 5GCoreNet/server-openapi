/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnrf_NFDiscovery

// NfInstanceInfo - Contains information on an NF profile matching a discovery request
type NfInstanceInfo struct {

	// String providing an URI formatted according to RFC 3986.
	NrfDiscApiUri string `json:"nrfDiscApiUri,omitempty"`

	PreferredSearch PreferredSearch `json:"preferredSearch,omitempty"`

	// The key of the map is the JSON Pointer of the priority IE in the NFProfile data type that is altered by the NRF 
	NrfAlteredPriorities map[string]int32 `json:"nrfAlteredPriorities,omitempty"`
}
