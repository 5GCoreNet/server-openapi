/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_AnalyticsInfo

// NfLoadLevelInformation - Represents load level information of a given NF instance.
type NfLoadLevelInformation struct {

	NfType NfType `json:"nfType,omitempty"`

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NfInstanceId string `json:"nfInstanceId,omitempty"`

	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.  
	NfSetId string `json:"nfSetId,omitempty"`

	NfStatus NfStatus `json:"nfStatus,omitempty"`

	NfCpuUsage int32 `json:"nfCpuUsage,omitempty"`

	NfMemoryUsage int32 `json:"nfMemoryUsage,omitempty"`

	NfStorageUsage int32 `json:"nfStorageUsage,omitempty"`

	NfLoadLevelAverage int32 `json:"nfLoadLevelAverage,omitempty"`

	NfLoadLevelpeak int32 `json:"nfLoadLevelpeak,omitempty"`

	NfLoadAvgInAoi int32 `json:"nfLoadAvgInAoi,omitempty"`

	Snssai Snssai `json:"snssai,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence int32 `json:"confidence,omitempty"`
}
