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

type MonitoringReport struct {

	ReferenceId int32 `json:"referenceId"`

	EventType EventType `json:"eventType"`

	Report Report `json:"report,omitempty"`

	ReachabilityForSmsReport ReachabilityForSmsReport `json:"reachabilityForSmsReport,omitempty"`

	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi string `json:"gpsi,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp time.Time `json:"timeStamp"`

	ReachabilityReport ReachabilityReport `json:"reachabilityReport,omitempty"`
}
