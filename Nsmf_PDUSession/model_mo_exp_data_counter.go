/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nsmf_PDUSession

import (
	"time"
)

// MoExpDataCounter - Contain the MO Exception Data Counter.
type MoExpDataCounter struct {

	// Unsigned integer identifying the MO Exception Data Counter, as specified in clause 5.31.14.3 of 3GPP TS 23.501. 
	Counter int32 `json:"counter"`

	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp time.Time `json:"timeStamp,omitempty"`
}
