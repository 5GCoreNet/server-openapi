/*
 * 3gpp-cp-parameter-provisioning
 *
 * API for provisioning communication pattern parameters.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CpProvisioning

// Local3dPointUncertaintyEllipsoid - Local 3D point with uncertainty ellipsoid
type Local3dPointUncertaintyEllipsoid struct {

	LocalOrigin LocalOrigin `json:"localOrigin"`

	Point RelativeCartesianLocation `json:"point"`

	UncertaintyEllipsoid UncertaintyEllipsoid `json:"uncertaintyEllipsoid"`

	// Indicates value of confidence.
	Confidence int32 `json:"confidence"`
}
