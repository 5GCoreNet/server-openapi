/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nsmf_PDUSession

type QosResourceTypeAnyOf string

// List of QosResourceTypeAnyOf
const (
	NON_GBR QosResourceTypeAnyOf = "NON_GBR"
	NON_CRITICAL_GBR QosResourceTypeAnyOf = "NON_CRITICAL_GBR"
	CRITICAL_GBR QosResourceTypeAnyOf = "CRITICAL_GBR"
)
