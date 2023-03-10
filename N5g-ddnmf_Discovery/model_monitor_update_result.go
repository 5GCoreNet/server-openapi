/*
 * N5g-ddnmf_Discovery API
 *
 * N5g-ddnmf_Discovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_N5g-ddnmf_Discovery

// MonitorUpdateResult - Represents the monitoring revocation results.
type MonitorUpdateResult struct {

	DiscType DiscoveryType `json:"discType"`

	// Contains the ProSe Restricted Code.
	ProseRestrictedCode string `json:"proseRestrictedCode"`

	// Contains the Application ID.
	AppId string `json:"appId"`

	// Contains the RPAUID.
	BannedRpauid string `json:"bannedRpauid"`

	// Contains the PDUID.
	BannedPduid string `json:"bannedPduid"`

	RevocationResult RevocationResult `json:"revocationResult"`
}
