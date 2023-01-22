/*
 * Common Data Types
 *
 * Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   
 *
 * API version: 1.4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package CommonData

// WirelineServiceAreaRestriction - The \"restrictionType\" attribute and the \"areas\" attribute shall be either both present or absent.  The empty array of areas is used when service is allowed/restricted nowhere. 
type WirelineServiceAreaRestriction struct {

	RestrictionType RestrictionType `json:"restrictionType,omitempty"`

	Areas []WirelineArea `json:"areas,omitempty"`
}
