/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_AnalyticsInfo

import (
	"time"
)

// UeCommunication - Represents UE communication information.
type UeCommunication struct {

	// indicating a time in seconds.
	CommDur int32 `json:"commDur,omitempty"`

	// string with format 'float' as defined in OpenAPI.
	CommDurVariance float32 `json:"commDurVariance,omitempty"`

	// indicating a time in seconds.
	PerioTime int32 `json:"perioTime,omitempty"`

	// string with format 'float' as defined in OpenAPI.
	PerioTimeVariance float32 `json:"perioTimeVariance,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	Ts time.Time `json:"ts,omitempty"`

	// string with format 'float' as defined in OpenAPI.
	TsVariance float32 `json:"tsVariance,omitempty"`

	RecurringTime ScheduledCommunicationTime `json:"recurringTime,omitempty"`

	TrafChar TrafficCharacterization `json:"trafChar,omitempty"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	Ratio int32 `json:"ratio,omitempty"`

	PerioCommInd bool `json:"perioCommInd,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence int32 `json:"confidence,omitempty"`

	AnaOfAppList AppListForUeComm `json:"anaOfAppList,omitempty"`

	SessInactTimer SessInactTimerForUeComm `json:"sessInactTimer,omitempty"`
}
