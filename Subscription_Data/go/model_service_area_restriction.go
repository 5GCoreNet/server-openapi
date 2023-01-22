/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Subscription_Data

// ServiceAreaRestriction - Provides information about allowed or not allowed areas.
type ServiceAreaRestriction struct {

	RestrictionType RestrictionType `json:"restrictionType,omitempty"`

	Areas []Area `json:"areas,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxNumOfTAs int32 `json:"maxNumOfTAs,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxNumOfTAsForNotAllowedAreas int32 `json:"maxNumOfTAsForNotAllowedAreas,omitempty"`
}
