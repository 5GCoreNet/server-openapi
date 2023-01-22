/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Subscription_Data

// UeSubscribedDataSets - Contains the subscribed data sets of a UE.
type UeSubscribedDataSets struct {

	AmData AccessAndMobilitySubscriptionData `json:"amData,omitempty"`

	SmfSelData SmfSelectionSubscriptionData `json:"smfSelData,omitempty"`

	SmsSubsData SmsSubscriptionData `json:"smsSubsData,omitempty"`

	SmData SmSubsData `json:"smData,omitempty"`

	TraceData *TraceData `json:"traceData,omitempty"`

	SmsMngData SmsManagementSubscriptionData `json:"smsMngData,omitempty"`

	LcsPrivacyData LcsPrivacyData `json:"lcsPrivacyData,omitempty"`

	LcsMoData LcsMoData `json:"lcsMoData,omitempty"`

	LcsBcaData LcsBroadcastAssistanceTypesData `json:"lcsBcaData,omitempty"`

	V2xData V2xSubscriptionData `json:"v2xData,omitempty"`

	ProseData ProseSubscriptionData `json:"proseData,omitempty"`

	OdbData OdbData `json:"odbData,omitempty"`

	EeProfileData EeProfileData `json:"eeProfileData,omitempty"`

	PpProfileData PpProfileData `json:"ppProfileData,omitempty"`

	NiddAuthData AuthorizationData `json:"niddAuthData,omitempty"`

	UcData UcSubscriptionData `json:"ucData,omitempty"`

	MbsSubscriptionData MbsSubscriptionData `json:"mbsSubscriptionData,omitempty"`
}
