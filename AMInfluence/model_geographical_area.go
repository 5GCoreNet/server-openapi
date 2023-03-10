/*
 * AMInfluence
 *
 * AMInfluence API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_AMInfluence

// GeographicalArea - Contains geographical area information (e.g.a civic address or shapes).
type GeographicalArea struct {

	CivicAddress CivicAddress `json:"civicAddress,omitempty"`

	Shapes GeographicArea `json:"shapes,omitempty"`
}
