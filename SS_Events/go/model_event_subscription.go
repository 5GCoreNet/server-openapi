/*
 * SS_Events
 *
 * API for SEAL Events management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package SS_Events

// EventSubscription - Represents the subscription to a single SEAL event.
type EventSubscription struct {

	EventId SealEvent `json:"eventId"`

	// Each element of the array represents the VAL group identifier(s) of a VAL service that the subscriber wants to know in the interested event. 
	ValGroups []ValGroupFilter `json:"valGroups,omitempty"`

	// Each element of the array represents the VAL User / UE IDs of a VAL service that the event subscriber wants to know in the interested event. 
	Identities []IdentityFilter `json:"identities,omitempty"`

	// List of event monitoring details that the subscriber wishes to mmonitor the VAL UEs, VAL group and/or VAL service. 
	MonFltr []MonitorFilter `json:"monFltr,omitempty"`

	// Represents the list of VAL User / UE IDs and the area of interest information which the subscriber wishes to monitor the location deviation of VAL User / UEs. 
	AreaInt []MonitorLocationInterestFilter `json:"areaInt,omitempty"`

	// Each element represents the location area monitoring details to monitor the VA UEs moving in and out of the provided location area. 
	LocAreaMon []MonLocAreaInterestFltr `json:"locAreaMon,omitempty"`
}