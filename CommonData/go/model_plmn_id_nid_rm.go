/*
 * Common Data Types
 *
 * Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   
 *
 * API version: 1.4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package CommonData

// PlmnIdNidRm - This data type is defined in the same way as the 'PlmnIdNid' data type, but with the OpenAPI 'nullable: true' property. 
type PlmnIdNidRm struct {

	// Mobile Country Code part of the PLMN, comprising 3 digits, as defined in clause 9.3.3.5 of 3GPP TS 38.413.  
	Mcc string `json:"mcc"`

	// Mobile Network Code part of the PLMN, comprising 2 or 3 digits, as defined in clause 9.3.3.5 of 3GPP TS 38.413.
	Mnc string `json:"mnc"`

	// This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).  
	Nid string `json:"nid,omitempty"`
}
