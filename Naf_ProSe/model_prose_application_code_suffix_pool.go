/*
 * Naf_ProSe API
 *
 * Naf_ProSe Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Naf_ProSe

// ProseApplicationCodeSuffixPool - Contains the Prose Application Code Suffix Pool.
type ProseApplicationCodeSuffixPool struct {

	// Contains the ProSe Application Code Suffix.
	CodeSuffix string `json:"codeSuffix,omitempty"`

	CodeSuffixRange ProseAppCodeSuffixRange `json:"codeSuffixRange,omitempty"`
}
