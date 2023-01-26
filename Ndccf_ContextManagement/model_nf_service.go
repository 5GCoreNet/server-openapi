/*
 * Ndccf_ContextManagement
 *
 * DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ndccf_ContextManagement

import (
	"time"
)

// NfService - Information of a given NF Service Instance; it is part of the NFProfile of an NF Instance 
type NfService struct {

	ServiceInstanceId string `json:"serviceInstanceId"`

	ServiceName ServiceName `json:"serviceName"`

	Versions []NfServiceVersion `json:"versions"`

	Scheme UriScheme `json:"scheme"`

	NfServiceStatus NfServiceStatus `json:"nfServiceStatus"`

	// Fully Qualified Domain Name
	Fqdn string `json:"fqdn,omitempty"`

	// Fully Qualified Domain Name
	InterPlmnFqdn string `json:"interPlmnFqdn,omitempty"`

	IpEndPoints []IpEndPoint `json:"ipEndPoints,omitempty"`

	ApiPrefix string `json:"apiPrefix,omitempty"`

	DefaultNotificationSubscriptions []DefaultNotificationSubscription `json:"defaultNotificationSubscriptions,omitempty"`

	AllowedPlmns []PlmnId `json:"allowedPlmns,omitempty"`

	AllowedSnpns []PlmnIdNid `json:"allowedSnpns,omitempty"`

	AllowedNfTypes []NfType `json:"allowedNfTypes,omitempty"`

	AllowedNfDomains []string `json:"allowedNfDomains,omitempty"`

	AllowedNssais []ExtSnssai `json:"allowedNssais,omitempty"`

	// A map (list of key-value pairs) where NF Type serves as key
	AllowedOperationsPerNfType map[string][]string `json:"allowedOperationsPerNfType,omitempty"`

	// A map (list of key-value pairs) where NF Instance Id serves as key
	AllowedOperationsPerNfInstance map[string][]string `json:"allowedOperationsPerNfInstance,omitempty"`

	AllowedOperationsPerNfInstanceOverrides bool `json:"allowedOperationsPerNfInstanceOverrides,omitempty"`

	Priority int32 `json:"priority,omitempty"`

	Capacity int32 `json:"capacity,omitempty"`

	Load int32 `json:"load,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	LoadTimeStamp time.Time `json:"loadTimeStamp,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	RecoveryTime time.Time `json:"recoveryTime,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	NfServiceSetIdList []string `json:"nfServiceSetIdList,omitempty"`

	SNssais []ExtSnssai `json:"sNssais,omitempty"`

	PerPlmnSnssaiList []PlmnSnssai `json:"perPlmnSnssaiList,omitempty"`

	// Vendor ID of the NF Service instance (Private Enterprise Number assigned by IANA)
	VendorId string `json:"vendorId,omitempty"`

	// A map (list of key-value pairs) where IANA-assigned SMI Network Management Private Enterprise Codes serves as key 
	SupportedVendorSpecificFeatures map[string][]VendorSpecificFeature `json:"supportedVendorSpecificFeatures,omitempty"`

	Oauth2Required bool `json:"oauth2Required,omitempty"`

	PerPlmnOauth2ReqList PlmnOauth2 `json:"perPlmnOauth2ReqList,omitempty"`
}