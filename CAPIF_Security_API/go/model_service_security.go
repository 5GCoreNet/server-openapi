/*
 * CAPIF_Security_API
 *
 * API for CAPIF security management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package CAPIF_Security_API

// ServiceSecurity - Represents the details of the security method for each service API interface. When included by the API invoker, it indicates the preferred method of security. When included by the CAPIF core function, it indicates the security method to be used for the service API interface. 
type ServiceSecurity struct {

	SecurityInfo []SecurityInformation `json:"securityInfo"`

	// string providing an URI formatted according to IETF RFC 3986.
	NotificationDestination string `json:"notificationDestination"`

	// Set to true by API invoker to request the CAPIF core function to send a test notification as defined in in clause 7.6. Set to false or omitted otherwise. 
	RequestTestNotification bool `json:"requestTestNotification,omitempty"`

	WebsockNotifConfig WebsockNotifConfig `json:"websockNotifConfig,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
}
