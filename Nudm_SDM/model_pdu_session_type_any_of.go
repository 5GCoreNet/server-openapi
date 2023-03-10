/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudm_SDM

type PduSessionTypeAnyOf string

// List of PduSessionTypeAnyOf
const (
	IPV4 PduSessionTypeAnyOf = "IPV4"
	IPV6 PduSessionTypeAnyOf = "IPV6"
	IPV4_V6 PduSessionTypeAnyOf = "IPV4V6"
	UNSTRUCTURED PduSessionTypeAnyOf = "UNSTRUCTURED"
	ETHERNET PduSessionTypeAnyOf = "ETHERNET"
)
