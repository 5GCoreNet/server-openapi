/*
 * Namf_Communication
 *
 * AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Namf_Communication

import (
	"time"
)

// UeContext - Represents an individual ueContext resource
type UeContext struct {

	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi string `json:"supi,omitempty"`

	SupiUnauthInd bool `json:"supiUnauthInd,omitempty"`

	GpsiList []string `json:"gpsiList,omitempty"`

	// String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.  
	Pei string `json:"pei,omitempty"`

	// Identifier of a group of NFs.
	UdmGroupId string `json:"udmGroupId,omitempty"`

	// Identifier of a group of NFs.
	AusfGroupId string `json:"ausfGroupId,omitempty"`

	// Identifier of a group of NFs.
	PcfGroupId string `json:"pcfGroupId,omitempty"`

	RoutingIndicator string `json:"routingIndicator,omitempty"`

	HNwPubKeyId int32 `json:"hNwPubKeyId,omitempty"`

	GroupList []string `json:"groupList,omitempty"`

	// string with format 'bytes' as defined in OpenAPI
	DrxParameter string `json:"drxParameter,omitempty"`

	// Unsigned integer representing the \"Subscriber Profile ID for RAT/Frequency Priority\"  as specified in 3GPP TS 36.413. 
	SubRfsp int32 `json:"subRfsp,omitempty"`

	// Unsigned integer representing the \"Subscriber Profile ID for RAT/Frequency Priority\"  as specified in 3GPP TS 36.413. 
	UsedRfsp int32 `json:"usedRfsp,omitempty"`

	SubUeAmbr Ambr `json:"subUeAmbr,omitempty"`

	// A map(list of key-value pairs) where Snssai serves as key.
	SubUeSliceMbrList map[string]SliceMbr `json:"subUeSliceMbrList,omitempty"`

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	SmsfId string `json:"smsfId,omitempty"`

	SeafData SeafData `json:"seafData,omitempty"`

	// string with format 'bytes' as defined in OpenAPI
	Var5gMmCapability string `json:"5gMmCapability,omitempty"`

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	PcfId string `json:"pcfId,omitempty"`

	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.  
	PcfSetId string `json:"pcfSetId,omitempty"`

	// NF Service Set Identifier (see clause 28.12 of 3GPP TS 23.003) formatted as the following  string \"set<Set ID>.sn<Service Name>.nfi<NF Instance ID>.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.sn<ServiceName>.nfi<NFInstanceID>.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)   <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NID> encoded as defined in clause 5.4.2 (\"Nid\" data type definition)  <NFInstanceId> encoded as defined in clause 5.3.2  <ServiceName> encoded as defined in 3GPP TS 29.510  <Set ID> encoded as a string of characters consisting of alphabetic    characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end    with either an alphabetic character or a digit. 
	PcfAmpServiceSetId string `json:"pcfAmpServiceSetId,omitempty"`

	// NF Service Set Identifier (see clause 28.12 of 3GPP TS 23.003) formatted as the following  string \"set<Set ID>.sn<Service Name>.nfi<NF Instance ID>.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.sn<ServiceName>.nfi<NFInstanceID>.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)   <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NID> encoded as defined in clause 5.4.2 (\"Nid\" data type definition)  <NFInstanceId> encoded as defined in clause 5.3.2  <ServiceName> encoded as defined in 3GPP TS 29.510  <Set ID> encoded as a string of characters consisting of alphabetic    characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end    with either an alphabetic character or a digit. 
	PcfUepServiceSetId string `json:"pcfUepServiceSetId,omitempty"`

	PcfBinding SbiBindingLevel `json:"pcfBinding,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	PcfAmPolicyUri string `json:"pcfAmPolicyUri,omitempty"`

	AmPolicyReqTriggerList []PolicyReqTrigger `json:"amPolicyReqTriggerList,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	PcfUePolicyUri string `json:"pcfUePolicyUri,omitempty"`

	UePolicyReqTriggerList []PolicyReqTrigger `json:"uePolicyReqTriggerList,omitempty"`

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	HpcfId string `json:"hpcfId,omitempty"`

	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.  
	HpcfSetId string `json:"hpcfSetId,omitempty"`

	RestrictedRatList []RatType `json:"restrictedRatList,omitempty"`

	ForbiddenAreaList []Area `json:"forbiddenAreaList,omitempty"`

	ServiceAreaRestriction ServiceAreaRestriction `json:"serviceAreaRestriction,omitempty"`

	RestrictedCoreNwTypeList []CoreNetworkType `json:"restrictedCoreNwTypeList,omitempty"`

	EventSubscriptionList []ExtAmfEventSubscription `json:"eventSubscriptionList,omitempty"`

	MmContextList []MmContext `json:"mmContextList,omitempty"`

	SessionContextList []PduSessionContext `json:"sessionContextList,omitempty"`

	EpsInterworkingInfo EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`

	TraceData *TraceData `json:"traceData,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	ServiceGapExpiryTime time.Time `json:"serviceGapExpiryTime,omitempty"`

	// String representing the STN-SR as defined in clause 18.6 of 3GPP TS 23.003.
	StnSr string `json:"stnSr,omitempty"`

	// String representing the C-MSISDN as defined in clause 18.7 of 3GPP TS 23.003.
	CMsisdn string `json:"cMsisdn,omitempty"`

	// string with format 'bytes' as defined in OpenAPI
	MsClassmark2 string `json:"msClassmark2,omitempty"`

	SupportedCodecList []string `json:"supportedCodecList,omitempty"`

	SmallDataRateStatusInfos []SmallDataRateStatusInfo `json:"smallDataRateStatusInfos,omitempty"`

	RestrictedPrimaryRatList []RatType `json:"restrictedPrimaryRatList,omitempty"`

	RestrictedSecondaryRatList []RatType `json:"restrictedSecondaryRatList,omitempty"`

	V2xContext V2xContext `json:"v2xContext,omitempty"`

	LteCatMInd bool `json:"lteCatMInd,omitempty"`

	RedCapInd bool `json:"redCapInd,omitempty"`

	MoExpDataCounter MoExpDataCounter `json:"moExpDataCounter,omitempty"`

	CagData CagData `json:"cagData,omitempty"`

	ManagementMdtInd bool `json:"managementMdtInd,omitempty"`

	ImmediateMdtConf ImmediateMdtConf `json:"immediateMdtConf,omitempty"`

	EcRestrictionDataWb EcRestrictionDataWb `json:"ecRestrictionDataWb,omitempty"`

	EcRestrictionDataNb bool `json:"ecRestrictionDataNb,omitempty"`

	IabOperationAllowed bool `json:"iabOperationAllowed,omitempty"`

	ProseContext ProseContext `json:"proseContext,omitempty"`

	AnalyticsSubscriptionList []AnalyticsSubscription `json:"analyticsSubscriptionList,omitempty"`

	PcfAmpBindingInfo string `json:"pcfAmpBindingInfo,omitempty"`

	PcfUepBindingInfo string `json:"pcfUepBindingInfo,omitempty"`

	UsedServiceAreaRestriction ServiceAreaRestriction `json:"usedServiceAreaRestriction,omitempty"`

	// A map(list of key-value pairs) where praId serves as key.
	PraInAmPolicy map[string]PresenceInfo `json:"praInAmPolicy,omitempty"`

	// A map(list of key-value pairs) where praId serves as key.
	PraInUePolicy map[string]PresenceInfo `json:"praInUePolicy,omitempty"`

	UpdpSubscriptionData UpdpSubscriptionData `json:"updpSubscriptionData,omitempty"`

	SmPolicyNotifyPduList []PduSessionInfo `json:"smPolicyNotifyPduList,omitempty"`

	PcfUeCallbackInfo *PcfUeCallbackInfo `json:"pcfUeCallbackInfo,omitempty"`

	// Positioning capabilities supported by the UE. A string encoding the \"ProvideCapabilities-r9-IEs\" IE as specified in clause 6.3 of 3GPP TS 37.355 (start from octet 1).
	UePositioningCap string `json:"uePositioningCap,omitempty"`

	AstiDistributionIndication bool `json:"astiDistributionIndication,omitempty"`

	TsErrorBudget int32 `json:"tsErrorBudget,omitempty"`

	SnpnOnboardInd bool `json:"snpnOnboardInd,omitempty"`

	SmfSelInfo *SmfSelectionData `json:"smfSelInfo,omitempty"`

	// A map(list of key-value pairs) where Snssai serves as key.
	PcfUeSliceMbrList map[string]SliceMbr `json:"pcfUeSliceMbrList,omitempty"`

	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.  
	SmsfSetId string `json:"smsfSetId,omitempty"`

	// NF Service Set Identifier (see clause 28.12 of 3GPP TS 23.003) formatted as the following  string \"set<Set ID>.sn<Service Name>.nfi<NF Instance ID>.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.sn<ServiceName>.nfi<NFInstanceID>.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)   <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NID> encoded as defined in clause 5.4.2 (\"Nid\" data type definition)  <NFInstanceId> encoded as defined in clause 5.3.2  <ServiceName> encoded as defined in 3GPP TS 29.510  <Set ID> encoded as a string of characters consisting of alphabetic    characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end    with either an alphabetic character or a digit. 
	SmsfServiceSetId string `json:"smsfServiceSetId,omitempty"`

	SmsfBindingInfo string `json:"smsfBindingInfo,omitempty"`
}
