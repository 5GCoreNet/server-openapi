/*
 * Namf_EventExposure
 *
 * AMF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Namf_EventExposure

import (
	"time"
)

// AmfEventMode - Describes how the reports shall be generated by a subscribed event
type AmfEventMode struct {

	Trigger AmfEventTrigger `json:"trigger"`

	MaxReports int32 `json:"maxReports,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	Expiry time.Time `json:"expiry,omitempty"`

	// indicating a time in seconds.
	RepPeriod int32 `json:"repPeriod,omitempty"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	SampRatio int32 `json:"sampRatio,omitempty"`

	PartitioningCriteria []PartitioningCriteria `json:"partitioningCriteria,omitempty"`

	NotifFlag NotificationFlag `json:"notifFlag,omitempty"`
}