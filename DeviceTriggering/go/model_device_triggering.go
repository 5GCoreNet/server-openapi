/*
 * 3gpp-device-triggering
 *
 * API for device trigger.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package DeviceTriggering

// DeviceTriggering - Represents device triggering related information.
type DeviceTriggering struct {

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Self string `json:"self,omitempty"`

	// string containing a local identifier followed by \"@\" and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \"@\" characters. See Clause 4.6.2 of 3GPP TS 23.682 for more information.
	ExternalId string `json:"externalId,omitempty"`

	// string formatted according to clause 3.3 of 3GPP TS 23.003 that describes an MSISDN.
	Msisdn string `json:"msisdn,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	// Unsigned integer identifying a period of time in units of seconds.
	ValidityPeriod int32 `json:"validityPeriod"`

	Priority Priority `json:"priority"`

	// Unsigned integer with valid values between 0 and 65535.
	ApplicationPortId int32 `json:"applicationPortId"`

	// Unsigned integer with valid values between 0 and 65535.
	AppSrcPortId int32 `json:"appSrcPortId,omitempty"`

	// String with format \"byte\" as defined in OpenAPI Specification, i.e, base64-encoded characters.
	TriggerPayload string `json:"triggerPayload"`

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	NotificationDestination string `json:"notificationDestination"`

	// Set to true by the SCS/AS to request the SCEF to send a test notification as defined in clause 5.2.5.3. Set to false or omitted otherwise.
	RequestTestNotification bool `json:"requestTestNotification,omitempty"`

	WebsockNotifConfig WebsockNotifConfig `json:"websockNotifConfig,omitempty"`

	DeliveryResult DeliveryResult `json:"deliveryResult,omitempty"`
}