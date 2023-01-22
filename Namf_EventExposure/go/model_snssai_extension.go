/*
 * Namf_EventExposure
 *
 * AMF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Namf_EventExposure

// SnssaiExtension - Extensions to the Snssai data type, sdRanges and wildcardSd shall not be present simultaneously 
type SnssaiExtension struct {

	// When present, it shall contain the range(s) of Slice Differentiator values supported for the Slice/Service Type value indicated in the sst attribute of the Snssai data type 
	SdRanges []SdRange `json:"sdRanges,omitempty"`

	// When present, it shall be set to true, to indicate that all SD values are supported for the Slice/Service Type value indicated in the sst attribute of the Snssai data type. 
	WildcardSd bool `json:"wildcardSd,omitempty"`
}
