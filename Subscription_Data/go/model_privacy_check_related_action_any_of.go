/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Subscription_Data

type PrivacyCheckRelatedActionAnyOf string

// List of PrivacyCheckRelatedActionAnyOf
const (
	NOT_ALLOWED PrivacyCheckRelatedActionAnyOf = "LOCATION_NOT_ALLOWED"
	ALLOWED_WITH_NOTIFICATION PrivacyCheckRelatedActionAnyOf = "LOCATION_ALLOWED_WITH_NOTIFICATION"
	ALLOWED_WITHOUT_NOTIFICATION PrivacyCheckRelatedActionAnyOf = "LOCATION_ALLOWED_WITHOUT_NOTIFICATION"
	ALLOWED_WITHOUT_RESPONSE PrivacyCheckRelatedActionAnyOf = "LOCATION_ALLOWED_WITHOUT_RESPONSE"
	RESTRICTED_WITHOUT_RESPONSE PrivacyCheckRelatedActionAnyOf = "LOCATION_RESTRICTED_WITHOUT_RESPONSE"
)
