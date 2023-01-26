/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_DataManagement

// PduSessionInfo - Represents session information.
type PduSessionInfo struct {

	N4SessId string `json:"n4SessId,omitempty"`

	// indicating a time in seconds.
	SessInactiveTimer int32 `json:"sessInactiveTimer,omitempty"`

	PduSessStatus PduSessionStatus `json:"pduSessStatus,omitempty"`
}