/*
 * UAE Server Real-time UAV Status Service
 *
 * UAE Server Real-time UAV Status Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_UAE_RealtimeUAVStatus

// LocalOrigin - Indicates a Local origin in a reference system
type LocalOrigin struct {

	CoordinateId string `json:"coordinateId,omitempty"`

	Point GeographicalCoordinates `json:"point,omitempty"`
}
