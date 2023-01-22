/*
 * Ndccf_DataManagement
 *
 * DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ndccf_DataManagement

// AmfEventState - Represents the state of a subscribed event
type AmfEventState struct {

	Active bool `json:"active"`

	RemainReports int32 `json:"remainReports,omitempty"`

	// indicating a time in seconds.
	RemainDuration int32 `json:"remainDuration,omitempty"`
}
