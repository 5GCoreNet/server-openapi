/*
 * Naf_EventExposure
 *
 * AF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Naf_EventExposure

// Exception - Represents the Exception information.
type Exception struct {

	ExcepId ExceptionId `json:"excepId"`

	ExcepLevel int32 `json:"excepLevel,omitempty"`

	ExcepTrend ExceptionTrend `json:"excepTrend,omitempty"`
}
