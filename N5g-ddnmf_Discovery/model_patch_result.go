/*
 * N5g-ddnmf_Discovery API
 *
 * N5g-ddnmf_Discovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_N5g-ddnmf_Discovery

// PatchResult - The execution report result on failed modification.
type PatchResult struct {

	// The execution report contains an array of report items. Each report item indicates one  failed modification. 
	Report []ReportItem `json:"report"`
}
