/*
 * Nnef_SMContext
 *
 * Nnef SMContext Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnef_SMContext

import (
	"time"
)

// SmallDataRateStatus - It indicates theSmall Data Rate Control Status
type SmallDataRateStatus struct {

	// When present, it shall contain the number of packets the UE is allowed to send uplink in the given time unit for the given PDU session (see clause 5.31.14.3 of 3GPP TS 23.501. 
	RemainPacketsUl int32 `json:"remainPacketsUl,omitempty"`

	// When present it shall contain the number of packets the AF is allowed to send downlink in the given time unit for the given PDU session (see clause 5.31.14.3 of 3GPP TS 23.501. 
	RemainPacketsDl int32 `json:"remainPacketsDl,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	ValidityTime time.Time `json:"validityTime,omitempty"`

	// When present, it shall indicate number of additional exception reports the UE is allowed to send uplink in the given time  unit for  the given PDU session (see clause 5.31.14.3 of 3GPP TS 23.501. 
	RemainExReportsUl int32 `json:"remainExReportsUl,omitempty"`

	// When present, it shall indicate number of additional exception reports the AF is allowed to send downlink  in the given time unit for the given PDU session (see clause 5.31.14.3 in 3GPP TS 23.501 
	RemainExReportsDl int32 `json:"remainExReportsDl,omitempty"`
}
