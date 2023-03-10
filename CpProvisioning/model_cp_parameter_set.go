/*
 * 3gpp-cp-parameter-provisioning
 *
 * API for provisioning communication pattern parameters.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CpProvisioning

import (
	"time"
)

// CpParameterSet - Represents an offered communication pattern parameter set.
type CpParameterSet struct {

	// SCS/AS-chosen correlator provided by the SCS/AS in the request to create a resource fo CP parameter set(s).
	SetId string `json:"setId"`

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Self string `json:"self,omitempty"`

	// string with format \"date-time\" as defined in OpenAPI.
	ValidityTime time.Time `json:"validityTime,omitempty"`

	PeriodicCommunicationIndicator CommunicationIndicator `json:"periodicCommunicationIndicator,omitempty"`

	// Unsigned integer identifying a period of time in units of seconds.
	CommunicationDurationTime int32 `json:"communicationDurationTime,omitempty"`

	// Unsigned integer identifying a period of time in units of seconds.
	PeriodicTime int32 `json:"periodicTime,omitempty"`

	ScheduledCommunicationTime ScheduledCommunicationTime `json:"scheduledCommunicationTime,omitempty"`

	ScheduledCommunicationType ScheduledCommunicationType `json:"scheduledCommunicationType,omitempty"`

	StationaryIndication StationaryIndication `json:"stationaryIndication,omitempty"`

	BatteryInds []BatteryIndication `json:"batteryInds,omitempty"`

	TrafficProfile TrafficProfile `json:"trafficProfile,omitempty"`

	// Identifies the UE's expected geographical movement. The attribute is only applicable in 5G.
	ExpectedUmts []UmtLocationArea5G `json:"expectedUmts,omitempty"`

	// integer between and including 1 and 7 denoting a weekday. 1 shall indicate Monday, and the subsequent weekdays shall be indicated with the next higher numbers. 7 shall indicate Sunday.
	ExpectedUmtDays int32 `json:"expectedUmtDays,omitempty"`
}
