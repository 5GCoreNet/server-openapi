/*
 * 3gpp-network-status-reporting
 *
 * API for reporting network status.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_ReportingNetworkStatus

// PointAltitude - Ellipsoid point with altitude.
type PointAltitude struct {

	Point GeographicalCoordinates `json:"point"`

	// Indicates value of altitude.
	Altitude float64 `json:"altitude"`
}
