/*
 * 3gpp-network-status-reporting
 *
 * API for reporting network status.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_ReportingNetworkStatus

type PolygonAllOf struct {

	// List of points.
	PointList []GeographicalCoordinates `json:"pointList"`
}
