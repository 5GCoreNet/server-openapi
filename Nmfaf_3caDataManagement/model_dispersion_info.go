/*
 * Nmfaf_3caDataManagement
 *
 * MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmfaf_3caDataManagement

import (
	"time"
)

// DispersionInfo - Represents the Dispersion information. When subscribed event is \"DISPERSION\", the  \"disperInfos\" attribute shall be included. 
type DispersionInfo struct {

	// string with format 'date-time' as defined in OpenAPI.
	TsStart time.Time `json:"tsStart"`

	// indicating a time in seconds.
	TsDuration int32 `json:"tsDuration"`

	DisperCollects []DispersionCollection `json:"disperCollects"`

	DisperType DispersionType `json:"disperType"`
}