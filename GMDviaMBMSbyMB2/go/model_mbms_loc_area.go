/*
 * GMDviaMBMSbyMB2
 *
 * API for Group Message Delivery via MBMS by MB2   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package GMDviaMBMSbyMB2

// MbmsLocArea - Represents a user location area whithin which is sent a group message delivery via MBMS request.
type MbmsLocArea struct {

	// Indicates a Cell Global Identification of the user which identifies the cell the UE is registered.
	CellId []string `json:"cellId,omitempty"`

	// Indicates an eNodeB in which the UE is currently located.
	EnodeBId []string `json:"enodeBId,omitempty"`

	// Identifies a geographic area of the user where the UE is located.
	GeographicArea []GeographicArea `json:"geographicArea,omitempty"`

	// Identifies an MBMS Service Area Identity of the user where the UE is located.
	MbmsServiceAreaId []string `json:"mbmsServiceAreaId,omitempty"`

	// Identifies a civic address of the user where the UE is located.
	CivicAddress []CivicAddress `json:"civicAddress,omitempty"`
}
