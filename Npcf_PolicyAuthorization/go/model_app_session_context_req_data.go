/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Npcf_PolicyAuthorization

// AppSessionContextReqData - Identifies the service requirements of an Individual Application Session Context.
type AppSessionContextReqData struct {

	// Contains an AF application identifier.
	AfAppId string `json:"afAppId,omitempty"`

	// Application provided charging identifier allowing correlation of charging information.
	AfChargId string `json:"afChargId,omitempty"`

	AfReqData AfRequestedData `json:"afReqData,omitempty"`

	AfRoutReq AfRoutingRequirement `json:"afRoutReq,omitempty"`

	// Contains an identity of an application service provider.
	AspId string `json:"aspId,omitempty"`

	// string identifying a BDT Reference ID as defined in clause 5.3.3 of 3GPP TS 29.154.
	BdtRefId string `json:"bdtRefId,omitempty"`

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn string `json:"dnn,omitempty"`

	EvSubsc EventsSubscReqData `json:"evSubsc,omitempty"`

	// Indication of MCPTT service request.
	McpttId string `json:"mcpttId,omitempty"`

	// Indication of MCVideo service request.
	McVideoId string `json:"mcVideoId,omitempty"`

	// Contains media component information. The key of the map is the medCompN attribute.
	MedComponents map[string]MediaComponent `json:"medComponents,omitempty"`

	IpDomain string `json:"ipDomain,omitempty"`

	MpsAction MpsAction `json:"mpsAction,omitempty"`

	// Indication of MPS service request.
	MpsId string `json:"mpsId,omitempty"`

	// Indication of MCS service request.
	McsId string `json:"mcsId,omitempty"`

	PreemptControlInfo PreemptionControlInformation `json:"preemptControlInfo,omitempty"`

	ResPrio ReservPriority `json:"resPrio,omitempty"`

	ServInfStatus ServiceInfoStatus `json:"servInfStatus,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	NotifUri string `json:"notifUri"`

	// Contains values of the service URN and may include subservices.
	ServUrn string `json:"servUrn,omitempty"`

	SliceInfo Snssai `json:"sliceInfo,omitempty"`

	// Contains an identity of a sponsor.
	SponId string `json:"sponId,omitempty"`

	SponStatus SponsoringStatus `json:"sponStatus,omitempty"`

	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi string `json:"supi,omitempty"`

	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi string `json:"gpsi,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat string `json:"suppFeat"`

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	UeIpv4 string `json:"ueIpv4,omitempty"`

	UeIpv6 Ipv6Addr `json:"ueIpv6,omitempty"`

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042. 
	UeMac string `json:"ueMac,omitempty"`

	TsnBridgeManCont BridgeManagementContainer `json:"tsnBridgeManCont,omitempty"`

	TsnPortManContDstt PortManagementContainer `json:"tsnPortManContDstt,omitempty"`

	TsnPortManContNwtts []PortManagementContainer `json:"tsnPortManContNwtts,omitempty"`
}
