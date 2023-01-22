/*
 * CAPIF_Events_API
 *
 * API for event subscription management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package CAPIF_Events_API

import (
	"time"
)

// ReportingInformation - Represents the type of reporting that the subscription requires.
type ReportingInformation struct {

	ImmRep bool `json:"immRep,omitempty"`

	NotifMethod NotificationMethod `json:"notifMethod,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxReportNbr int32 `json:"maxReportNbr,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	MonDur time.Time `json:"monDur,omitempty"`

	// indicating a time in seconds.
	RepPeriod int32 `json:"repPeriod,omitempty"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	SampRatio int32 `json:"sampRatio,omitempty"`

	// Criteria for partitioning the UEs before applying the sampling ratio.
	PartitionCriteria []PartitioningCriteria `json:"partitionCriteria,omitempty"`

	// indicating a time in seconds.
	GrpRepTime int32 `json:"grpRepTime,omitempty"`

	NotifFlag NotificationFlag `json:"notifFlag,omitempty"`
}
