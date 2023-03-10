/*
 * 3gpp-ecs-address-provision
 *
 * API for ECS Address Provisioning.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_EcsAddressProvision

// SpatialValidityCond - Contains the Spatial Validity Condition.
type SpatialValidityCond struct {

	TrackingAreaList []Tai `json:"trackingAreaList,omitempty"`

	Countries []string `json:"countries,omitempty"`

	GeographicalServiceArea GeoServiceArea `json:"geographicalServiceArea,omitempty"`
}
