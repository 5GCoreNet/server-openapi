/*
 * Ntsctsf_ASTI Service API
 *
 * TSCTSF  Access Stratum time distribution configuration Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ntsctsf_ASTI

// AccessTimeDistributionData - Contains the parameters for the creation of 5G access stratum time distribution configuration. 
type AccessTimeDistributionData struct {

	Supis []string `json:"supis,omitempty"`

	Gpsis []string `json:"gpsis,omitempty"`

	// String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.  
	InterGrpId string `json:"interGrpId,omitempty"`

	// String identifying External Group Identifier that identifies a group made up of one or more  subscriptions associated to a group of IMSIs, as specified in clause 19.7.3 of 3GPP TS 23.003.  
	ExterGrpId string `json:"exterGrpId,omitempty"`

	AsTimeDisParam AsTimeDistributionParam `json:"asTimeDisParam"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat string `json:"suppFeat,omitempty"`
}
