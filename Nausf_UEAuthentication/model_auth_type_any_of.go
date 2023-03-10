/*
 * AUSF API
 *
 * AUSF UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nausf_UEAuthentication

type AuthTypeAnyOf string

// List of AuthTypeAnyOf
const (
	_5_G_AKA AuthTypeAnyOf = "5G_AKA"
	EAP_AKA_PRIME AuthTypeAnyOf = "EAP_AKA_PRIME"
	EAP_TLS AuthTypeAnyOf = "EAP_TLS"
	EAP_TTLS AuthTypeAnyOf = "EAP_TTLS"
)
