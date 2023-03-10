/*
 * Nadrf_DataManagement
 *
 * ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nadrf_DataManagement

import (
	"time"
)

// AdditionalMeasurement - Represents additional measurement information.
type AdditionalMeasurement struct {

	UnexpLoc NetworkAreaInfo `json:"unexpLoc,omitempty"`

	UnexpFlowTeps []IpEthFlowDescription `json:"unexpFlowTeps,omitempty"`

	UnexpWakes []time.Time `json:"unexpWakes,omitempty"`

	DdosAttack AddressList `json:"ddosAttack,omitempty"`

	WrgDest AddressList `json:"wrgDest,omitempty"`

	Circums []CircumstanceDescription `json:"circums,omitempty"`
}
