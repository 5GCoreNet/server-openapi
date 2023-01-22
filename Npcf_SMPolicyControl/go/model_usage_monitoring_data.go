/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Npcf_SMPolicyControl

import (
	"time"
)

// UsageMonitoringData - Contains usage monitoring related control information.
type UsageMonitoringData struct {

	// Univocally identifies the usage monitoring policy data within a PDU session.
	UmId string `json:"umId"`

	// Unsigned integer identifying a volume in units of bytes with \"nullable=true\" property.
	VolumeThreshold *int64 `json:"volumeThreshold,omitempty"`

	// Unsigned integer identifying a volume in units of bytes with \"nullable=true\" property.
	VolumeThresholdUplink *int64 `json:"volumeThresholdUplink,omitempty"`

	// Unsigned integer identifying a volume in units of bytes with \"nullable=true\" property.
	VolumeThresholdDownlink *int64 `json:"volumeThresholdDownlink,omitempty"`

	// indicating a time in seconds with OpenAPI defined 'nullable: true' property.
	TimeThreshold *int32 `json:"timeThreshold,omitempty"`

	// string with format 'date-time' as defined in OpenAPI with 'nullable:true' property.  
	MonitoringTime *time.Time `json:"monitoringTime,omitempty"`

	// Unsigned integer identifying a volume in units of bytes with \"nullable=true\" property.
	NextVolThreshold *int64 `json:"nextVolThreshold,omitempty"`

	// Unsigned integer identifying a volume in units of bytes with \"nullable=true\" property.
	NextVolThresholdUplink *int64 `json:"nextVolThresholdUplink,omitempty"`

	// Unsigned integer identifying a volume in units of bytes with \"nullable=true\" property.
	NextVolThresholdDownlink *int64 `json:"nextVolThresholdDownlink,omitempty"`

	// indicating a time in seconds with OpenAPI defined 'nullable: true' property.
	NextTimeThreshold *int32 `json:"nextTimeThreshold,omitempty"`

	// indicating a time in seconds with OpenAPI defined 'nullable: true' property.
	InactivityTime *int32 `json:"inactivityTime,omitempty"`

	// Contains the PCC rule identifier(s) which corresponding service data flow(s) shall be excluded from PDU Session usage monitoring. It is only included in the UsageMonitoringData instance for session level usage monitoring. 
	ExUsagePccRuleIds *[]string `json:"exUsagePccRuleIds,omitempty"`
}
