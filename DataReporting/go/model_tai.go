/*
 * 3gpp-data-reporting
 *
 * API for 3GPP Data Reporting.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package DataReporting

// Tai - Contains the tracking area identity as described in 3GPP 23.003
type Tai struct {

	PlmnId PlmnId `json:"plmnId"`

	// 2 or 3-octet string identifying a tracking area code as specified in clause 9.3.3.10  of 3GPP TS 38.413, in hexadecimal representation. Each character in the string shall  take a value of \"0\" to \"9\", \"a\" to \"f\" or \"A\" to \"F\" and shall represent 4 bits. The most significant character representing the 4 most significant bits of the TAC shall  appear first in the string, and the character representing the 4 least significant bit  of the TAC shall appear last in the string.  
	Tac string `json:"tac"`

	// This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).  
	Nid string `json:"nid,omitempty"`
}
