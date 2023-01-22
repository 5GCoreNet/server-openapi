/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_AnalyticsInfo

// NumberAverage - Represents average and variance information.
type NumberAverage struct {

	// string with format 'float' as defined in OpenAPI.
	Number float32 `json:"number"`

	// string with format 'float' as defined in OpenAPI.
	Variance float32 `json:"variance"`

	// string with format 'float' as defined in OpenAPI.
	Skewness float32 `json:"skewness,omitempty"`
}
