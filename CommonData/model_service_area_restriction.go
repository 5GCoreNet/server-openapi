/*
 * Common Data Types
 *
 * Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   
 *
 * API version: 1.4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CommonData

// ServiceAreaRestriction - Provides information about allowed or not allowed areas.
type ServiceAreaRestriction struct {

	RestrictionType RestrictionType `json:"restrictionType,omitempty"`

	Areas []Area `json:"areas,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxNumOfTAs int32 `json:"maxNumOfTAs,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxNumOfTAsForNotAllowedAreas int32 `json:"maxNumOfTAsForNotAllowedAreas,omitempty"`
}
