/*
 * 3gpp-device-triggering
 *
 * API for device trigger.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_DeviceTriggering

type DeliveryResultAnyOf string

// List of DeliveryResultAnyOf
const (
	SUCCESS DeliveryResultAnyOf = "SUCCESS"
	UNKNOWN DeliveryResultAnyOf = "UNKNOWN"
	FAILURE DeliveryResultAnyOf = "FAILURE"
	TRIGGERED DeliveryResultAnyOf = "TRIGGERED"
	EXPIRED DeliveryResultAnyOf = "EXPIRED"
	UNCONFIRMED DeliveryResultAnyOf = "UNCONFIRMED"
	REPLACED DeliveryResultAnyOf = "REPLACED"
	TERMINATE DeliveryResultAnyOf = "TERMINATE"
)
