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

// AuthDataForRestricted - Represents obtained authorization Data for restricted discovery for a discoverer UE
type AuthDataForRestricted struct {

	ProseQueryCodes []string `json:"proseQueryCodes"`

	// Contains the ProSe Response Code.
	ProseRespCode string `json:"proseRespCode"`

	// string with format 'date-time' as defined in OpenAPI.
	ValidityTime time.Time `json:"validityTime"`
}