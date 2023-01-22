/*
 * Common Data Types
 *
 * Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   
 *
 * API version: 1.4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package CommonData

// SpatialValidityCond - Contains the Spatial Validity Condition.
type SpatialValidityCond struct {

	TrackingAreaList []Tai `json:"trackingAreaList,omitempty"`

	Countries []string `json:"countries,omitempty"`

	GeographicalServiceArea GeoServiceArea `json:"geographicalServiceArea,omitempty"`
}
