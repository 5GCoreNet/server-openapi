/*
 * 3gpp-nidd
 *
 * API for non IP data delivery.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_NIDD

// RdsPort - Represents the port configuration for Reliable Data Transfer.
type RdsPort struct {

	// Unsigned integer with valid values between 0 and 65535.
	PortUE int32 `json:"portUE"`

	// Unsigned integer with valid values between 0 and 65535.
	PortSCEF int32 `json:"portSCEF"`
}
