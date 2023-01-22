/*
 * Common Data Types
 *
 * Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   
 *
 * API version: 1.4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package CommonData

// RoutingAreaId - Contains a Routing Area Identification as defined in 3GPP TS 23.003, clause 4.2.
type RoutingAreaId struct {

	PlmnId PlmnId `json:"plmnId"`

	// Location Area Code
	Lac string `json:"lac"`

	// Routing Area Code
	Rac string `json:"rac"`
}
