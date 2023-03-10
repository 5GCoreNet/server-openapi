/*
 * 3GPP 5GC NRM
 *
 * OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_5GcNrm

type NotificationType string

// List of NotificationType
const (
	N1_MESSAGES NotificationType = "N1_MESSAGES"
	N2_INFORMATION NotificationType = "N2_INFORMATION"
	LOCATION_NOTIFICATION NotificationType = "LOCATION_NOTIFICATION"
	DATA_REMOVAL_NOTIFICATION NotificationType = "DATA_REMOVAL_NOTIFICATION"
	DATA_CHANGE_NOTIFICATION NotificationType = "DATA_CHANGE_NOTIFICATION"
	LOCATION_UPDATE_NOTIFICATION NotificationType = "LOCATION_UPDATE_NOTIFICATION"
	NSSAA_REAUTH_NOTIFICATION NotificationType = "NSSAA_REAUTH_NOTIFICATION"
	NSSAA_REVOC_NOTIFICATION NotificationType = "NSSAA_REVOC_NOTIFICATION"
)
