/*
 * SS_NetworkResourceAdaptation
 *
 * SS Network Resource Adaptation Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_SS_NetworkResourceAdaptation

// StreamSpecification - Stream specification includes MAC addresses of the source and destination DS-TT ports. 
type StreamSpecification struct {

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042. 
	SrcMacAddr string `json:"srcMacAddr"`

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042. 
	DstMacAddr string `json:"dstMacAddr"`
}
