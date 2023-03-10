/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Subscription_Data

type LcsClientClassAnyOf string

// List of LcsClientClassAnyOf
const (
	BROADCAST_SERVICE LcsClientClassAnyOf = "BROADCAST_SERVICE"
	OM_IN_HPLMN LcsClientClassAnyOf = "OM_IN_HPLMN"
	OM_IN_VPLMN LcsClientClassAnyOf = "OM_IN_VPLMN"
	ANONYMOUS_LOCATION_SERVICE LcsClientClassAnyOf = "ANONYMOUS_LOCATION_SERVICE"
	SPECIFIC_SERVICE LcsClientClassAnyOf = "SPECIFIC_SERVICE"
)
