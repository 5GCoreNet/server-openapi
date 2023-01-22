/*
 * TS 29.122 Common Data Types
 *
 * Data types applicable to several APIs.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package CommonData

// TestNotification - Represents a notification that can be sent to test whether a chosen notification mechanism works.
type TestNotification struct {

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Subscription string `json:"subscription"`
}