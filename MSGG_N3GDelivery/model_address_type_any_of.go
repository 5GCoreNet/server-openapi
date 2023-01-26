/*
 * MSGG_N3GDelivery
 *
 * API for MSGG N3G Message Delivery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_MSGG_N3GDelivery

type AddressTypeAnyOf string

// List of AddressTypeAnyOf
const (
	UE AddressTypeAnyOf = "UE"
	AS AddressTypeAnyOf = "AS"
	GROUP AddressTypeAnyOf = "GROUP"
	BC AddressTypeAnyOf = "BC"
	TOPIC AddressTypeAnyOf = "TOPIC"
)