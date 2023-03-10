/*
 * Eecs_ServiceProvisioning
 *
 * API for ECS Service Provisioning. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Eecs_ServiceProvisioning

// HorizontalVelocity - Horizontal velocity.
type HorizontalVelocity struct {

	// Indicates value of horizontal speed.
	HSpeed float32 `json:"hSpeed"`

	// Indicates value of angle.
	Bearing int32 `json:"bearing"`
}
