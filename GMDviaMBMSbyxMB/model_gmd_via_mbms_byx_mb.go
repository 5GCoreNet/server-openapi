/*
 * GMDviaMBMSbyxMB
 *
 * API for Group Message Delivery via MBMS by xMB   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_GMDviaMBMSbyxMB

import (
	"time"
)

// GmdViaMbmsByxMb - Represents a group message delivery via MBMS by xMB.
type GmdViaMbmsByxMb struct {

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Self string `json:"self,omitempty"`

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	NotificationDestination string `json:"notificationDestination"`

	// Set to true by the SCS/AS to request the SCEF to send a test notification as defined in clause 5.2.5.3. Set to false or omitted otherwise.
	RequestTestNotification bool `json:"requestTestNotification,omitempty"`

	WebsockNotifConfig WebsockNotifConfig `json:"websockNotifConfig,omitempty"`

	MbmsLocArea MbmsLocArea `json:"mbmsLocArea,omitempty"`

	// string with format \"date-time\" as defined in OpenAPI.
	MessageDeliveryStartTime time.Time `json:"messageDeliveryStartTime,omitempty"`

	// string with format \"date-time\" as defined in OpenAPI.
	MessageDeliveryStopTime time.Time `json:"messageDeliveryStopTime,omitempty"`

	// String with format \"byte\" as defined in OpenAPI Specification, i.e, base64-encoded characters.
	GroupMessagePayload string `json:"groupMessagePayload,omitempty"`

	// string identifying a Ipv4 address formatted in the \"dotted decimal\" notation as defined in IETF RFC 1166, with \"readOnly=true\" property.
	ScefMessageDeliveryIPv4 string `json:"scefMessageDeliveryIPv4,omitempty"`

	// string identifying a Ipv6 address formatted according to clause 4 in IETF RFC 5952, with \"readOnly=true\" property. The mixed Ipv4 Ipv6 notation according to clause 5 of IETF RFC 5952 shall not be used.
	ScefMessageDeliveryIPv6 string `json:"scefMessageDeliveryIPv6,omitempty"`

	// Unsigned integer with valid values between 0 and 65535, with \"readOnly=true\" property.
	ScefMessageDeliveryPort int32 `json:"scefMessageDeliveryPort,omitempty"`
}
