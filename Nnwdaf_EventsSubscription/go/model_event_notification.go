/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_EventsSubscription

import (
	"time"
)

// EventNotification - Represents a notification on events that occurred.
type EventNotification struct {

	Event NwdafEvent `json:"event"`

	// string with format 'date-time' as defined in OpenAPI.
	Start time.Time `json:"start,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	Expiry time.Time `json:"expiry,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	TimeStampGen time.Time `json:"timeStampGen,omitempty"`

	FailNotifyCode NwdafFailureCode `json:"failNotifyCode,omitempty"`

	// indicating a time in seconds.
	RvWaitTime int32 `json:"rvWaitTime,omitempty"`

	AnaMetaInfo AnalyticsMetadataInfo `json:"anaMetaInfo,omitempty"`

	NfLoadLevelInfos []NfLoadLevelInformation `json:"nfLoadLevelInfos,omitempty"`

	NsiLoadLevelInfos []NsiLoadLevelInfo `json:"nsiLoadLevelInfos,omitempty"`

	SliceLoadLevelInfo SliceLoadLevelInformation `json:"sliceLoadLevelInfo,omitempty"`

	SvcExps []ServiceExperienceInfo `json:"svcExps,omitempty"`

	QosSustainInfos []QosSustainabilityInfo `json:"qosSustainInfos,omitempty"`

	UeComms []UeCommunication `json:"ueComms,omitempty"`

	UeMobs []UeMobility `json:"ueMobs,omitempty"`

	UserDataCongInfos []UserDataCongestionInfo `json:"userDataCongInfos,omitempty"`

	AbnorBehavrs []AbnormalBehaviour `json:"abnorBehavrs,omitempty"`

	NwPerfs []NetworkPerfInfo `json:"nwPerfs,omitempty"`

	DnPerfInfos []DnPerfInfo `json:"dnPerfInfos,omitempty"`

	DisperInfos []DispersionInfo `json:"disperInfos,omitempty"`

	RedTransInfos []RedundantTransmissionExpInfo `json:"redTransInfos,omitempty"`

	WlanInfos []WlanPerformanceInfo `json:"wlanInfos,omitempty"`

	SmccExps []SmcceInfo `json:"smccExps,omitempty"`
}
