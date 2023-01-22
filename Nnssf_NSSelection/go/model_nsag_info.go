/*
 * NSSF NS Selection
 *
 * NSSF Network Slice Selection Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnssf_NSSelection

// NsagInfo - Contains the association of NSAGs and S-NSSAI(s) along with the TA(s) within which the association is valid.
type NsagInfo struct {

	NsagIds []int32 `json:"nsagIds"`

	SnssaiList []Snssai `json:"snssaiList"`

	TaiList []Tai `json:"taiList,omitempty"`

	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`
}