/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_AnalyticsInfo

// NsmfEventExposureNotification - Represents notifications on events that occurred.
type NsmfEventExposureNotification struct {

	// Notification correlation ID
	NotifId string `json:"notifId"`

	// Notifications about Individual Events
	EventNotifs []EventNotification1 `json:"eventNotifs"`

	// String providing an URI formatted according to RFC 3986.
	AckUri string `json:"ackUri,omitempty"`
}
