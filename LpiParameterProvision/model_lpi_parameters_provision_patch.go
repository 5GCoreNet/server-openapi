/*
 * 3gpp-lpi-pp
 *
 * API for Location Privacy Indication Parameters Provisioning.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_LpiParameterProvision

// LpiParametersProvisionPatch - Represents the parameters to modify an existing Individual LPI Parameters Provisionings resource.
type LpiParametersProvisionPatch struct {

	Lpi Lpi `json:"lpi,omitempty"`

	// String uniquely identifying MTC provider information.
	MtcProviderId string `json:"mtcProviderId,omitempty"`
}
