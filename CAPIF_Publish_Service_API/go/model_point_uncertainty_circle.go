/*
 * CAPIF_Publish_Service_API
 *
 * API for publishing service APIs.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package CAPIF_Publish_Service_API

// PointUncertaintyCircle - Ellipsoid point with uncertainty circle.
type PointUncertaintyCircle struct {

	Point GeographicalCoordinates `json:"point"`

	// Indicates value of uncertainty.
	Uncertainty float32 `json:"uncertainty"`
}
