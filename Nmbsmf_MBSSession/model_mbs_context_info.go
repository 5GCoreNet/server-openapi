/*
 * Nmbsmf-MBSSession
 *
 * MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmbsmf_MBSSession

import (
	"time"
)

// MbsContextInfo - MBS context information
type MbsContextInfo struct {

	// string with format 'date-time' as defined in OpenAPI.
	StartTime time.Time `json:"startTime,omitempty"`

	AnyUeInd bool `json:"anyUeInd,omitempty"`

	LlSsm Ssm `json:"llSsm,omitempty"`

	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	CTeid int32 `json:"cTeid,omitempty"`

	MbsServiceArea MbsServiceArea `json:"mbsServiceArea,omitempty"`

	// A map (list of key-value pairs) where the key identifies an areaSessionId 
	MbsServiceAreaInfoList map[string]MbsServiceAreaInfo `json:"mbsServiceAreaInfoList,omitempty"`
}
