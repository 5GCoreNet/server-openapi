/*
 * SS_GroupManagement
 *
 * API for SEAL Group management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package SS_GroupManagement

// ValGroupDocument - Represents details of the VAL group document information.
type ValGroupDocument struct {

	// The VAL group idenitity.
	ValGroupId string `json:"valGroupId"`

	// The text description of the VAL group.
	GrpDesc string `json:"grpDesc,omitempty"`

	// The list of VAL User IDs or VAL UE IDs, which are members of the VAL group.
	Members []ValTargetUe `json:"members,omitempty"`

	// Configuration data for the VAL group.
	ValGrpConf string `json:"valGrpConf,omitempty"`

	// The list of VAL services enabled on the group.
	ValServiceIds []string `json:"valServiceIds,omitempty"`

	// VAL service specific information.
	ValSvcInf string `json:"valSvcInf,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat string `json:"suppFeat,omitempty"`

	// string providing an URI formatted according to IETF RFC 3986.
	ResUri string `json:"resUri,omitempty"`

	LocInfo LocationInfo `json:"locInfo,omitempty"`

	AddLocInfo LocationArea5G `json:"addLocInfo,omitempty"`

	// string containing a local identifier followed by \"@\" and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \"@\" characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information.
	ExtGrpId string `json:"extGrpId,omitempty"`

	Com5GLanType PduSessionType `json:"com5GLanType,omitempty"`
}