/*
 * SS_LocationAreaInfoRetrieval
 *
 * API for SEAL Location Area Info Retrieval.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package SS_LocationAreaInfoRetrieval

// PointAltitudeUncertainty - Ellipsoid point with altitude and uncertainty ellipsoid.
type PointAltitudeUncertainty struct {

	Point GeographicalCoordinates `json:"point"`

	// Indicates value of altitude.
	Altitude float64 `json:"altitude"`

	UncertaintyEllipse UncertaintyEllipse `json:"uncertaintyEllipse"`

	// Indicates value of uncertainty.
	UncertaintyAltitude float32 `json:"uncertaintyAltitude"`

	// Indicates value of confidence.
	Confidence int32 `json:"confidence"`
}
