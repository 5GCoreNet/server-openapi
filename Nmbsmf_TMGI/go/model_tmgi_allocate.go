/*
 * Nmbsmf_TMGI
 *
 * MB-SMF TMGI Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nmbsmf_TMGI

// TmgiAllocate - Data within TMGI Allocate Request
type TmgiAllocate struct {

	// The number of requested TMGIs
	TmgiNumber int32 `json:"tmgiNumber,omitempty"`

	// The list of TMGIs to be refreshed
	TmgiList []Tmgi `json:"tmgiList,omitempty"`
}
