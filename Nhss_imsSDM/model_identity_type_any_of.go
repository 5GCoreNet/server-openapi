/*
 * Nhss_imsSDM
 *
 * Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nhss_imsSDM

type IdentityTypeAnyOf string

// List of IdentityTypeAnyOf
const (
	DISTINCT_IMPU IdentityTypeAnyOf = "DISTINCT_IMPU"
	DISTINCT_PSI IdentityTypeAnyOf = "DISTINCT_PSI"
	WILDCARDED_IMPU IdentityTypeAnyOf = "WILDCARDED_IMPU"
	WILDCARDED_PSI IdentityTypeAnyOf = "WILDCARDED_PSI"
)
