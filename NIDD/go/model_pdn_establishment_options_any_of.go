/*
 * 3gpp-nidd
 *
 * API for non IP data delivery.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package NIDD

type PdnEstablishmentOptionsAnyOf string

// List of PdnEstablishmentOptionsAnyOf
const (
	WAIT_FOR_UE PdnEstablishmentOptionsAnyOf = "WAIT_FOR_UE"
	INDICATE_ERROR PdnEstablishmentOptionsAnyOf = "INDICATE_ERROR"
	SEND_TRIGGER PdnEstablishmentOptionsAnyOf = "SEND_TRIGGER"
)
