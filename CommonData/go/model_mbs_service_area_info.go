/*
 * Common Data Types
 *
 * Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   
 *
 * API version: 1.4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package CommonData

// MbsServiceAreaInfo - MBS Service Area Information for location dependent MBS session
type MbsServiceAreaInfo struct {

	// Integer where the allowed values correspond to the value range of an unsigned 16-bit integer.
	AreaSessionId int32 `json:"areaSessionId"`

	MbsServiceArea MbsServiceArea `json:"mbsServiceArea"`
}