/*
 * GMDviaMBMSbyxMB
 *
 * API for Group Message Delivery via MBMS by xMB   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_GMDviaMBMSbyxMB

// ServiceCreation - Represents an individual xMB Service resource.
type ServiceCreation struct {

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Self string `json:"self,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	// string containing a local identifier followed by \"@\" and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \"@\" characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information.
	ExternalGroupId string `json:"externalGroupId,omitempty"`

	// Identifies the MBMS User Service supplied by the SCEF.
	UserServiceId string `json:"userServiceId,omitempty"`

	// The service class that service belongs to supplied by the SCEF.
	ServiceClass string `json:"serviceClass,omitempty"`

	// List of language of the service content supplied by the SCEF.
	ServiceLanguages []string `json:"serviceLanguages,omitempty"`

	// List of Service Names supplied by the SCEF.
	ServiceNames []string `json:"serviceNames,omitempty"`

	// When set to 'true', the Content Provider indicates that the service is a Receive Only Mode service. This parameter is supplied by the SCEF.
	ReceiveOnlyMode bool `json:"receiveOnlyMode,omitempty"`

	ServiceAnnouncementMode ServiceAnnouncementMode `json:"serviceAnnouncementMode,omitempty"`
}
