/*
 * SS_LocationAreaInfoRetrieval
 *
 * API for SEAL Location Area Info Retrieval.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package SS_LocationAreaInfoRetrieval

// Polygon - Polygon.
type Polygon struct {

	// List of points.
	PointList []GeographicalCoordinates `json:"pointList"`
}
