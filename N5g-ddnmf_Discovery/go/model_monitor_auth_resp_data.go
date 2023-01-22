/*
 * N5g-ddnmf_Discovery API
 *
 * N5g-ddnmf_Discovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package N5g-ddnmf_Discovery

// MonitorAuthRespData - Represents the obtained Monitor Authorize Data for a UE
type MonitorAuthRespData struct {

	AuthDataOpen MonitorAuthDataForOpen `json:"authDataOpen,omitempty"`

	AuthDataRestricted MonitorAuthDataForRestricted `json:"authDataRestricted,omitempty"`
}
