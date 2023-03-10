/*
 * 3gpp-racs-parameter-provisioning
 *
 * API for provisioning UE radio capability parameters.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_RacsParameterProvisioning

// RacsProvisioningDataPatch - Represents parameters to request the modification of a UE's radio capability data.
type RacsProvisioningDataPatch struct {

	// Identifies the configuration related to manufactuer specific UE radio capability. Each element uniquely identifies an RACS configuration for an RACS ID and is identified in the map via the RACS ID as key.
	RacsConfigs map[string]RacsConfigurationRm `json:"racsConfigs,omitempty"`
}
