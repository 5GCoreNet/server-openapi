/*
 * Eees_EASDiscovery
 *
 * API for EAS Discovery. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Eees_EASDiscovery

// HorizontalVelocityWithUncertainty - Horizontal velocity with speed uncertainty.
type HorizontalVelocityWithUncertainty struct {

	// Indicates value of horizontal speed.
	HSpeed float32 `json:"hSpeed"`

	// Indicates value of angle.
	Bearing int32 `json:"bearing"`

	// Indicates value of speed uncertainty.
	HUncertainty float32 `json:"hUncertainty"`
}
