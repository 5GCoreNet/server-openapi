/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nudm_SDM

// MbsfnArea - Contains an MBSFN area information.
type MbsfnArea struct {

	// This IE shall contain the MBSFN Area ID.
	MbsfnAreaId int32 `json:"mbsfnAreaId,omitempty"`

	// When present, this IE shall contain the Carrier Frequency (EARFCN).
	CarrierFrequency int32 `json:"carrierFrequency,omitempty"`
}
