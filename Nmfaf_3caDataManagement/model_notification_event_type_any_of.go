/*
 * Nmfaf_3caDataManagement
 *
 * MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmfaf_3caDataManagement

type NotificationEventTypeAnyOf string

// List of NotificationEventTypeAnyOf
const (
	REGISTERED NotificationEventTypeAnyOf = "NF_REGISTERED"
	DEREGISTERED NotificationEventTypeAnyOf = "NF_DEREGISTERED"
	PROFILE_CHANGED NotificationEventTypeAnyOf = "NF_PROFILE_CHANGED"
)
