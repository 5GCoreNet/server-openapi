/*
 * TS 29.122 Common Data Types
 *
 * Data types applicable to several APIs.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CommonData

// NotificationData - Represents the information to be conveyed in a bearer level event(s) notification.
type NotificationData struct {

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Transaction string `json:"transaction"`

	// Contains the reported event and applicable information
	EventReports []EventReport `json:"eventReports"`
}