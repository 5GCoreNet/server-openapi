/*
 * AUSF API
 *
 * AUSF UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nausf_UEAuthentication

// ProSeEapSession - Contains information related to the EAP session.
type ProSeEapSession struct {

	// contains an EAP packet
	EapPayload *string `json:"eapPayload"`

	// Contains the KNR_ProSe.
	KnrProSe string `json:"knrProSe,omitempty"`

	// A map(list of key-value pairs) where member serves as key
	Links map[string]LinksValueSchema `json:"_links,omitempty"`

	AuthResult AuthResult `json:"authResult,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	// contains an Nonce2
	Nonce2 *string `json:"nonce2,omitempty"`

	// The 5GPRUK ID is string in NAI format as specified in clause 28.7.19 of 3GPP TS 23.003. 
	Var5gPrukId string `json:"5gPrukId,omitempty"`
}
