/*
 * N5g-ddnmf_Discovery API
 *
 * N5g-ddnmf_Discovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_N5g-ddnmf_Discovery

// MonitorDiscDataForRestricted - Represents Data for restricted discovery used to request the authorization to monitor for a UE
type MonitorDiscDataForRestricted struct {

	// Contains the RPAUID.
	Rpauid string `json:"rpauid"`

	// Contains the PDUID.
	TargetPduid string `json:"targetPduid"`

	// Contains the Application ID.
	AppId string `json:"appId"`

	// Contains the RPAUID.
	TargetRpauid string `json:"targetRpauid"`
}