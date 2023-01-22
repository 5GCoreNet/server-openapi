/*
 * EES EEL Managed ACR Service
 *
 * EES EEL Managed ACR Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Eees_EELManagedACR

// UncertaintyEllipse - Ellipse with uncertainty.
type UncertaintyEllipse struct {

	// Indicates value of uncertainty.
	SemiMajor float32 `json:"semiMajor"`

	// Indicates value of uncertainty.
	SemiMinor float32 `json:"semiMinor"`

	// Indicates value of orientation angle.
	OrientationMajor int32 `json:"orientationMajor"`
}
