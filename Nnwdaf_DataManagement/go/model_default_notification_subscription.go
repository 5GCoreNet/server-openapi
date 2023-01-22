/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_DataManagement

// DefaultNotificationSubscription - Data structure for specifying the notifications the NF service subscribes by default, along with callback URI 
type DefaultNotificationSubscription struct {

	NotificationType NotificationType `json:"notificationType"`

	// String providing an URI formatted according to RFC 3986.
	CallbackUri string `json:"callbackUri"`

	N1MessageClass N1MessageClass `json:"n1MessageClass,omitempty"`

	N2InformationClass N2InformationClass `json:"n2InformationClass,omitempty"`

	Versions []string `json:"versions,omitempty"`

	Binding string `json:"binding,omitempty"`

	AcceptedEncoding string `json:"acceptedEncoding,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	// A map of service specific information. The name of the corresponding service (as specified in ServiceName data type) is the key. 
	ServiceInfoList map[string]DefSubServiceInfo `json:"serviceInfoList,omitempty"`
}
