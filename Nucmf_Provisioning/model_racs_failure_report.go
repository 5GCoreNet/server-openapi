/*
 * Nucmf_Provisioning
 *
 * UCMF_Provisioning Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nucmf_Provisioning

// RacsFailureReport - Represents a radio capability data provisioning failure report.
type RacsFailureReport struct {

	// Identifies the RACS ID(s) for which the RACS data are not provisioned successfully.
	RacsIds []string `json:"racsIds"`

	FailureCode RacsFailureCode `json:"failureCode"`
}
