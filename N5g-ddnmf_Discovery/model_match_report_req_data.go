/*
 * N5g-ddnmf_Discovery API
 *
 * N5g-ddnmf_Discovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_N5g-ddnmf_Discovery

// MatchReportReqData - Represents the Match Report information
type MatchReportReqData struct {

	DiscType DiscoveryType `json:"discType"`

	ProseAppCodes []string `json:"proseAppCodes,omitempty"`

	MoniteredPlmnId PlmnId `json:"moniteredPlmnId,omitempty"`
}
