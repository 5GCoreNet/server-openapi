/*
 * Nudm_PP
 *
 * Nudm Parameter Provision Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nudm_PP

import (
	"time"
)

type PpMaximumResponseTime struct {

	// indicating a time in seconds.
	MaximumResponseTime int32 `json:"maximumResponseTime"`

	AfInstanceId string `json:"afInstanceId"`

	ReferenceId int32 `json:"referenceId"`

	// string with format 'date-time' as defined in OpenAPI.
	ValidityTime time.Time `json:"validityTime,omitempty"`

	// String uniquely identifying MTC provider information.
	MtcProviderInformation string `json:"mtcProviderInformation,omitempty"`
}
