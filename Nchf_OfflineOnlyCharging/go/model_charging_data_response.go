/*
 * Nchf_OfflineOnlyCharging
 *
 * OfflineOnlyCharging Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nchf_OfflineOnlyCharging

import (
	"time"
)

type ChargingDataResponse struct {

	// string with format 'date-time' as defined in OpenAPI.
	InvocationTimeStamp time.Time `json:"invocationTimeStamp"`

	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	InvocationSequenceNumber int32 `json:"invocationSequenceNumber"`

	InvocationResult InvocationResult `json:"invocationResult,omitempty"`

	SessionFailover SessionFailover `json:"sessionFailover,omitempty"`

	Triggers []Trigger `json:"triggers,omitempty"`

	PDUSessionChargingInformation PduSessionChargingInformation `json:"pDUSessionChargingInformation,omitempty"`

	RoamingQBCInformation RoamingQbcInformation `json:"roamingQBCInformation,omitempty"`
}
