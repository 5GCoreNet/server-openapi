/*
 * N5g-ddnmf_Discovery API
 *
 * N5g-ddnmf_Discovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package N5g-ddnmf_Discovery

import (
	"time"
)

// AnnounceDiscDataForRestricted - Represents Data for restricted discovery used to request the authorization to announce for a UE
type AnnounceDiscDataForRestricted struct {

	// Contains the RPAUID.
	Rpauid string `json:"rpauid"`

	// Contains the Application ID.
	AppId string `json:"appId"`

	// string with format 'date-time' as defined in OpenAPI.
	ValidityTime time.Time `json:"validityTime"`

	// Contains the ProSe Restricted Code.
	ProseRestrictedCode string `json:"proseRestrictedCode,omitempty"`

	// Contains the ProSe Restricted Code Prefix.
	ProseRestrictedPrefix string `json:"proseRestrictedPrefix,omitempty"`

	CodeSuffixPool RestrictedCodeSuffixPool `json:"codeSuffixPool,omitempty"`
}
