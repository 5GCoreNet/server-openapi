/*
 * 3gpp-as-session-with-qos
 *
 * API for setting us an AS session with required QoS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_AsSessionWithQoS

// UserPlaneNotificationData - Represents the parameters to be conveyed in a user plane event(s) notification.
type UserPlaneNotificationData struct {

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Transaction string `json:"transaction"`

	// Contains the reported event and applicable information
	EventReports []UserPlaneEventReport `json:"eventReports"`
}
