/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Subscription_Data

import (
	"time"
)

type ExpectedUeBehaviourData struct {

	StationaryIndication StationaryIndication `json:"stationaryIndication,omitempty"`

	// indicating a time in seconds.
	CommunicationDurationTime int32 `json:"communicationDurationTime,omitempty"`

	// indicating a time in seconds.
	PeriodicTime int32 `json:"periodicTime,omitempty"`

	ScheduledCommunicationTime ScheduledCommunicationTime `json:"scheduledCommunicationTime,omitempty"`

	ScheduledCommunicationType ScheduledCommunicationType `json:"scheduledCommunicationType,omitempty"`

	// Identifies the UE's expected geographical movement. The attribute is only applicable in 5G.
	ExpectedUmts []LocationArea `json:"expectedUmts,omitempty"`

	TrafficProfile TrafficProfile `json:"trafficProfile,omitempty"`

	BatteryIndication BatteryIndication `json:"batteryIndication,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	ValidityTime time.Time `json:"validityTime,omitempty"`
}
