/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Subscription_Data

type DlDataDeliveryStatusAnyOf string

// List of DlDataDeliveryStatusAnyOf
const (
	BUFFERED DlDataDeliveryStatusAnyOf = "BUFFERED"
	TRANSMITTED DlDataDeliveryStatusAnyOf = "TRANSMITTED"
	DISCARDED DlDataDeliveryStatusAnyOf = "DISCARDED"
)
