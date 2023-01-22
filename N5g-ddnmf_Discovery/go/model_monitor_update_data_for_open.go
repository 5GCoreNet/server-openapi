/*
 * N5g-ddnmf_Discovery API
 *
 * N5g-ddnmf_Discovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package N5g-ddnmf_Discovery

// MonitorUpdateDataForOpen - Represents Monitor Update Data for the Discovery Type \"OPEN\".
type MonitorUpdateDataForOpen struct {

	// Contains the ProSe Application ID name.
	ProseAppIdName string `json:"proseAppIdName"`

	Ttl int32 `json:"ttl"`
}