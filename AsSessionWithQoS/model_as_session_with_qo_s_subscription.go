/*
 * 3gpp-as-session-with-qos
 *
 * API for setting us an AS session with required QoS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_AsSessionWithQoS

// AsSessionWithQoSSubscription - Represents an individual AS session with required QoS subscription resource.
type AsSessionWithQoSSubscription struct {

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Self string `json:"self,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn string `json:"dnn,omitempty"`

	Snssai Snssai `json:"snssai,omitempty"`

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	NotificationDestination string `json:"notificationDestination"`

	// Identifies the external Application Identifier.
	ExterAppId string `json:"exterAppId,omitempty"`

	// Describe the data flow which requires QoS.
	FlowInfo []FlowInfo `json:"flowInfo,omitempty"`

	// Identifies Ethernet packet flows.
	EthFlowInfo []EthFlowDescription `json:"ethFlowInfo,omitempty"`

	// Identifies the Ethernet flows which require QoS. Each Ethernet flow consists of a flow idenifer and the corresponding UL and/or DL flows. 
	EnEthFlowInfo []EthFlowInfo `json:"enEthFlowInfo,omitempty"`

	// Identifies a pre-defined QoS information
	QosReference string `json:"qosReference,omitempty"`

	// Identifies an ordered list of pre-defined QoS information. The lower the index of the array for a given entry, the higher the priority.
	AltQoSReferences []string `json:"altQoSReferences,omitempty"`

	// Identifies an ordered list of alternative service requirements that include individual QoS parameter sets. The lower the index of the array for a given entry, the higher the priority.
	AltQosReqs []AlternativeServiceRequirementsData `json:"altQosReqs,omitempty"`

	// Indicates whether the QoS flow parameters signalling to the UE when the SMF is notified by the NG-RAN of changes in the fulfilled QoS situation is disabled (true) or not (false). Default value is false. The fulfilled situation is either the QoS profile or an Alternative QoS Profile. 
	DisUeNotif bool `json:"disUeNotif,omitempty"`

	// string identifying a Ipv4 address formatted in the \"dotted decimal\" notation as defined in IETF RFC 1166.
	UeIpv4Addr string `json:"ueIpv4Addr,omitempty"`

	IpDomain string `json:"ipDomain,omitempty"`

	// string identifying a Ipv6 address formatted according to clause 4 in IETF RFC 5952. The mixed Ipv4 Ipv6 notation according to clause 5 of IETF RFC 5952 shall not be used.
	UeIpv6Addr string `json:"ueIpv6Addr,omitempty"`

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042. 
	MacAddr string `json:"macAddr,omitempty"`

	UsageThreshold UsageThreshold `json:"usageThreshold,omitempty"`

	SponsorInfo SponsorInformation `json:"sponsorInfo,omitempty"`

	QosMonInfo QosMonitoringInformation `json:"qosMonInfo,omitempty"`

	// Indicates whether the direct event notification is requested (true) or not (false). Default value is false. 
	DirectNotifInd bool `json:"directNotifInd,omitempty"`

	TscQosReq TscQosRequirement `json:"tscQosReq,omitempty"`

	// Set to true by the SCS/AS to request the SCEF to send a test notification as defined in clause 5.2.5.3. Set to false or omitted otherwise.
	RequestTestNotification bool `json:"requestTestNotification,omitempty"`

	WebsockNotifConfig WebsockNotifConfig `json:"websockNotifConfig,omitempty"`

	// Represents the list of user plane event(s) to which the SCS/AS requests to subscribe to.
	Events []UserPlaneEvent `json:"events,omitempty"`
}
