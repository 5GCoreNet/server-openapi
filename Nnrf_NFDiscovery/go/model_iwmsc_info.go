/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnrf_NFDiscovery

// IwmscInfo - Information of an SMS-IWMSC NF Instance
type IwmscInfo struct {

	MsisdnRanges []IdentityRange `json:"msisdnRanges,omitempty"`

	SupiRanges []SupiRange `json:"supiRanges,omitempty"`

	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`

	ScNumber string `json:"scNumber,omitempty"`
}
