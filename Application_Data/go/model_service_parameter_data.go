/*
 * Unified Data Repository Service API file for Application Data
 *
 * The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Application_Data

// ServiceParameterData - Represents the service parameter data.
type ServiceParameterData struct {

	// Identifies an application.
	AppId string `json:"appId,omitempty"`

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn string `json:"dnn,omitempty"`

	Snssai Snssai `json:"snssai,omitempty"`

	// String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.  
	InterGroupId string `json:"interGroupId,omitempty"`

	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi string `json:"supi,omitempty"`

	// string identifying a Ipv4 address formatted in the \"dotted decimal\" notation as defined in IETF RFC 1166.
	UeIpv4 string `json:"ueIpv4,omitempty"`

	// string identifying a Ipv6 address formatted according to clause 4 in IETF RFC 5952. The mixed Ipv4 Ipv6 notation according to clause 5 of IETF RFC 5952 shall not be used.
	UeIpv6 string `json:"ueIpv6,omitempty"`

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042. 
	UeMac string `json:"ueMac,omitempty"`

	AnyUeInd bool `json:"anyUeInd,omitempty"`

	// Represents configuration parameters for V2X communications over PC5 reference point. 
	ParamOverPc5 string `json:"paramOverPc5,omitempty"`

	// Represents configuration parameters for V2X communications over Uu reference point. 
	ParamOverUu string `json:"paramOverUu,omitempty"`

	// Represents the service parameters for 5G ProSe direct discovery.
	ParamForProSeDd string `json:"paramForProSeDd,omitempty"`

	// Represents the service parameters for 5G ProSe direct communications.
	ParamForProSeDc string `json:"paramForProSeDc,omitempty"`

	// Represents the service parameters for 5G ProSe UE-to-network relay UE.
	ParamForProSeU2NRelUe string `json:"paramForProSeU2NRelUe,omitempty"`

	// Represents the service parameters for 5G ProSe Remate UE.
	ParamForProSeRemUe string `json:"paramForProSeRemUe,omitempty"`

	// Contains the service parameter used to guide the URSP.
	UrspGuidance []UrspRuleRequest `json:"urspGuidance,omitempty"`

	// Contains the outcome of the UE Policy Delivery.
	DeliveryEvents []Event `json:"deliveryEvents,omitempty"`

	// Contains the Notification Correlation Id allocated by the NEF for the notification of UE Policy delivery outcome. 
	PolicDelivNotifCorreId string `json:"policDelivNotifCorreId,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	PolicDelivNotifUri string `json:"policDelivNotifUri,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat string `json:"suppFeat,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	ResUri string `json:"resUri,omitempty"`

	Headers []string `json:"headers,omitempty"`

	ResetIds []string `json:"resetIds,omitempty"`
}
