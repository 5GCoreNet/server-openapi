/*
 * GMDviaMBMSbyMB2
 *
 * API for Group Message Delivery via MBMS by MB2   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_GMDviaMBMSbyMB2

import (
	"time"
)

// GmdViaMbmsbyMb2Patch - Represents a modification request of a group message delivery via MBMS by MB2.
type GmdViaMbmsbyMb2Patch struct {

	// string containing a local identifier followed by \"@\" and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \"@\" characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information.
	ExternalGroupId string `json:"externalGroupId,omitempty"`

	MbmsLocArea MbmsLocArea `json:"mbmsLocArea,omitempty"`

	// string with format \"date-time\" as defined in OpenAPI.
	MessageDeliveryStartTime time.Time `json:"messageDeliveryStartTime,omitempty"`

	// String with format \"byte\" as defined in OpenAPI Specification, i.e, base64-encoded characters.
	GroupMessagePayload string `json:"groupMessagePayload,omitempty"`

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	NotificationDestination string `json:"notificationDestination,omitempty"`
}