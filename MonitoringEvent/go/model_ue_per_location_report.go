/*
 * 3gpp-monitoring-event
 *
 * API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MonitoringEvent

// UePerLocationReport - Represents the number of UEs found at the indicated location.
type UePerLocationReport struct {

	// Identifies the number of UEs.
	UeCount int32 `json:"ueCount"`

	// Each element uniquely identifies a user.
	ExternalIds []string `json:"externalIds,omitempty"`

	// Each element identifies the MS internal PSTN/ISDN number allocated for a UE.
	Msisdns []string `json:"msisdns,omitempty"`

	// Each element uniquely identifies a UAV.
	ServLevelDevIds []string `json:"servLevelDevIds,omitempty"`
}
