/*
 * Nchf_OfflineOnlyCharging
 *
 * OfflineOnlyCharging Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nchf_OfflineOnlyCharging

type PduSessionTypeAnyOf string

// List of PduSessionTypeAnyOf
const (
	IPV4 PduSessionTypeAnyOf = "IPV4"
	IPV6 PduSessionTypeAnyOf = "IPV6"
	IPV4_V6 PduSessionTypeAnyOf = "IPV4V6"
	UNSTRUCTURED PduSessionTypeAnyOf = "UNSTRUCTURED"
	ETHERNET PduSessionTypeAnyOf = "ETHERNET"
)
