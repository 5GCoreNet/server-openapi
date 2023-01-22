/*
 * NhssUEAU
 *
 * HSS UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nhss_UEAU

type AuthTypeAnyOf string

// List of AuthTypeAnyOf
const (
	_5_G_AKA AuthTypeAnyOf = "5G_AKA"
	EAP_AKA_PRIME AuthTypeAnyOf = "EAP_AKA_PRIME"
	EAP_TLS AuthTypeAnyOf = "EAP_TLS"
	NONE AuthTypeAnyOf = "NONE"
	EAP_TTLS AuthTypeAnyOf = "EAP_TTLS"
)