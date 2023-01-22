/*
 * VAE_HDMapDynamicInfo
 *
 * API for VAE HDMapDynamicInfo Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package VAE_HDMapDynamicInfo

// HdMapDynamicInfoData - Represents an individual HdMap DynamicInfo Subscription resource for a V2X UE ID.
type HdMapDynamicInfoData struct {

	// Represents the identifier of the V2X UE.
	UeId string `json:"ueId"`

	// String providing an URI formatted according to RFC 3986.
	NotifUri string `json:"notifUri"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Range int32 `json:"range"`

	// Set to true by the NF service consumer to request the VAE server to send a test notification as defined in clause 6.3.5.3. Set to false or omitted otherwise. 
	RequestTestNotification bool `json:"requestTestNotification,omitempty"`

	WebsockNotifConfig WebsockNotifConfig `json:"websockNotifConfig,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat string `json:"suppFeat,omitempty"`
}
